//go:build !wasm

package tinycrypto

import "crypto/rand"

func readRandom(b []byte) (err error) {
	_, err = rand.Read(b)
	return
}
