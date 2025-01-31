package command

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/maxigraf/gator/internal/database"
)

func handlerRegister(s *State, cmd Command) error {
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

	fmt.Printf("User '%s' created: %+v\n", cmd.Args[0], user)

	return nil
}
