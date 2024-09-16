// /types/data_types.go
package types

type EncryptedData struct {
    DataEncrypted string `json:"data_encrypted"`
    AESKey        string `json:"aes_key"`
    SHA512Hash    string `json:"sha512"`
}

type DecryptedData struct {
    OriginalData string `json:"original_data"`
    IsValid      bool   `json:"is_valid"`
}