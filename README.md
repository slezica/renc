# renc

Command-line tool to encode and decode streams in a variety of common encodings.

## Installation

Grab a binary from `/bin` and place it in your `$PATH`, or build it yourself with `make build`. You
can use `make build-all` to cross-compile for all platforms.

## Usage

`renc` reads from `stdin`, decodes and/or encodes, and writes to `stdout`. You can specify the encodings
using this syntax:

```bash
renc [decode]:[encode]
```

For example:

```bash
cat ./base64file | renc base64:hex # re-encodes base64 as hex
```

Currently supported encodings are `raw`, `hex`, `base32` and `base64`. If either encoding is
omitted, it defaults to `raw`. For example:

```bash
renc :hex
renc raw:hex # same thing
```
```bash
renc base64:
renc base64:raw # same thing
```

### Examples

Decode hex to raw bytes (no re-encoding):
```bash
$ echo -n "48656c6c6f20776f726c6421" | renc hex:
Hello world!
```

Decode `base64`, then encode in `hex`:

```bash
$ echo -n "SGVsbG8gd29ybGQh" | renc base64:hex
48656c6c6f20776f726c6421
```

Encode a binary file to `base64`:

```bash
$ cat data.bin | renc :base64
UHJldHR5IGxhbWUgZWFzdGVyIGVnZywgaHVoPw==
```
