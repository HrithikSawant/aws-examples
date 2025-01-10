# Checksum Demo

This repository demonstrates how to calculate a file's checksum using:
1. A Bash script that calculates and uploads the checksum metadata to an AWS S3 bucket.
2. A Go program that calculates the SHA-256 checksum of a file.

---

## 1. Create an S3 Bucket
To begin, create an S3 bucket where files and their checksum metadata will be stored:
```bash
aws s3 mb s3://checksum-examples-ab-2342
```

## 2. Create a Test File
Create a sample file for which we will calculate the checksum:
```bash
echo "Hello checksum example" > myfile.txt
```

---

## 3. Calculate MD5 Checksum and Upload Manually
To manually calculate the MD5 checksum and upload a file:
1. Calculate the checksum:
   ```bash
   md5sum myfile.txt
   # Output: 2c02890b2273a66964217601f550241c  myfile.txt
   ```
2. Upload the file to S3:
   ```bash
   aws s3 cp myfile.txt s3://checksum-examples-ab-2342
   ```
3. Verify the ETag (MD5 checksum):
   ```bash
   aws s3api head-object --bucket checksum-examples-ab-2342 --key myfile.txt
   ```
   Example Output:
   ```json
   {
       "AcceptRanges": "bytes",
       "LastModified": "2024-12-23T05:52:20+00:00",
       "ContentLength": 23,
       "ETag": "\"2c02890b2273a66964217601f550241c\"",
       "ContentType": "text/plain",
       "ServerSideEncryption": "AES256",
       "Metadata": {}
   }
   ```

---

## 4. Go Program: `checksum-calculator.go`
This Go program calculates the SHA-256 checksum of a given file.

### Prerequisites
- Install [Go](https://golang.org/doc/install).

### How to Use
1. Save the Go code to a file, e.g., `checksum-calculator.go`.
2. Create a test file to calculate its checksum:
   ```bash
   echo "Hello checksum example" > myfile.txt
   ```
3. Run the Go program:
   ```bash
   go run checksum-calculator.go
   ```
4. The output will display the SHA-256 checksum of the file.

---

## 5. Use the Bash Script for Advanced Checksum Uploads
To upload a file with additional checksum metadata, run the provided Bash script:
```bash
./put-object-checksum checksum-examples-ab-2342 ./myfile.txt
```

---

## Example Output
### Bash Script Output:
```plaintext
== Starting put-object-checksum script ==
Uploading file 'myfile.txt' to bucket 'checksum-examples-ab-2342'...
SHA256 Checksum: <calculated-SHA256-value>
MD5 Checksum: <calculated-MD5-value>
File uploaded successfully with metadata.
Fetching metadata for object 'myfile.txt'...
<metadata details>
== Finished put-object-checksum script ==
```

### Go Program Output:
```plaintext
SHA-256 checksum of the file: <calculated-SHA256-value>
```

---

## Notes
- Ensure the AWS CLI has sufficient permissions to upload and retrieve object metadata from the specified bucket.
- Modify the `filePath` variable in the Go program to test with different files.

