package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"flag"
	"fmt"
	"math/rand"
)

// genKey generates a 128-bit AES key from a given seed
func genKey(seed int64) []byte {
	rand.Seed(seed)
	key := make([]byte, 16) // AES-128
	for i := range key {
		key[i] = byte(rand.Intn(254) + 1)
	}
	return key
}

// decrypt decrypts the encrypted data using the generated key and returns the original text
func decrypt(seed int64, encryptedBase64 string) (string, error) {
	// Generate the encryption key using the same seed
	key := genKey(seed)

	// Decode the base64-encoded data
	encryptedData, err := base64.URLEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return "", fmt.Errorf("base64 decode: %w", err)
	}

	// The first 16 bytes should be the IV
	iv := encryptedData[:aes.BlockSize]
	encryptedText := encryptedData[aes.BlockSize:]

	// Create a new AES cipher using the generated key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("new cipher: %w", err)
	}

	// Decrypt the data using CFB mode
	stream := cipher.NewCFBDecrypter(block, iv)
	decrypted := make([]byte, len(encryptedText))
	stream.XORKeyStream(decrypted, encryptedText)

	return string(decrypted), nil
}

func main() {
	// Define command-line flags
	seedPtr := flag.Int64("seed", 0, "Seed used to generate the encryption key")
	encryptedBase64Ptr := flag.String("data", "", "Base64-encoded encrypted data to decrypt")

	// Parse the flags
	flag.Parse()

	// Validate inputs
	if *seedPtr == 0 || *encryptedBase64Ptr == "" {
		fmt.Println("Usage: decrypt -seed=<seed> -data=\"<encrypted data>\"")
		return
	}

	// Decrypt the text using provided command-line arguments
	decryptedText, err := decrypt(*seedPtr, *encryptedBase64Ptr)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}

	fmt.Println("Decrypted text:", decryptedText)
}
