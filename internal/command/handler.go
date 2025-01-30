package command

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/maxigraf/gator/internal/database"
	"github.com/maxigraf/gator/internal/feed"
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

func HandlerUsers(s *State, cmd Command) error {
	users, err := s.Database.GetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("failed to get users: %v", err)
	}

	for _, user := range users {
		fmt.Printf("* %s", user.Name)
		if user.Name == s.Config.CurrentUserName {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}

	return nil
}

func HandlerAggregate(s *State, cmd Command) error {
	result, err := feed.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		return fmt.Errorf("failed to get feed: %v", err)
	}

	fmt.Printf("%+v\n", result)

	return nil
}

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("missing arguments")
	}

	user, err := s.Database.GetUser(context.Background(), s.Config.CurrentUserName)

	if err != nil {
		return fmt.Errorf("not logged in or user doesn't exist: %v", err)
	}

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

	return nil
}
