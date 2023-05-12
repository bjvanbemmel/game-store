package cmd

import "errors"

type Command interface {
	Execute() error
}

var ArgNotFoundErr error = errors.New("Argument not found.")

type Handler struct{}

func (h Handler) Match(arg string) (Command, error) {
	cmds := map[string]Command{
		"help":    Help{},
		"serve":   Serve{},
		"migrate": Migrate{},
		"seed":    Seed{},
	}

	if c, ok := cmds[arg]; ok {
		return c, nil
	}

	return nil, ArgNotFoundErr
}
