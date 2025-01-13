# S3 Encryption Client Demo

This README provides a guide for running a client-side encryption demo with AWS S3 using an SDK written in Go. The script encrypts data locally, uploads it to S3, retrieves it, and decrypts it.

## Prerequisites

- **AWS CLI** installed and configured.
- **Go Programming Language** installed.
- An AWS S3 bucket.

## Steps

### 1. Create a Bucket

Run the following command to create a new S3 bucket:
```sh
aws s3 mb s3://encryption-client-fun-hrithik-fg23y54035
```


### 2. Run the SDK Script

#### Code Updation required before running
#### Main Script
- Update bucket name
- Update region as per your requirement

Navigate to the directory containing the Go script and run the following commands:
```sh
go mod tidy
go run main.go
```

### 3. Expected Output

When the script runs successfully, you should see the following output:
```
PUT: Encrypted object uploaded
GET WITH KEY: Encrypted data retrieved
Decrypted data: handshake
```

### 4. Cleanup

To remove the object and delete the bucket, run the following commands:
```sh
aws s3 rm s3://encryption-client-fun-hrithik-fg23y54035/hello.txt
aws s3 rb s3://encryption-client-fun-hrithik-fg23y54035
```

## How It Works

The Go script performs the following steps:

1. **Generate an RSA Key Pair**
   - An RSA key pair is created for encrypting the content encryption key (CEK).

2. **Generate a Content Encryption Key (CEK)**
   - A random 256-bit key is generated for AES-GCM encryption.

3. **Encrypt the Content Encryption Key**
   - The CEK is encrypted using the RSA public key.

4. **Encrypt the Data**
   - The plaintext is encrypted using AES-GCM and the CEK.

5. **Upload Encrypted Data to S3**
   - The encrypted data is uploaded to the S3 bucket along with metadata for the encrypted CEK and algorithm.

6. **Retrieve and Decrypt Data**
   - The script retrieves the encrypted object, decrypts the CEK using the RSA private key, and decrypts the data using the CEK.

