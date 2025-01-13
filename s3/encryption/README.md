# S3 Encryption Demo

This README provides a guide for using AWS S3 server-side encryption (SSE) and client-side encryption (SSE-C) to securely upload and manage objects in an S3 bucket. 

## Prerequisites

- **AWS CLI** installed and configured with appropriate permissions.
- Access to a supported AWS region: 
  - **US East (N. Virginia)**
  - **US West (Oregon)**
  - **Europe (Ireland)**
- **KMS Key** for SSE-KMS encryption.

## Steps

### 1. Create an S3 Bucket
```sh
aws s3 mb s3://encryption-fun-hrithik-fg654238
```

### 2. Create a File and Upload to the Bucket
Create a sample file and upload it to the S3 bucket:
```sh
echo "Welcome Hrithik Encryption" > ./hello.txt
aws s3 cp hello.txt s3://encryption-fun-hrithik-fg654238
```

### 3. Create or Retrieve a KMS Key
To use SSE-KMS encryption, you need a KMS Key ID. Follow these steps:

1. **Sign In to the AWS Management Console**
   - Go to the [AWS KMS Console](https://console.aws.amazon.com/kms).

2. **Create a New KMS Key (If Not Already Created)**
   - Navigate to **Customer Managed Keys**.
   - Click **Create Key**.
   - Choose the key type (e.g., **Symmetric**).
   - Provide a **key alias** (a friendly name for your key).
   - Set up key administrators and usage permissions (IAM roles or users who can manage or use the key).
   - Complete the creation process.

3. **Find the Key ID**
   - After the key is created, locate it under the **Customer Managed Keys** section.
   - The **Key ID** will be listed in the details of your key. It will look like:  
     `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`.

4. **Using the Key ID**
   - Use the Key ID in your `--ssekms-key-id` parameter when uploading files with KMS encryption.

Alternatively, use the AWS CLI to retrieve the Key ID:

```sh
aws kms list-keys
```

This will return the Key ARNs of all keys. To get details for a specific key, use:

```sh
aws kms describe-key --key-id <KEY_ARN>
```
Replace `<KEY_ARN>` with the ARN of the key (e.g., `arn:aws:kms:region:account-id:key/key-id`).

### 4. Put Object with Server-Side Encryption using KMS (SSE-KMS)
Use the following command to upload a file with SSE-KMS encryption:
```sh
aws s3api put-object \
--bucket encryption-fun-hrithik-fg654238 \
--key hello.txt \
--body hello.txt \
--server-side-encryption "aws:kms" \
--ssekms-key-id "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
```

### 5. Attempt to Put Object with SSE-C (Failed)
The following command demonstrates an unsuccessful attempt at uploading a file with SSE-C encryption due to an MD5 mismatch:
```sh
export BASE64_ENCODED_KEY=$(openssl rand -base64 32)
echo  $BASE64_ENCODED_KEY

export MD5_VALUE=$(echo $BASE64_ENCODED_KEY | md5sum | awk '{print $1}' | base64 -w0)
echo  $MD5_VALUE

aws s3api put-object \
--bucket encryption-fun-hrithik-fg654238 \
--key hello.txt \
--body hello.txt \
--sse-customer-algorithm AES256 \
--sse-customer-key $BASE64_ENCODED_KEY \
#--sse-customer-key-md5 $MD5_VALUE
```
**Error Output:**
```
An error occurred (InvalidArgument) when calling the PutObject operation: The calculated MD5 hash of the key did not match the hash that was provided.
```

### 6. Successful Put Object with SSE-C
Generate a 256-bit encryption key and upload the file with SSE-C:
```sh
openssl rand -out ssec.key 32

aws s3 cp hello.txt s3://encryption-fun-hrithik-fg654238/hello.txt \
--sse-c AES256 \
--sse-c-key fileb://ssec.key
```

### 7. Download Object Encrypted with SSE-C
Download the file from S3 using the same encryption key:
```sh
aws s3 cp s3://encryption-fun-hrithik-fg654238/hello.txt hello.txt --sse-c AES256 --sse-c-key fileb://ssec.key
```
**Note:** If the encryption key is not supplied during download, the following error will occur:
```
Bad Request
```

## References
- [AWS S3 SSE-C Documentation](https://catalog.us-east-1.prod.workshops.aws/workshops/aad9ff1e-b607-45bc-893f-121ea5224f24/en-US/s3/serverside/ssec)

## Important Notes
- **KMS Key Costs:** KMS keys incur a monthly cost and are only supported in specific regions (US East (N. Virginia), US West (Oregon), Europe (Ireland)).
- **SSE-C:** Ensure that the same key used for encryption is securely stored and provided during file downloads or further operations.

- **Clean-up** Delete the KNS key and bucket