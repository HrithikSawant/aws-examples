
# S3 Bucket ACL Management with AWS CLI

This guide demonstrates how to manage **ACLs** for an S3 bucket, including operations like creating a new bucket, turning off public access, changing bucket ownership, setting ACLs for cross-account access, and cleaning up.

---

### **Step 1: Create a New Bucket**

Create a new S3 bucket with the specified name:

```bash
aws s3api create-bucket --bucket acl-example-hrithik-fh54265345478 --region us-east-1
```

Example output:

```json
{
    "Location": "/acl-example-hrithik-fh54265345478"
}
```

---

### **Step 2: Turn Off Block Public Access for ACLs**

Use the `s3-acl-manager` script to modify the public access block settings for the bucket. This example disables public ACLs while blocking public policies and restricting public bucket access:

```bash
./s3-acl-manager --bucket acl-example-hrithik-fh54265345478 --block-public-acls false --ignore-public-acls false --block-public-policy true --restrict-public-buckets true
```

Expected output:

```bash
Setting public access block...
Commands executed successfully for bucket: acl-example-hrithik-fh54265345478
Fetching updated public access block configuration...
{
    "PublicAccessBlockConfiguration": {
        "BlockPublicAcls": false,
        "IgnorePublicAcls": false,
        "BlockPublicPolicy": true,
        "RestrictPublicBuckets": true
    }
}
Fetching bucket ACL...
{
    "Owner": {
        "DisplayName": "<Your owner name>",
        "ID": "<owner ID>"
    },
    "Grants": [
        {
            "Grantee": {
                "DisplayName": "<display-name>",
                "ID": "id",
                "Type": "CanonicalUser"
            },
            "Permission": "FULL_CONTROL"
        }
    ]
}
```

---

### **Step 3: Change Bucket Ownership**

To set the ownership of objects within the bucket to the bucket owner, use the following command:

```bash
aws s3api put-bucket-ownership-controls --bucket acl-example-hrithik-fh54265345478 --ownership-controls="Rules=[{ObjectOwnership=BucketOwnerPreferred}]"
```

---

### **Step 4: Get Canonical ID of Another AWS User**

To fetch the **Canonical ID** of another AWS account (useful for cross-account access), you can run the following command:

```bash
aws s3api list-buckets --query Owner.ID --output text
```

Alternatively, you can also find the **Canonical ID** in the AWS Console under **S3 > Permissions > Access Control List**.

---

### **Step 5: Change ACLs to Allow Access for a User in Another AWS Account**

If you want to allow a user in another AWS account to access the bucket, you'll need to create and apply an ACL policy file. For example, the file `policy.json`

Then apply the policy to the bucket:

```bash
aws s3api put-bucket-acl --bucket acl-example-hrithik-fh54265345478 --access-control-policy file:///path/to/policy.json
```

---

### **Step 6: Access the Bucket from Another Account**

To test access from another AWS account, create a file and upload it to the bucket:

```bash
touch bootcamp.txt
aws s3 cp bootcamp.txt s3://acl-example-hrithik-fh54265345478
aws s3 ls s3://acl-example-hrithik-fh54265345478
```

This will list the contents of the bucket, confirming access.

---

### **Step 7: Cleanup**

Finally, you can delete the uploaded file and remove the bucket:

```bash
aws s3 rm s3://acl-example-hrithik-fh54265345478/bootcamp.txt
aws s3 rb s3://acl-example-hrithik-fh54265345478
```

This will remove the file and the bucket.

---

### Summary

This guide covers how to:
1. Create an S3 bucket.
2. Manage public access block settings for the bucket.
3. Change bucket ownership.
4. Configure ACLs for cross-account access.
5. Clean up after testing.

By following these steps, you can effectively manage ACLs and access control on your S3 buckets.

---
