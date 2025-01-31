package command

import (
	"context"
	"fmt"

	"github.com/maxigraf/gator/internal/database"
)

func handlerFollowing(s *State, cmd Command, user database.User) error {
	feedFollows, err := s.Database.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return fmt.Errorf("could not get feed follows: %v", err)
	}

	fmt.Println("Followed feeds:")

	for _, feedFollow := range feedFollows {
		fmt.Printf("* %s follows feed '%s'\n", feedFollow.UserName, feedFollow.FeedName)
	}
	return nil
}
