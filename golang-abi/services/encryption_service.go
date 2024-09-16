// /services/encryption_service.go
package services

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/sha512"
    "encoding/base64"
    "io"
    "computations_crypto/types"
)

type EncryptionService struct{}

// EncryptData encrypts the input JSON with AES-256 and generates a SHA-512 hash
func (e *EncryptionService) EncryptData(input string) (*types.EncryptedData, error) {
    // Generate a random AES key (32 bytes for AES-256)
    key := make([]byte, 32)
    if _, err := rand.Read(key); err != nil {
        return nil, err
    }

    // Create AES cipher
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    // Encrypt the input data
    cipherText := make([]byte, aes.BlockSize+len(input))
    iv := cipherText[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(input))

    // Generate SHA-512 hash of the original data
    hash := sha512.Sum512([]byte(input))

    return &types.EncryptedData{
        DataEncrypted: base64.StdEncoding.EncodeToString(cipherText),
        AESKey:        base64.StdEncoding.EncodeToString(key),
        SHA512Hash:    base64.StdEncoding.EncodeToString(hash[:]),
    }, nil
}