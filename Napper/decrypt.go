/*
AES Decryption Tool for Napper Challenge

Author: Lukas Johannes Möller
Date: 24.02.2024
License: MIT License

This Go script is designed to decrypt data encrypted with AES-128 in CFB mode. It is specifically tailored for use in the Hack The Box machine "Napper" challenge, providing a means to decrypt data using a seed-based key generation approach. The script accepts two crucial pieces of information via command-line arguments: a seed used for key generation and the base64-encoded encrypted data.

The decryption process involves the following steps:
1. Generating a 128-bit AES key from the provided seed.
2. Decoding the base64-encoded encrypted data to retrieve the ciphertext.
3. Using the AES key and Cipher Feedback (CFB) mode to decrypt the ciphertext and obtain the original plaintext.

Prerequisites:
- Go (Golang) environment set up on the machine where the script will be run.

Usage:
The script is executed from the command line, where the seed and encrypted data are passed as arguments:
    go run decrypt.go -seed=<seed> -data="<base64-encoded-data>"

Example:
    go run decrypt.go -seed=46385390 -data="tbjZvSCUhZtSmOqEYO1TFmX-ibTWLnMJc6CQJHZ_aM6alBTptvEaiMEvjv_Jfx33T7spOEMKOXg="

This script serves as an educational tool for understanding AES decryption and should be used in accordance with ethical guidelines and applicable laws.

*/
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
