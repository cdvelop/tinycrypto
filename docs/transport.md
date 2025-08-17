# Transport API

This section describes the transport layer API, which is responsible for the physical communication with the CryptoAuthentication device.

---

## `Transport` Interface

The `Transport` interface defines the methods that a transport implementation must provide.

```go
type Transport interface {
    Receive(wordAddress byte, data []byte) (err error)
    Send(wordAddress byte, data []byte) (err error)
    WakeUp() (err error)
    Idle() (err error)
}
```

*   **Receive**: Receives data from the device.
*   **Send**: Sends data to the device.
*   **WakeUp**: Wakes up the device.
*   **Idle**: Puts the device in an idle state.

---

## I2C Transport

The library provides a built-in I2C transport implementation.

### `NewI2CTransport(...)`

```go
NewI2CTransport(deviceAddress uint8, bus I2C, baud uint32) Transport
```

Creates a new I2C transport object.

*   **deviceAddress**: The I2C address of the device.
*   **bus**: An implementation of the `I2C` interface.
*   **baud**: The I2C bus speed in Hz.
*   **Returns**: A `Transport` object that can be used with `NewDevice`.

---

### `I2C` Interface

To use the `NewI2CTransport` function, you need to provide an object that implements the `I2C` interface. This interface abstracts the low-level I2C communication.

```go
type I2C interface {
    ReadRegister(addr uint8, r uint8, buf []byte) error
    WriteRegister(addr uint8, r uint8, buf []byte) error
    Tx(addr uint16, w, r []byte) error
    SetBaudRate(br uint32)
}
```

You will need to implement this interface for your specific hardware platform.
