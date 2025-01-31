package command

import (
	"context"
	"fmt"
)

func handlerUsers(s *State, cmd Command) error {
	users, err := s.Database.GetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("failed to get users: %v", err)
	}

	fmt.Println("Registered users:")

	for _, user := range users {
		fmt.Printf("* %s", user.Name)
		if user.Name == s.Config.CurrentUserName {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}

	return nil
}
