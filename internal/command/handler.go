package command

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/maxigraf/gator/internal/database"
)

func HandlerLogin(s *State, cmd Command) error {
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

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("missing arguments")
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	user, err := s.Database.CreateUser(context.Background(), params)

	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	s.Config.SetUser(user.Name)

	fmt.Printf("User '%s' created: %v\n", cmd.Args[0], user)

	return nil
}

func HandlerReset(s *State, cmd Command) error {
	err := s.Database.Reset(context.Background())

	if err != nil {
		return fmt.Errorf("reset failed: %v", err)
	}

	fmt.Println("Reset successful")

	return nil
}
