package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fail(fmt.Errorf("%v", r))
		}
	}()

	cmd, err := ParseCommand(os.Args[1:])
	if err != nil {
		fail(err)
	}
  
  decoder := cmd.decoder.NewDecoder(os.Stdin)
  encoder := cmd.encoder.NewEncoder(os.Stdout)

	_, err = io.Copy(encoder, decoder)
	if err != nil {
		fail(err)
	}

  encoder.Close()
}

func fail(err error) {
  os.Stderr.WriteString(fmt.Sprintf("%s\n%s\n", Help(), ErrorMsg(err)))
	os.Exit(1)
}
