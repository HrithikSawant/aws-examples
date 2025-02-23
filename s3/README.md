# Amazon S3 Examples

This directory contains examples and scripts to help you interact with Amazon S3 using AWS CLI.

<details>
  <summary> 0. S3 Bucket Operations and Scripts</summary>

- **Create an S3 Bucket**  
  - Script to create a new S3 bucket with specific configurations (including object lock enabled).
  - Example usage:

       ```bash
       ./create-bucket
       ```

- **Delete an S3 Bucket**  
  - Script to delete an S3 bucket.
  - Example usage:

       ```bash
       ./delete-bucket
       ```

- **Get the Most Recently Created S3 Bucket**  
  - Script to get the most recently created S3 bucket.
  - Example usage:

       ```bash
       ./get-newest-bucket
       ```

- **Upload Files to S3**  
  - Script to upload random files to a specific S3 bucket.
  - Example usage:

       ```bash
       ./sync
       ```

- **Delete Objects from S3 Bucket**  
  - Script to delete objects in a specific S3 bucket.
  - Example usage:

       ```bash
       ./delete-objects
       ```

- **List Objects in S3 Bucket**  
  - Script to list all objects in a specific S3 bucket.
  - Example usage:

       ```bash
       ./list-objects
       ```

</details>

1. [S3 Bucket and Prefix Demo](https://github.com/HrithikSawant/aws-examples/blob/main/s3/prefixes/README.md)

2. [S3 Bucket Policy Demo](https://github.com/HrithikSawant/aws-examples/blob/main/s3/bucket-policy/README.md)

3. [S3 etags](https://github.com/HrithikSawant/aws-examples/blob/main/s3/etags/README.md)

4. [S3 metadata](https://github.com/HrithikSawant/aws-examples/blob/main/s3/metadata/README.md)

5. [S3 checksums](https://github.com/HrithikSawant/aws-examples/blob/main/s3/checksums/README.md)

6. [S3 Server-Side Encryption (SSE) and Server-Side Encryption with Customer-Provided Keys (SSE-C)](https://github.com/HrithikSawant/aws-examples/blob/main/s3/encryption/README.md)

7. [S3 Encrypting Data Locally](https://github.com/HrithikSawant/aws-examples/blob/main/s3/encryption-client/README.md)

8. [S3 Bucket ACL Demo](https://github.com/HrithikSawant/aws-examples/blob/main/s3/acl/README.md)

9. [S3 Bucket via CNF](https://github.com/HrithikSawant/aws-examples/blob/main/s3/cnf/README.md)

## Prerequisites

To use these examples, ensure the following:

- AWS CLI is installed and configured with valid credentials.
- S3 permissions are properly set up in your AWS account.

## Additional Notes

- For more details follow [AWS S3 Documentation](https://docs.aws.amazon.com/s3/index.html).
