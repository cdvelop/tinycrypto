# Hashing API

This section describes the API for hashing data using SHA256 and HMAC.

---

### `(d *Device) SHA256(...)`

```go
SHA256(input []byte) (output []byte, err error)
```

Computes the SHA256 hash of the input data.

*   **input**: The data to hash.
*   **Returns**: The 32-byte SHA256 hash and an error if any.

---

### `(d *Device) HMAC(...)`

```go
HMAC(input []byte, keySlot uint16, target uint8) (output []byte, err error)
```

Computes the HMAC-SHA256 of the input data using a key from a slot.

*   **input**: The data to hash.
*   **keySlot**: The slot containing the key to use for the HMAC calculation.
*   **target**: The target for the HMAC output.
*   **Returns**: The 32-byte HMAC and an error if any.

---

### `(d *Device) ShaExt(...)`

```go
ShaExt(mode uint8, length uint16, message []byte) (data []byte, err error)
```

An extended version of the SHA command that allows for more control over the hashing process, including starting, updating, and ending a hash calculation.

*   **mode**: The SHA mode (e.g., `SHA_MODE_SHA256_START`, `SHA_MODE_SHA256_UPDATE`, `SHA_MODE_SHA256_END`).
*   **length**: The length of the message.
*   **message**: The message data.
*   **Returns**: The result of the SHA operation and an error if any.
