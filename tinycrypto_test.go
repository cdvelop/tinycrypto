package tinycrypto

import (
	"bytes"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	key := make([]byte, 32)
	if err := readRandom(key); err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	plaintext := []byte("hello world")

	ciphertext, err := Encrypt(plaintext, key)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	decrypted, err := Decrypt(ciphertext, key)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("expected %s, got %s", plaintext, decrypted)
	}
}

func TestKeyParsingErrors(t *testing.T) {
	// Test Sign with invalid private key
	_, err := Sign([]byte("message"), []byte("invalid key"))
	if err == nil {
		t.Error("expected error for invalid private key, got nil")
	}

	// Test Verify with invalid public key
	_, err = Verify([]byte("message"), []byte("signature"), []byte("invalid key"))
	if err == nil {
		t.Error("expected error for invalid public key, got nil")
	}

	// Test EncryptAsymmetric with invalid public key
	_, err = EncryptAsymmetric([]byte("plaintext"), []byte("invalid key"))
	if err == nil {
		t.Error("expected error for invalid public key, got nil")
	}

	// Test DecryptAsymmetric with invalid private key
	_, err = DecryptAsymmetric([]byte("ciphertext"), []byte("invalid key"))
	if err == nil {
		t.Error("expected error for invalid private key, got nil")
	}
}

func TestEncryptDecryptError(t *testing.T) {
	// Test with wrong key size
	key := make([]byte, 16)
	plaintext := []byte("hello world")
	_, err := Encrypt(plaintext, key)
	if err == nil {
		t.Error("expected error for wrong key size, got nil")
	}

	// Test with corrupted ciphertext
	realKey := make([]byte, 32)
	if err := readRandom(realKey); err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}
	ciphertext, _ := Encrypt(plaintext, realKey)
	ciphertext[0] ^= 0xff // corrupt the nonce
	_, err = Decrypt(ciphertext, realKey)
	if err == nil {
		t.Error("expected error for corrupted ciphertext, got nil")
	}
}

func TestGenerateKeyPair(t *testing.T) {
	pub, priv, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("GenerateKeyPair failed: %v", err)
	}

	if len(pub) == 0 {
		t.Error("public key is empty")
	}

	if len(priv) == 0 {
		t.Error("private key is empty")
	}
}

func TestSignVerify(t *testing.T) {
	pub, priv, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("GenerateKeyPair failed: %v", err)
	}

	message := []byte("this is a test message")
	signature, err := Sign(message, priv)
	if err != nil {
		t.Fatalf("Sign failed: %v", err)
	}

	ok, err := Verify(message, signature, pub)
	if err != nil {
		t.Fatalf("Verify failed: %v", err)
	}

	if !ok {
		t.Error("signature verification failed")
	}
}

func TestSignVerifyError(t *testing.T) {
	pub, priv, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("GenerateKeyPair failed: %v", err)
	}

	// Test with wrong key
	wrongPub, _, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("GenerateKeyPair failed: %v", err)
	}

	message := []byte("this is a test message")
	signature, err := Sign(message, priv)
	if err != nil {
		t.Fatalf("Sign failed: %v", err)
	}

	ok, err := Verify(message, signature, wrongPub)
	if err != nil {
		t.Fatalf("Verify failed: %v", err)
	}
	if ok {
		t.Error("expected signature verification to fail with wrong key")
	}

	// Test with corrupted signature
	signature[0] ^= 0xff
	ok, err = Verify(message, signature, pub)
	if err != nil {
		t.Fatalf("Verify failed: %v", err)
	}
	if ok {
		t.Error("expected signature verification to fail with corrupted signature")
	}
}

func TestEncryptDecryptAsymmetricError(t *testing.T) {
	pub, priv, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("GenerateKeyPair failed: %v", err)
	}

	// Test with wrong key
	_, wrongPriv, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("GenerateKeyPair failed: %v", err)
	}

	plaintext := []byte("hello asymmetric world")
	ciphertext, err := EncryptAsymmetric(plaintext, pub)
	if err != nil {
		t.Fatalf("EncryptAsymmetric failed: %v", err)
	}

	_, err = DecryptAsymmetric(ciphertext, wrongPriv)
	if err == nil {
		t.Error("expected error for wrong private key, got nil")
	}

	// Test with corrupted ciphertext
	ciphertext[0] ^= 0xff
	_, err = DecryptAsymmetric(ciphertext, priv)
	if err == nil {
		t.Error("expected error for corrupted ciphertext, got nil")
	}
}

func TestEncryptDecryptAsymmetric(t *testing.T) {
	pub, priv, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("GenerateKeyPair failed: %v", err)
	}

	plaintext := []byte("hello asymmetric world")

	ciphertext, err := EncryptAsymmetric(plaintext, pub)
	if err != nil {
		t.Fatalf("EncryptAsymmetric failed: %v", err)
	}

	decrypted, err := DecryptAsymmetric(ciphertext, priv)
	if err != nil {
		t.Fatalf("DecryptAsymmetric failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("expected %s, got %s", plaintext, decrypted)
	}
}
