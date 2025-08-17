# TinyCrypto

A lightweight Go library for cryptographic operations, designed for WebAssembly and small devices using TinyGo.

## Features

- **Simple API:** Easy-to-use API for symmetric and asymmetric encryption, and digital signatures.
- **TinyGo Optimized:** Designed to produce small binaries when compiled with TinyGo.
- **WebAssembly Ready:** Can be used in browser environments.
- **Zero Dependencies on Go Standard Library:** Uses `github.com/cdvelop/tinystring` for string, number, and error handling to minimize binary size.

## Basic Usage

To use the library, create a new `TinyCrypto` engine and call its methods:

```go
package main

import (
	"fmt"

	"github.com/waj334/tinygo-cryptoauthlib"
)

func main() {
	// Create a new crypto engine
	crypto := tinycrypto.New()

	// Symmetric encryption
	key := make([]byte, 32) // AES-256 key
	plaintext := []byte("hello world")
	ciphertext, err := crypto.Encrypt(plaintext, key)
	if err != nil {
		panic(err)
	}
	decrypted, err := crypto.Decrypt(ciphertext, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("Symmetric decrypted:", string(decrypted))

	// Asymmetric signatures
	pub, priv, err := crypto.GenerateKeyPair()
	if err != nil {
		panic(err)
	}
	message := []byte("this is a test message")
	signature, err := crypto.Sign(message, priv)
	if err != nil {
		panic(err)
	}
	ok, err := crypto.Verify(message, signature, pub)
	if err != nil {
		panic(err)
	}
	fmt.Println("Signature verified:", ok)
}
```

## Documentation

The detailed API documentation is organized into the following sections:

- [Symmetric Encryption](./docs/symmetric.md)
- [Asymmetric Encryption](./docs/asymmetric.md)
- [Digital Signatures](./docs/signatures.md)
