package command

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/maxigraf/gator/internal/database"
)

func handlerFollow(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("missing arguments")
	}

	user, err := s.Database.GetUser(context.Background(), s.Config.CurrentUserName)

	if err != nil {
		return fmt.Errorf("not logged in or user doesn't exist: %v", err)
	}

	feed, err := s.Database.GetFeedByUrl(context.Background(), cmd.Args[0])

	if err != nil {
		return fmt.Errorf("could not get feed: %v", err)
	}

	feedFollow, err := addFeedFollow(s, user, feed)

	if err != nil {
		return err
	}

	fmt.Printf("User %s now following feed '%s'\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}

func addFeedFollow(s *State, user database.User, feed database.Feed) (database.CreateFeedFollowRow, error) {
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.Database.CreateFeedFollow(context.Background(), params)

	if err != nil {
		return database.CreateFeedFollowRow{}, fmt.Errorf("failed to follow feed: %v", err)
	}

	return feedFollow, nil
}
