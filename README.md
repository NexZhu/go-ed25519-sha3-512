# go-ed25519-sha3-512

This module is a fork of [crypto/ed25519](https://pkg.go.dev/crypto/ed25519) package
modified to use **SHA3-512** hash function instead of **SHA-512** used in the original code.

## Usage

See [godoc](https://pkg.go.dev/github.com/crpt/go-ed25519-sha3-512).

If you're using [crypto/ed25519](https://pkg.go.dev/crypto/ed25519) package, just replace:

```go
import "crypto/ed25519"
```

with:

```go
import ed25519 "github.com/crpt/go-ed25519-sha3-512"
```

and you're good to go.

## TODO

Find new test data to replace `testdata/sign.input.gz` in `ed25519_test.go` > `func TestGolden`
