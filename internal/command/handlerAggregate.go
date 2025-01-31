package command

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/maxigraf/gator/internal/database"
	"github.com/maxigraf/gator/internal/feed"
)

func handlerAggregate(s *State, cmd Command) error {
	time_between_reqs, err := time.ParseDuration(cmd.Args[0])

	if err != nil {
		return fmt.Errorf("invalid duration: %v", err)
	}

	fmt.Printf("Colleting feeds every %v\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *State) {
	feedToFetch, err := s.Database.GetNextFeedToFetch(context.Background())

	if err != nil {
		fmt.Printf("Could not load feed: %v\n", err)
		return
	}

	param := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: feedToFetch.ID,
	}

	err = s.Database.MarkFeedFetched(context.Background(), param)

	if err != nil {
		fmt.Printf("Failed to mark as fetched: %v\n", err)
		return
	}

	content, err := feed.FetchFeed(context.Background(), feedToFetch.Url)

	if err != nil {
		fmt.Printf("Failed to fetch feed data: %v\n", err)
		return
	}

	savePosts(s, feedToFetch, content)
}

func savePosts(s *State, feed database.Feed, content *feed.RSSFeed) {
	for _, item := range content.Channel.Item {
		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)

		if err != nil {
			pubDate = time.Now()
		}

		description := sql.NullString{}

		if len(item.Description) > 0 {
			description.String = item.Description
			description.Valid = true
		}

		param := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: description,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		}

		_, err = s.Database.CreatePost(context.Background(), param)

		if err != nil {
			target := &pq.Error{}
			if errors.As(err, &target) {
				if target.Constraint == "posts_url_key" {
					return
				}
			}
			fmt.Printf("could not store post: %s %+v\n", reflect.TypeOf(err), err)
		}
	}
}
