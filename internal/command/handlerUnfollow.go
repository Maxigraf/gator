package command

import (
	"context"
	"fmt"

	"github.com/maxigraf/gator/internal/database"
)

func handlerUnfollow(s *State, cmd Command, user database.User) error {
	feed, err := s.Database.GetFeedByUrl(context.Background(), cmd.Args[0])

	if err != nil {
		return fmt.Errorf("could not get feed: %v", err)
	}

	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.Database.DeleteFeedFollow(context.Background(), params)

	if err != nil {
		return fmt.Errorf("unfollow failed: %v", err)
	}

	fmt.Printf("%s no longer following '%s'\n", user.Name, feed.Name)

	return nil
}
