package command

import (
	"context"
	"fmt"
)

func handlerFollowing(s *State, cmd Command) error {
	user, err := s.Database.GetUser(context.Background(), s.Config.CurrentUserName)

	if err != nil {
		return fmt.Errorf("not logged in or user doesn't exist: %v", err)
	}

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
