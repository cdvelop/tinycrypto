# ECDH API

This section describes the API for performing Elliptic Curve Diffie-Hellman (ECDH) key agreement.

---

### `(d *Device) Ecdh(...)`

```go
Ecdh(keyId uint16, pubKey []byte) (pms []byte, err error)
```

Performs an ECDH key agreement with the given public key and the private key in the specified slot.

*   **keyId**: The slot containing the private key to use for the ECDH calculation.
*   **pubKey**: The public key of the other party.
*   **Returns**: The premaster secret (PMS) and an error if any.

---

### `(d *Device) EcdhEncrypted(...)`

```go
EcdhEncrypted(keyId uint16, pubKey []byte, readKey []byte, readKeyId uint16, num []byte) (pms []byte, err error)
```

Performs an ECDH key agreement and returns the premaster secret encrypted with a read key.

*   **keyId**: The slot containing the private key.
*   **pubKey**: The public key of the other party.
*   **readKey**: The key to encrypt the PMS with.
*   **readKeyId**: The slot of the read key.
*   **num**: A 20-byte nonce.
*   **Returns**: The encrypted premaster secret and an error if any.

---

### `(d *Device) EcdhExt(...)`

```go
EcdhExt(mode uint8, keyId uint16, pubKey []byte) (pms []byte, nonce []byte, err error)
```

An extended version of the ECDH command that allows for more control over the key agreement process.

*   **mode**: The ECDH mode.
*   **keyId**: The slot containing the private key.
*   **pubKey**: The public key of the other party.
*   **Returns**: The premaster secret, a nonce, and an error if any.
