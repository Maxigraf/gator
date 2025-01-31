package command

import (
	"context"
	"fmt"
)

func handlerReset(s *State, cmd Command) error {
	err := s.Database.Reset(context.Background())

	if err != nil {
		return fmt.Errorf("reset failed: %v", err)
	}

	fmt.Println("Reset successful")

	return nil
}
