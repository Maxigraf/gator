package command

import "fmt"

type Commands interface {
	Register(string, func(*State, Command) error)
	Run(*State, Command) error
}

func NewCommands() Commands {
	cmds := Commands(&commands{})
	registerHandlers(cmds)
	return cmds
}

type commands struct {
	handler map[string]func(*State, Command) error
}

func (c *commands) Register(name string, f func(*State, Command) error) {
	if c.handler == nil {
		c.handler = make(map[string]func(*State, Command) error)
	}

	c.handler[name] = f
}

func (c *commands) Run(s *State, cmd Command) error {
	if c.handler == nil {
		return fmt.Errorf("no commands available")
	}

	handler, exists := c.handler[cmd.Name]

	if !exists {
		return fmt.Errorf("invalid command '%s'", cmd.Name)
	}

	return handler(s, cmd)
}
