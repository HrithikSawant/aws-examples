# S3 Metadata Demo

This guide demonstrates how to use **S3 object metadata** with the AWS CLI by executing script. Follow the steps below to interact with S3 buckets and objects, create metadata, and retrieve object metadata.

## Prerequisites

- AWS CLI installed and configured.
- `metadata` script.

## Steps to Demo Metadata Commands

### Step 1: Create an S3 Bucket

To begin, create a new S3 bucket. Replace the `random-bucket-name` with a unique name for your bucket.

```bash
./metadata create-bucket random-bucket-name
```

### Step 2: Upload a File to the S3 Bucket with Metadata
Upload a file to the newly created bucket and add metadata (e.g., Planet=Mars):

```bash
./metadata upload-file random-bucket-name ./myfile.txt Planet=Mars
```

#### output

```bash
{
    "ETag": "\"311c0df03d20d005c3a0a2ecb1353a46\"",
    "ServerSideEncryption": "AES256"
}
```

### Step 3: Retrieve Metadata of the Uploaded File
To retrieve the metadata of a file in the bucket, run the following command:

```bash
./metadata get-metadata random-bucket-name myfile.txt
```

Example output:

```bash
{
    "AcceptRanges": "bytes",
    "LastModified": "2025-01-09T06:46:16+00:00",
    "ContentLength": 29,
    "ETag": "\"311c0df03d20d005c3a0a2ecb1353a46\"",
    "ContentType": "binary/octet-stream",
    "ServerSideEncryption": "AES256",
    "Metadata": {
        "planet": "Mars"
    }
}
```

### Step 4: Clean Up (Delete the File or Bucket)
If you want to delete the uploaded file, run:

```bash
./metadata cleanup random-bucket-name myfile.txt
```

Example output:

```bash
Deleting file: myfile.txt from bucket: random-bucket-name
delete: s3://random-bucket-name/myfile.txt
```

To delete the entire bucket, use the following command:

```bash
./metadata cleanup random-bucket-name
```

Example output:

```bash
Deleting bucket: random-bucket-name
remove_bucket: random-bucket-name
```

### Step 5: Verify Cleanup
You can verify the bucket and file deletion by listing all buckets:

```bash
aws s3 ls
```
