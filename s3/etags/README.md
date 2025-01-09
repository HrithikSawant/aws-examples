
# ETag Demo with Terraform

This guide demonstrates how to use the `etag` attribute in an S3 object resource to ensure file integrity during uploads using Terraform.

---

## Prerequisites

- Install [Terraform](https://www.terraform.io/downloads.html).
- Ensure you have AWS credentials configured locally.
- A file named `myfile.txt` must exist in the same directory as `main.tf`.

---

## Steps

### 1. Initialize Terraform
Run the following command to initialize Terraform and download the required provider:

```bash
terraform init
```

### 2. Apply the Configuration
Use the following command to create the S3 bucket and upload the file to S3:

```bash
terraform apply
```

- Review the changes and type `yes` to confirm.

### 3. Verify ETag Integrity
- After the upload, Terraform will ensure that the uploaded object in S3 matches the MD5 checksum of the local file (`myfile.txt`) using the `etag` attribute.

### 4. Cleanup
To delete the resources created by this configuration, run:

```bash
terraform destroy
```

- Type `yes` to confirm.

---

## Notes

- The Terraform configuration is defined in the `main.tf` file, which creates an S3 bucket and uploads a file (`myfile.txt`) to it.
- The `etag` attribute ensures file integrity by comparing the MD5 hash of the local file with the ETag of the uploaded file in S3.

This demo highlights the importance of using ETags to verify file consistency during uploads.
