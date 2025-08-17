# Basic API

This section describes the basic functions for interacting with the CryptoAuthentication device.

---

### `NewDevice(transport Transport) *Device`

Creates a new `Device` object.

*   **transport**: An implementation of the `Transport` interface (e.g., `NewI2CTransport`).
*   **Returns**: A pointer to a `Device` object.

---

### `(d *Device) ReadZone(...)`

```go
ReadZone(zone uint8, slot uint16, blockNo uint8, offset uint8, mode IOMode) (data []byte, err error)
```

Reads a block or a word from a given zone.

*   **zone**: The zone to read from (e.g., `ATCA_ZONE_CONFIG`, `ATCA_ZONE_DATA`, `ATCA_ZONE_OTP`).
*   **slot**: The slot to read from.
*   **blockNo**: The block number within the zone.
*   **offset**: The offset within the block.
*   **mode**: `ReadBlock` (32 bytes) or `ReadWord` (4 bytes).
*   **Returns**: The data read from the device and an error if any.

---

### `(d *Device) ReadBytesZone(...)`

```go
ReadBytesZone(zone uint8, slot uint16, offset int, data []byte) (err error)
```

Reads a specified number of bytes from a zone.

*   **zone**: The zone to read from.
*   **slot**: The slot to read from.
*   **offset**: The starting offset to read from.
*   **data**: A byte slice to store the read data. The length of the slice determines the number of bytes to read.
*   **Returns**: An error if any.

---

### `(d *Device) ReadSerialNumber()`

```go
ReadSerialNumber() (result [9]byte, err error)
```

Reads the 9-byte serial number of the device.

*   **Returns**: The serial number as a 9-byte array and an error if any.

---

### `(d *Device) WriteZone(...)`

```go
WriteZone(zone uint8, slot uint16, block uint8, offset uint8, data []byte) (err error)
```

Writes a block (32 bytes) or a word (4 bytes) to a given zone.

*   **zone**: The zone to write to.
*   **slot**: The slot to write to.
*   **block**: The block number to write to.
*   **offset**: The offset within the block.
*   **data**: The data to write. Must be 4 or 32 bytes long.
*   **Returns**: An error if any.

---

### `(d *Device) WriteBytesZone(...)`

```go
WriteBytesZone(zone uint8, slot uint16, offset int, data []byte) (err error)
```

Writes a specified number of bytes to a zone. The length of the data must be a multiple of 4.

*   **zone**: The zone to write to.
*   **slot**: The slot to write to.
*   **offset**: The starting offset to write to.
*   **data**: The data to write.
*   **Returns**: An error if any.

---

### `(d *Device) ReadPublicKey(...)`

```go
ReadPublicKey(slot uint16) (pubkey []byte, err error)
```

Reads a public key from a given slot.

*   **slot**: The slot containing the public key (8-15).
*   **Returns**: The public key and an error if any.

---

### `(d *Device) WritePubKey(...)`

```go
WritePubKey(slot uint16, pubKey []byte) (err error)
```

Writes a public key to a given slot.

*   **slot**: The slot to write the public key to.
*   **pubKey**: The 64-byte public key to write.
*   **Returns**: An error if any.

---

### `(d *Device) ReadSignature(...)`

```go
ReadSignature(slot uint16) (signature []byte, err error)
```

Reads a signature from a given slot.

*   **slot**: The slot containing the signature (8-15).
*   **Returns**: The signature and an error if any.

---

### `(d *Device) ReadEncrypted(...)`

```go
ReadEncrypted(keyId uint16, block uint8, encKey []byte, encKeyId uint16, num []byte) (data []byte, err error)
```

Reads data from a slot that has been encrypted with a key from another slot.

*   **keyId**: The slot to read from.
*   **block**: The block to read from.
*   **encKey**: The encryption key.
*   **encKeyId**: The slot of the encryption key.
*   **num**: A 20-byte nonce.
*   **Returns**: The decrypted data and an error if any.

---

### `(d *Device) WriteEnc(...)`

```go
WriteEnc(keyId uint16, block uint8, data []byte, encKey []byte, encKeyId uint16, num []byte) (err error)
```

Writes data to a slot after encrypting it with a key from another slot.

*   **keyId**: The slot to write to.
*   **block**: The block to write to.
*   **data**: The data to encrypt and write.
*   **encKey**: The encryption key.
*   **encKeyId**: The slot of the encryption key.
*   **num**: A 20-byte nonce.
*   **Returns**: An error if any.
