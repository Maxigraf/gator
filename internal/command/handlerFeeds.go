package command

import (
	"context"
	"fmt"
)

func handlerFeeds(s *State, cmd Command) error {
	feeds, err := s.Database.GetFeeds(context.Background())

	if err != nil {
		return fmt.Errorf("could not get feeds: %v", err)
	}

	fmt.Println("Available feeds:")

	for _, feed := range feeds {
		user, err := s.Database.GetUserById(context.Background(), feed.UserID)

		if err != nil {
			return fmt.Errorf("could not get user: %v", err)
		}

		fmt.Printf("* Feed '%s' [%s] added by %s\n", feed.Name, feed.Url, user.Name)
	}

	return nil
}
