package main

import (
	"flag"
)

type Command struct {
	decoder string
	encoder string
}

func ParseCommand(args []string) *Command {
	flagSet := flag.NewFlagSet("renc", flag.ExitOnError)
	flagSet.Parse(args)

  codecNames := flagSet.Args()[0:2]

	return &Command{
    decoder: codecNames[0],
    encoder: codecNames[1],
  }
}
