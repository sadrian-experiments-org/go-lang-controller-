package main

import (
	"flag"
	"fmt"
	"os"
	"cli/controller"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command>. The only one is tf-api-run")
		os.Exit(1)
	}

	// Create a new subcommand dispatcher
	cmdDispatcher := flag.NewFlagSet("cli", flag.ExitOnError)
	cmdDispatcher.Parse(os.Args[1:2])

	switch cmdDispatcher.Arg(0) {
	case "api-run":
		controller.ApiRunController()
	default:
		fmt.Println("Unknown command:", cmdDispatcher.Arg(0))
	}
}