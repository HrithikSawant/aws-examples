# AWS Just-In-Time (JIT) Access Demo

This demo showcases how to implement **Just-In-Time (JIT)** access in AWS using IAM and Security Token Service (STS). JIT access ensures that users are granted temporary permissions only when needed.

## Prerequisites

1. **AWS CLI** installed and configured.
2. An existing IAM user (e.g., `JITUser`).
3. An S3 bucket to which access is being granted (e.g., `example-bucket`).

## Files

1. `create_resources.sh`: Sets up the required IAM role and attaches minimal permissions.
2. `assume_role.sh`: Generates temporary credentials for the JIT user by assuming the IAM role.
3. `cleanup.sh`: Cleans up the IAM resources created during the demo.

## Usage

### Step 1: Create Resources

Run the `create_resources.sh` script to set up the IAM role and minimal access policy.

```bash
./create_resources.sh
```

This script will:

- Create an IAM policy with read-only access to the specified S3 bucket.
- Create an IAM role that allows the designated user (JITUser) to assume it.
- Attach the policy to the role.

### Step 2: Assume Role
Run the assume_role.sh script to assume the IAM role and obtain temporary credentials.

```bash 
./assume_role.sh
```

After running this script:

- Temporary credentials (Access Key, Secret Key, and Session Token) will be exported to your environment.
- Use these credentials to perform AWS CLI actions, such as listing objects in the S3 bucket:

```bash
aws s3 ls s3://example-bucket/
```

Temporary credentials will expire after the default session duration (1 hour).

### Step 3: Clean Up Resources

To avoid unnecessary costs, run the cleanup.sh script to delete the IAM role, policy, and other resources created during the demo.

```bash
./cleanup.sh
```

## Notes

- Replace placeholders like YOUR_ACCOUNT_ID, example-bucket, and JITUser in the scripts with your actual AWS resources and user details.
- Temporary credentials are automatically invalidated after their expiration, ensuring limited-time access.


### Example Use Case

**Scenario:** A developer needs access to an S3 bucket for a short task.
**Solution:** They run assume_role.sh to gain temporary credentials for the task.
**Outcome:** The developer completes the task, and credentials automatically expire after the session duration, ensuring no lingering access.
