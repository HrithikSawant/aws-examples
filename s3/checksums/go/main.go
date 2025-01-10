package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// Function to calculate SHA-256 checksum of a file
func calculateSHA256Checksum(filePath string) (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a new SHA-256 hash object
	hash := sha256.New()

	// Copy the file content into the hash function
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	// Return the checksum in hexadecimal format
	checksum := hex.EncodeToString(hash.Sum(nil))
	return checksum, nil
}

func main() {
	// Replace this with the path to your file
	filePath := "./myfile.txt"

	// Calculate SHA-256 checksum
	checksum, err := calculateSHA256Checksum(filePath)
	if err != nil {
		fmt.Println("Error calculating checksum:", err)
		return
	}

	// Print the SHA-256 checksum
	fmt.Println("SHA-256 checksum of the file:", checksum)
}
