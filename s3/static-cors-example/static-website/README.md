# Static Website Hosting on AWS S3

This guide demonstrates how to host a static website on an AWS S3 bucket.

## Prerequisites

1. AWS CLI installed and configured.
2. A valid AWS account.
3. The necessary permissions to create and configure S3 buckets.

---
## Steps to Host a Static Website

### 1. Create an S3 Bucket

Run the following command to create an S3 bucket. Replace `static-site-example-fh57430` with your desired bucket name:

```bash
aws s3 mb s3://static-site-example-fh57430
```

### 2. Update Bucket Public Access Settings
Update the bucket's public access settings to allow website hosting while maintaining certain security policies:

```bash
aws s3api put-public-access-block \
--bucket static-site-example-fh57430 \
--public-access-block-configuration "BlockPublicAcls=true,IgnorePublicAcls=true,BlockPublicPolicy=false,RestrictPublicBuckets=false"
```

### 3. Apply a Bucket Policy
Create a bucket-policy.json file to define public read permissions for your website content.

Apply the policy to the bucket:

```bash
aws s3api put-bucket-policy --bucket static-site-example-fh57430 --policy file://bucket-policy.json
```

### 4. Enable Static Website Hosting
Create a website.json file.

Apply the configuration to the bucket:

```bash
aws s3api put-bucket-website --bucket static-site-example-fh57430 --website-configuration file://website.json
```

### 5. Upload the Website Content
Upload your index.html file to the bucket:

```bash
aws s3 cp index.html s3://static-site-example-fh57430
```

The static website is now hosted on your S3 bucket. Access it using the following URL formats:

```bash
https://static-site-example-fh57430.s3.ap-south-1.amazonaws.com/index.html
```


### 6. Clean-up
Delete your objects files and bucket.

```bash
aws s3 rm s3://static-site-example-fh57430 --recursive
aws s3 rb s3://static-site-example-fh57430
```

