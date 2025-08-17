# tinygo-cryptoauthlib

CryptoAutoLib port for TinyGo.

This library provides a Go interface for the Microchip CryptoAuthentication library, allowing interaction with devices like ATECC608A.

## Basic Usage

To use the library, you first need to create a `Device` object with a `Transport`. The following example shows how to create a device using the I2C transport.

```go
package main

import (
	"machine"
	"time"

	"github.com/tinygo-org/tinygo-cryptoauthlib"
)

// Your I2C implementation
type MyI2C struct {
    // ...
}

func (i *MyI2C) ReadRegister(addr uint8, r uint8, buf []byte) error {
    // ...
}

func (i *MyI2C) WriteRegister(addr uint8, r uint8, buf []byte) error {
    // ...
}

func (i *MyI2C) Tx(addr uint16, w, r []byte) error {
    // ...
}

func (i *MyI2C) SetBaudRate(br uint32) {
    // ...
}


func main() {
    // Configure the I2C bus
	i2c := &MyI2C{}

	// Create a new I2C transport
	transport := cryptoauthlib.NewI2CTransport(0xC0, i2c, 100000)

	// Create a new device
	device := cryptoauthlib.NewDevice(transport)

	// Now you can use the device to interact with the CryptoAuthentication chip
    serialNumber, err := device.ReadSerialNumber()
    if err != nil {
        // Handle error
    }
    println(string(serialNumber[:]))
}
```

## Documentation

The detailed API documentation is organized into the following sections:

*   [Basic API](./docs/basic.md)
*   [Configuration](./docs/config.md)
*   [Key Generation](./docs/keys.md)
*   [Signing and Verification](./docs/sign-verify.md)
*   [Hashing](./docs/hash.md)
*   [ECDH](./docs/ecdh.md)
*   [Transport](./docs/transport.md)
