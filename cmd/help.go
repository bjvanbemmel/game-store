package cmd

import (
	"fmt"

	"github.com/bjvanbemmel/game-store/docs"
)

type Help struct{}

func (h Help) Execute() error {
	fmt.Print(string(docs.HelpEmbed))

	return nil
}
