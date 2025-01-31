package command

import (
	"context"
	"fmt"

	"github.com/maxigraf/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	ensureLoggedIn := func(s *State, cmd Command) error {
		user, err := s.Database.GetUser(context.Background(), s.Config.CurrentUserName)

		if err != nil {
			return fmt.Errorf("not logged in or user doesn't exist: %v", err)
		}

		return handler(s, cmd, user)
	}

	return ensureLoggedIn
}
