package cmd

import "errors"

type Command interface {
	Execute() error
}

type Handler struct{}

func (h Handler) MatchAndRun(arg string) error {
	cmds := map[string]Command{
		"help":    Help{},
		"serve":   Serve{},
		"migrate": Migrate{},
		"seed":    Seed{},
	}

	if c, ok := cmds[arg]; ok {
		return c.Execute()
	}

	return errors.New("Argument not found.")
}
