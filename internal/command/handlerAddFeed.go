package command

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/maxigraf/gator/internal/database"
)

func handlerAddFeed(s *State, cmd Command, user database.User) error {
	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}

	feed, err := s.Database.CreateFeed(context.Background(), params)

	if err != nil {
		return fmt.Errorf("failed to create feed: %v", err)
	}

	fmt.Printf("Feed '%s' created: %+v\n", feed.Name, feed)

	_, err = addFeedFollow(s, user, feed)

	if err != nil {
		return err
	}

	return nil
}
