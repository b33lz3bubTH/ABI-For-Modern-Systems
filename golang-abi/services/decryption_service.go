package services

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/sha512"
    "encoding/base64"
    "errors"
    "computations_crypto/types"
)

type DecryptionService struct{}

// DecryptData decrypts the input and verifies the SHA-512 hash
func (d *DecryptionService) DecryptData(encData *types.EncryptedData) (*types.DecryptedData, error) {
    // Decode AES key and cipher text
    key, err := base64.StdEncoding.DecodeString(encData.AESKey)
    if err != nil {
        return nil, err
    }

    cipherText, err := base64.StdEncoding.DecodeString(encData.DataEncrypted)
    if err != nil {
        return nil, err
    }

    // Create AES cipher
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    // Decrypt the data
    iv := cipherText[:aes.BlockSize]
    cipherText = cipherText[aes.BlockSize:]
    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(cipherText, cipherText)

    originalData := string(cipherText)

    // Verify SHA-512 hash
    hash := sha512.Sum512([]byte(originalData))
    if base64.StdEncoding.EncodeToString(hash[:]) != encData.SHA512Hash {
        return nil, errors.New("hash verification failed")
    }

    return &types.DecryptedData{
        OriginalData: originalData,
        IsValid:      true,
    }, nil
}