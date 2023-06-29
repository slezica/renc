
`renc` is a command-line tool to decode/encode streams using a variety of encodings.

### Installation

Grab a binary from the `bin/` directory and place it in your `$PATH`, or use the `Makefile`.

### Usage

`renc` decodes `stdin` and re-encodes into `stdout`. It takes two arguments, a decoder and an encoder.

```sh
renc [decoder] [encoder=raw]
```

The `decoder` and `encoder` arguments can be:

* `raw`: just let bytes through
* `hex`: hexadecimal
* `base32`: standard base32
* `base64`: standard base64

### Examples

Decode `base64`, then encode into `hex`:


`` `
$ echo -n "SGVsbG8gd29ybGQh" | renc base64 hex
48656c6c6f20776f726c6421
```

Decode `hex` into raw bytes (i.e. no re-encoding):
```
$ echo -n "48656c6c6f20776f726c6421" | renc hex
Hello world!
```

Encode a binary file to `base64`:

```
$ cat data.bin | renc raw base64
UHJldHR5IGxhbWUgZWFzdGVyIGVnZywgaHVoPw==
```
