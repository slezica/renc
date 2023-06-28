package main

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
)

type Codec struct {
	name   string
	encode func([]byte) string
	decode func(string) ([]byte, error)
}

func (c *Codec) Name() string {
	return c.name
}

func (c *Codec) Encode(data []byte) string {
	return c.encode(data)
}

func (c *Codec) Decode(text string) ([]byte, error) {
	return c.decode(text)
}

var (
	Ascii  = &Codec{"ascii", encodeAscii, decodeAscii}
	Hex    = &Codec{"hex", hex.EncodeToString, hex.DecodeString}
	Base32 = &Codec{"base32", base32.StdEncoding.EncodeToString, decodeAnyBase32}
	Base64 = &Codec{"base64", base64.StdEncoding.EncodeToString, decodeAnyBase64}
)

// -------------------------------------------------------------------------------------------------

var codecByName = map[string]*Codec{
	Ascii.name:  Ascii,
	Hex.name:    Hex,
	Base32.name: Base32,
	Base64.name: Base64,
}

func ByName(name string) *Codec {
	return codecByName[name]
}

// -------------------------------------------------------------------------------------------------

func encodeAscii(data []byte) string {
	return string(data)
}

func decodeAscii(text string) ([]byte, error) {
	return []byte(text), nil
}

func decodeAnyBase32(text string) ([]byte, error) {
	data, stdErr := base32.StdEncoding.DecodeString(text)
	if stdErr == nil {
		return data, nil
	}

	data, err := base32.HexEncoding.DecodeString(text)
	if err == nil {
		return data, nil
	}

	return nil, stdErr
}

func decodeAnyBase64(text string) ([]byte, error) {
	data, stdErr := base64.StdEncoding.DecodeString(text)
	if stdErr == nil {
		return data, nil
	}

	data, err := base64.RawStdEncoding.DecodeString(text)
	if err == nil {
		return data, nil
	}

	data, err = base64.URLEncoding.DecodeString(text)
	if err == nil {
		return data, nil
	}

	data, err = base64.RawURLEncoding.DecodeString(text)
	if err == nil {
		return data, nil
	}

	return nil, stdErr
}
