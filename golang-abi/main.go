// /main.go
package main


import (
    "C"
	"fmt"
    "encoding/json"
    "computations_crypto/services" // Update with the correct module path
    "computations_crypto/types"
)

// Global services
var encryptionService = &services.EncryptionService{}
var decryptionService = &services.DecryptionService{}

//export encrypt
func encrypt(input *C.char) *C.char {
    goInput := C.GoString(input)

    // Encrypt the input JSON
    encryptedData, err := encryptionService.EncryptData(goInput)
    if err != nil {
        return C.CString("Error: " + err.Error())
    }

    // Marshal the result into JSON
    result, err := json.Marshal(encryptedData)
    if err != nil {
        return C.CString("Error: " + err.Error())
    }

    return C.CString(string(result))
}

//export decrypt
func decrypt(input *C.char) *C.char {
    goInput := C.GoString(input)

    // Parse the input JSON into EncryptedData struct
    var encData types.EncryptedData
    if err := json.Unmarshal([]byte(goInput), &encData); err != nil {
        return C.CString("Error: " + err.Error())
    }

    // Decrypt the data
    decryptedData, err := decryptionService.DecryptData(&encData)
    if err != nil {
        return C.CString("Error while decrypting: " + err.Error())
    }

    // Marshal the result into JSON
    result, err := json.Marshal(decryptedData)
    if err != nil {
        return C.CString("Error: " + err.Error())
    }

    return C.CString(string(result))
}

func main() {
    // Create service instances locally within the main function
    encryptionService := &services.EncryptionService{}
    decryptionService := &services.DecryptionService{}

    // Simulate JSON input
    inputJSON := `{
        "name": "Sample Maksudi Product, gustir gud",
        "price": 2500,
        "category": "Electronics"
    }`

    // Step 1: Encrypt the input JSON
    fmt.Println("Original Data:", inputJSON)
    encryptedData, err := encryptionService.EncryptData(inputJSON)
    if err != nil {
        fmt.Println("Encryption Error:", err)
        return
    }

    // Convert encrypted data to JSON string for better visualization
    encDataJSON, err := json.MarshalIndent(encryptedData, "", "  ")
    if err != nil {
        fmt.Println("Error converting encrypted data to JSON:", err)
        return
    }
    fmt.Println("Encrypted Data:", string(encDataJSON))


    // Step 2: Decrypt the data
    decryptedData, err := decryptionService.DecryptData(encryptedData)
    if err != nil {
        fmt.Println("Decryption Error:", err)
        return
    }

    // Convert decrypted data to JSON string for better visualization
    decDataJSON, err := json.MarshalIndent(decryptedData, "", "  ")
    if err != nil {
        fmt.Println("Error converting decrypted data to JSON:", err)
        return
    }

    // Step 3: Output decrypted data
    fmt.Println("Decrypted Data:", string(decDataJSON))




	// Checking the decrypt function using C-style interaction
	const decrypting_raw_data = `{
		"data_encrypted": "NnDfr6VTCn88lBGNMxbS2Df9TQMSbHNzTEYxdo+OR8hSSQKIP+gJXmKXJYc8By7KKnuyLHDyiRaR6MJ99rGWk5Fx7N3UQ3x20kmvG6NsQYbu+n5wzfxYuMiUIVAaUg+4iCt/GuyNE36BI/wG1LsvAOM4",
		"aes_key": "2ivsryxAqFYX+i163pvx/FPgU7nWPU5i6b5HC8WQX6M=",
		"sha512": "3UCl+9TI1fLrN5qD+7W9vuc9r1Ke/yD5XsrChKEVSyufQILQMSgmwbcN5hw6OKxLvH+fLlIvho6Dnnz06z4sHg=="
	}`

	// Convert Go string to C string
	cInput := C.CString(decrypting_raw_data)

	// Call the decrypt function with the C string
	decrypted := decrypt(cInput)

	// Print the result of decryption
	fmt.Println("Decrypted result:", C.GoString(decrypted))
}
