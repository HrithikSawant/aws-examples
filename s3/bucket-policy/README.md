
# S3 Bucket Policy Demo

This guide demonstrates how to add an S3 bucket policy using the AWS CLI and test cross-account access.

---

## Steps

### 1. Create a Bucket
Use the following command to create an S3 bucket:

```bash
aws s3 mb s3://bucket-policy-example-x7y9z3a
```

### 2. Create and Attach the Bucket Policy
Add a bucket policy to allow cross-account access. Replace `policy.json` with the appropriate policy file:

```bash
aws s3api put-bucket-policy --bucket bucket-policy-example-x7y9z3a --policy file://policy.json
```

### 3. Access the Bucket from Another AWS Account
Perform actions such as uploading and listing objects in the bucket from the second account.

#### a. Upload a File
Create a sample file and upload it to the bucket:

```bash
touch bootcamp.txt
aws s3 cp bootcamp.txt s3://bucket-policy-example-x7y9z3a
```

#### b. List the Contents of the Bucket
Verify the uploaded file by listing the bucket contents:

```bash
aws s3 ls s3://bucket-policy-example-x7y9z3a
```

### 4. Cleanup
Remove the test file and delete the bucket to clean up resources:

#### a. Remove the Test File
```bash
aws s3 rm s3://bucket-policy-example-x7y9z3a/bootcamp.txt
```

#### b. Delete the Bucket
```bash
aws s3 rb s3://bucket-policy-example-x7y9z3a
```

---

## Notes
- Ensure the `policy.json` file contains the correct bucket policy to allow access for the specified IAM user or account.
- Verify that the appropriate permissions are granted to the IAM user in the other AWS account.

This demo showcases a simple cross-account access setup with an S3 bucket policy.
