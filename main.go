package main

import (
	"log"
	"os"

	"github.com/bjvanbemmel/game-store/cmd"
)

func main() {
	var help cmd.Help = cmd.Help{}
	var args []string = os.Args

	if len(args) < 2 {
		help.Execute()
		return
	}

	var handler cmd.Handler = cmd.Handler{}

	err := handler.MatchAndRun(args[1])
	if err != nil {
		log.Fatal(err.Error())
	}
}
