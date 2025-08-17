# Signing and Verification API

This section describes the API for signing messages and verifying signatures.

---

### `(d *Device) Sign(...)`

```go
Sign(keyId uint16, message []byte) (signature []byte, err error)
```

Signs a 32-byte message using the private key in the specified slot.

*   **keyId**: The slot containing the private key to use for signing.
*   **message**: The 32-byte message to sign.
*   **Returns**: The signature and an error if any.

---

### `(d *Device) SignInternal(...)`

```go
SignInternal(keyId uint16, invalidate bool, fullSerialNumber bool) (signature []byte, err error)
```

Generates a signature using an internally stored digest.

*   **keyId**: The slot containing the private key.
*   **invalidate**: If true, the private key will be invalidated after signing.
*   **fullSerialNumber**: If true, the full serial number of the device will be included in the signature.
*   **Returns**: The signature and an error if any.

---

### `(d *Device) SignExt(...)`

```go
SignExt(mode uint8, keyId uint16) (signature []byte, err error)
```

An extended version of `Sign` that allows for more control over the signing process.

*   **mode**: The signing mode.
*   **keyId**: The slot containing the private key.
*   **Returns**: The signature and an error if any.

---

### `(d *Device) Verify(...)`

```go
Verify(mode uint8, keyId uint16, signature []byte, pubKey []byte, otherData []byte) (mac []byte, err error)
```

Verifies a signature with a public key.

*   **mode**: The verification mode.
*   **keyId**: The slot of the key to use for verification.
*   **signature**: The signature to verify.
*   **pubKey**: The public key to use for verification (for external verification).
*   **otherData**: Other data for the verification process.
*   **Returns**: A MAC and an error if any. A `StatusCheckmacVerifyFailed` error indicates that the signature is invalid.

---

### `(d *Device) VerifyExternal(...)`

```go
VerifyExternal(message, signature, pubKey []byte) (verified bool, err error)
```

Verifies a signature of a message using an external public key.

*   **message**: The message that was signed.
*   **signature**: The signature to verify.
*   **pubKey**: The public key to use for verification.
*   **Returns**: `true` if the signature is valid, `false` otherwise, and an error if any.
