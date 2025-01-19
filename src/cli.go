package main

import (
	"flag"
	"fmt"
	"strings"
)

const DEFAULT_ENCODER = "raw"
const DEFAULT_DECODER = "raw"

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

	if len(args) == 0 {
		return nil, fmt.Errorf("no [encoder]:[decoder] argument")

	} else if len(args) > 1 {
		return nil, fmt.Errorf("too many arguments. Did you mean '[encoder]:[decoder]' in a single argument?")
	}

	codecNames, err := parseCodecPair(args[0])
	if err != nil {
		return nil, fmt.Errorf("invalid argument: %w", err)
	}

	decoder, err = CodecByName(codecNames[0])
	if err != nil {
		return nil, fmt.Errorf("invalid decoder: %w", err)
	}

	encoder, err = CodecByName(codecNames[1])
	if err != nil {
		return nil, fmt.Errorf("invalid encoder: %w", err)
	}

	return &Command{decoder, encoder}, nil
}

func parseCodecPair(arg string) ([]string, error) {
	names := strings.Split(arg, ":")

	if len(names) != 2 {
		return nil, fmt.Errorf("invalid '[encoder]:[decoder]' argument: %s", arg)
	}

	if names[0] == "" {
		names[0] = DEFAULT_DECODER
	}

	if names[1] == "" {
		names[1] = DEFAULT_ENCODER
	}

	return names, nil
}

var usage = `Usage: renc [decoder=raw]:[encoder=raw]`
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
