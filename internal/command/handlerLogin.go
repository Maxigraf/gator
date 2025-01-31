package command

import (
	"context"
	"fmt"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("missing arguments")
	}

	user, err := s.Database.GetUser(context.Background(), cmd.Args[0])

	if err != nil {
		return fmt.Errorf("user does not exist")
	}

	err = s.Config.SetUser(user.Name)

	if err != nil {
		return fmt.Errorf("could not set username: %v", err)
	}

	fmt.Printf("User '%s' has been set \n", s.Config.CurrentUserName)

	return nil
}
