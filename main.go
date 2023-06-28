package main

import (
	"fmt"
	"os"
)

func main() {
  cmd := ParseCommand(os.Args[1:])

  codecs := make([]*Codec, 2)
  codecs[0] = ByName(cmd.encoder)
  codecs[1] = ByName(cmd.decoder)

      

  fmt.Println(codecs)
}
