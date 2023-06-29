package main

import (
	"flag"
	"fmt"
	"strings"
)

type Command struct {
	decoder Codec
	encoder Codec
}

func ParseCommand(raw_args []string) (*Command, error) {
	flagSet := flag.NewFlagSet("renc", flag.ExitOnError)
	flagSet.Parse(raw_args)
  args := flagSet.Args()

	var decoder Codec = &BinaryCodec{}
	var encoder Codec = &BinaryCodec{}
	var err error

	if len(args) > 2 {
		return nil, fmt.Errorf("Usage: renc [decoder=raw] [encoder=raw]\nError: too many arguments")
	}
	if len(args) > 0 {
		decoder, err = CodecByName(args[0])
		if err != nil {
			return nil, fmt.Errorf("invalid decoder: %w", err)
		}
	}
	if len(args) > 1 {
		encoder, err = CodecByName(args[1])
		if err != nil {
			return nil, fmt.Errorf("invalid encoder: %w", err)
		}
	}

	return &Command{decoder, encoder}, nil
}

var usage = `Usage: renc [decoder=raw] [encoder=raw]`
var help = `
Available encoders/decoders:
  - raw
  - hex
  - base32
  - base64
`

func Usage() string {
	return usage
}

func Help() string {
  return fmt.Sprintf("%s\n%s\n", usage, strings.TrimSpace(help))
}

func ErrorMsg(err error) string {
  return fmt.Sprintf("Error: %s", err.Error())
}
