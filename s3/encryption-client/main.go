package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// Define S3 bucket and object key
	bucket := "encryption-client-fun-hrithik-fg23y54035"
	objectKey := "hello.txt"

	// Generate an RSA key
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Fatalf("Failed to generate RSA key: %v", err)
	}

	// The plaintext data to upload
	plaintext := []byte("handshake")

	// Encrypt the plaintext locally using RSA for the content encryption key
	contentEncryptionKey := make([]byte, 32) // AES-256 requires 32 bytes key
	_, err = rand.Read(contentEncryptionKey)
	if err != nil {
		log.Fatalf("Failed to generate content encryption key: %v", err)
	}

	encryptedKey, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, &key.PublicKey, contentEncryptionKey, nil)
	if err != nil {
		log.Fatalf("Failed to encrypt content encryption key: %v", err)
	}

	// Encrypt the data with AES-GCM using the content encryption key
	ciphertext, err := encryptWithAESGCM(contentEncryptionKey, plaintext)
	if err != nil {
		log.Fatalf("Failed to encrypt data: %v", err)
	}

	// Create an AWS session and S3 client
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("YOUR REGION"), // Update region
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}
	s3Client := s3.New(sess)

	// Upload the encrypted object
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
		Body:   bytes.NewReader(ciphertext),
		Metadata: map[string]*string{
			"x-amz-meta-encrypted-key": aws.String(base64.StdEncoding.EncodeToString(encryptedKey)),
			"x-amz-meta-algorithm":     aws.String("AES/GCM/NoPadding"),
		},
	})
	if err != nil {
		log.Fatalf("Failed to put object: %v", err)
	}
	fmt.Println("PUT: Encrypted object uploaded")

	// Retrieve the encrypted object
	getResp, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Fatalf("Failed to get object: %v", err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(getResp.Body)
	encryptedData := buf.Bytes()
	fmt.Println("GET WITH KEY: Encrypted data retrieved")

	// Decrypt the content encryption key
	decryptedKey, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, key, encryptedKey, nil)
	if err != nil {
		log.Fatalf("Failed to decrypt content encryption key: %v", err)
	}

	// Decrypt the data
	decryptedData, err := decryptWithAESGCM(decryptedKey, encryptedData)
	if err != nil {
		log.Fatalf("Failed to decrypt data: %v", err)
	}
	fmt.Println("Decrypted data:", string(decryptedData))
}

// encryptWithAESGCM encrypts data using AES-GCM
func encryptWithAESGCM(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, 12) // AES-GCM nonce size
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// decryptWithAESGCM decrypts data using AES-GCM
func decryptWithAESGCM(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
