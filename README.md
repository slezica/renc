
`renc` is a command-line tool to decode/encode streams using a variety of encodings.

### Installation

Grab a binary from the `bin/` directory and place it in your `$PATH`, or use the `Makefile`.

### Usage

`renc` decodes `stdin` and re-encodes into `stdout`. It takes one argument in the form `[decoder]:[encoder]` (both parts are optional and default to `raw`).

```sh
renc [decoder=raw]:[encoder=raw]
```

The `decoder` and `encoder` parameters can be:

* `raw`: just let bytes through
* `hex`: hexadecimal
* `base32`: standard base32
* `base64`: standard base64

### Examples

Decode raw hex bytes (no re-encoding):
```
$ echo -n "48656c6c6f20776f726c6421" | renc hex:
Hello world!
```

Decode `base64`, then encode into `hex`:

```
$ echo -n "SGVsbG8gd29ybGQh" renc base64:hex
48656c6c6f20776f726c6421
```

Encode a binary file to `base64`:

```
$ cat data.bin | renc :base64
UHJldHR5IGxhbWUgZWFzdGVyIGVnZywgaHVoPw==
```
