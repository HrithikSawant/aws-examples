# AWS Just-Enough-Access (JEA) Demo

This demo demonstrates the **Just-Enough-Access (JEA)** principle, which involves granting the minimum required permissions to perform specific tasks in AWS. In this example, we create a role with read-only access to a specific S3 bucket.

## Prerequisites

1. **AWS CLI** installed and configured.
2. An S3 bucket (e.g., `example-bucket`) to which access will be granted.

## Files

1. `create_jea_resources`: Creates the minimal IAM policy and role for JEA.
2. `test_jea_access`: Tests access with the role's temporary credentials.
3. `cleanup_jea_resources`: Deletes the resources created during the demo.

## Usage

### Step 1: Create Resources

Run the `create_jea_resources` script to create the JEA policy and role.

```bash
./create_jea_resources
```

This script will:

- Create an IAM policy with read-only access to the specified S3 bucket.
- Create an IAM role that allows assuming the policy.

## Step 2: Test Access

Use the **test_jea_access** script to assume the role and test the minimal access.

```bash
./test_jea_access
```

This script will:

- Assume the role and retrieve temporary credentials.
- Use the temporary credentials to attempt access to the S3 bucket.

## Step 3: Cleanup Resources

After the demo, run the **cleanup_jea_resources** script to delete the created resources.

```bash
./cleanup_jea_resources
```

Note:

Replace placeholders like **YOUR_ACCOUNT_ID** and example-bucket in the scripts with actual values.
The JEA policy grants read-only access (s3:GetObject) to objects in the specified bucket.
Temporary credentials are automatically invalidated after their expiration period.