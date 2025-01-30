package command

import "fmt"

type Commands struct {
	handler map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	if c.handler == nil {
		c.handler = make(map[string]func(*State, Command) error)
	}

	c.handler[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	if c.handler == nil {
		return fmt.Errorf("no commands available")
	}

	handler, exists := c.handler[cmd.Name]

	if !exists {
		return fmt.Errorf("invalid command '%s'", cmd.Name)
	}

	return handler(s, cmd)
}
