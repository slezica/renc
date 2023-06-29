package main

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

type Codec interface {
  Name() string
  NewEncoder(w io.Writer) io.Writer
  NewDecoder(r io.Reader) io.Reader
}

// -------------------------------------------------------------------------------------------------

var codecs = []Codec{
  &BinaryCodec{},
  &HexCodec{},
  &Base32Codec{},
  &Base64Codec{},
}

var codecsByName = map[string]Codec{}

func CodecByName(name string) (Codec, error) {
  if len(codecsByName) == 0 {
    for _, codec := range codecs {
      codecsByName[codec.Name()] = codec
    }
  }

  codec, ok := codecsByName[name]
  if !ok {
    return nil, fmt.Errorf("no such codec: %s", name)
  }

  return codec, nil
}

// -------------------------------------------------------------------------------------------------

type BinaryCodec struct {}

func (e *BinaryCodec) Name() string {
  return "dummy"
}

func (e *BinaryCodec) NewEncoder(w io.Writer) io.Writer {
  return w
}

func (e *BinaryCodec) NewDecoder(r io.Reader) io.Reader {
  return r
}

// -------------------------------------------------------------------------------------------------

type HexCodec struct {}

func (e *HexCodec) Name() string {
  return "hex"
}

func (e *HexCodec) NewEncoder(w io.Writer) io.Writer {
  return hex.NewEncoder(w)
}

func (e *HexCodec) NewDecoder(r io.Reader) io.Reader {
  return hex.NewDecoder(r)
}

// -------------------------------------------------------------------------------------------------

type Base32Codec struct {}

func (e *Base32Codec) Name() string {
  return "base32"
}

func (e *Base32Codec) NewEncoder(w io.Writer) io.Writer {
  return base32.NewEncoder(base32.StdEncoding, w)
}

func (e *Base32Codec) NewDecoder(r io.Reader) io.Reader {
  return base32.NewDecoder(base32.StdEncoding, r)
}

// -------------------------------------------------------------------------------------------------


type Base64Codec struct {}

func (e *Base64Codec) Name() string {
  return "base64"
}

func (e *Base64Codec) NewEncoder(w io.Writer) io.Writer {
  return base32.NewEncoder(base32.StdEncoding, w)
}

func (e *Base64Codec) NewDecoder(r io.Reader) io.Reader {
  return base64.NewDecoder(base64.StdEncoding, r)
}
