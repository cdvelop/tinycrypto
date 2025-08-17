# Key Generation API

This section describes the API for generating cryptographic keys on the device.

---

### `(d *Device) GenKey(...)`

```go
GenKey(keyId uint16) (pubKey []byte, err error)
```

Generates a new private key in the specified slot and returns the corresponding public key.

*   **keyId**: The slot where the private key will be generated and stored.
*   **Returns**: The generated public key and an error if any.

---

### `(d *Device) GenKeyMac()`

```go
GenKeyMac() (pubKey []byte, mac []byte, err error)
```

Generates a new key and returns the public key and a MAC.

*   **Returns**: The public key, a MAC, and an error if any.

---

### `(d *Device) GetPublicKey(...)`

```go
GetPublicKey(keyId uint16) (pubkey []byte, err error)
```

Retrieves the public key from a given slot without generating a new key. This can be used to get the public key of a key that was previously generated.

*   **keyId**: The slot containing the key.
*   **Returns**: The public key and an error if any.

---

### `(d *Device) GenKeyExt(...)`

```go
GenKeyExt(mode uint8, keyId uint16, otherData []byte) (pubKey []byte, err error)
```

An extended version of `GenKey` that allows for more control over the key generation process.

*   **mode**: The key generation mode.
*   **keyId**: The slot for the key.
*   **otherData**: Additional data for the key generation.
*   **Returns**: The public key and an error if any.
