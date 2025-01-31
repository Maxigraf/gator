package command

import (
	"fmt"
)

func middlewareArgsRequired(handler func(s *State, cmd Command) error) func(*State, Command) error {
	return middlewareArgsCountRequired(handler, 1)
}

func middlewareArgsCountRequired(handler func(s *State, cmd Command) error, minCount int) func(*State, Command) error {
	ensureArgsPresent := func(s *State, cmd Command) error {
		if len(cmd.Args) < minCount {
			return fmt.Errorf("missing arguments")
		}

		return handler(s, cmd)
	}

	return ensureArgsPresent
}
