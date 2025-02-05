
# AWS IAM User and Role Setup with STS and Access Control

This guide provides step-by-step instructions for creating a new IAM user with no permissions, generating access keys, creating an IAM role, assuming that role using the user credentials, and testing access. It also includes cleanup steps. This guide uses placeholder values that should be replaced with your own specific details.

## Step 1: Create a New User with No Permissions

Create an IAM user with no permissions. Replace `<username>` with the desired IAM username.

```bash
aws iam create-user --user-name <username>
```

### Generate Access Keys for the User

Generate access keys for the new IAM user. Replace `<username>` with the IAM username created earlier.

```bash
aws iam create-access-key --user-name <username> --output table
```

Make sure to save the `AccessKeyId` and `SecretAccessKey` for later use.

---

## Step 2: Verify Current User Permissions

Before proceeding with the role assumption, verify that the newly created user has no permissions.

### Test Access with `sts get-caller-identity`

Ensure that you're working under the newly created user's profile by running:

```bash
aws sts get-caller-identity --profile <profile-name>
```

### Ensure No Access to S3

Try listing the S3 buckets. This should return an error indicating that the user does not have access to S3.

```bash
aws s3 ls --profile <profile-name>
```

---

## Step 3: Create an IAM Role for Resource Access

Create a role that will allow access to resources. Replace `<role-name>` and `<path-to-trust-policy>` with your desired role name and path to your trust policy JSON file.

```bash
aws iam create-role \
  --role-name <role-name> \
  --assume-role-policy-document file://<path-to-trust-policy>
```

Attach necessary policies to the role. Replace `<role-name>` with your role name, and `<policy-arn>` with the policy ARN (e.g., `AmazonS3ReadOnlyAccess`).

```bash
aws iam attach-role-policy \
  --role-name <role-name> \
  --policy-arn <policy-arn>
```

---

## Step 4: Attach a Policy to Allow the User to Assume the Role

Assign a policy to the user allowing them to assume the role. Replace `<username>` and `<path-to-policy>` with the IAM username and the path to your policy JSON file.

```bash
aws iam put-user-policy \
  --user-name <username> \
  --policy-name <policy-name> \
  --policy-document file://<path-to-policy>
```

The `policy.json` should contain the necessary permissions to assume the role

---

## Step 5: Assume the Role Using the User's Credentials

Assume the role using the `sts-machine-user` profile. Replace `<role-arn>` with the role ARN you created, and `<session-name>` with a name for the session.

```bash
aws sts assume-role \
  --role-arn arn:aws:iam::<aws-account-id>:role/<role-name> \
  --role-session-name <session-name> \
  --profile <profile-name>
```

This command will output temporary credentials (AccessKeyId, SecretAccessKey, and SessionToken).

---

## Step 6: Configure AWS CLI to Use the Assumed Role

Configure a new AWS CLI profile using the temporary credentials returned by the `assume-role` command. Replace the placeholders `<access-key-id>`, `<secret-access-key>`, and `<session-token>` with the actual temporary credentials.

```bash
aws configure set aws_access_key_id "<access-key-id>" --profile <new-profile-name>
aws configure set aws_secret_access_key "<secret-access-key>" --profile <new-profile-name>
aws configure set aws_session_token "<session-token>" --profile <new-profile-name>
```

---

## Step 7: Test the Assumed Role's Permissions

Now that the `assumed-role` profile is configured, test its permissions. For example, try listing the S3 buckets:

```bash
aws s3 ls --profile <new-profile-name>
```

If the role has the correct permissions, the S3 buckets should be listed.

---

## Step 8: Clean Up Resources

After testing, clean up by deleting the IAM policies, roles, users, and other resources you created.

### Delete User Policy

```bash
aws iam delete-user-policy --user-name <username> --policy-name <policy-name>
```

### Delete Role Policy

```bash
aws iam delete-role-policy --role-name <role-name> --policy-name <policy-name>
```

### Delete the Role

```bash
aws iam delete-role --role-name <role-name>
```

### Delete the S3 Bucket (If Applicable)

If you created an S3 bucket during the test, delete it using:

```bash
aws s3 rb s3://<bucket-name> --force
```

### Delete Access Keys

Delete any access keys associated with the user:

```bash
aws iam list-access-keys --user-name <username>
aws iam delete-access-key --user-name <username> --access-key-id <access-key-id>
```

### Delete the User

Finally, delete the IAM user:

```bash
aws iam delete-user --user-name <username>
```
---

## Conclusion

By following this guide, you have successfully created an IAM user, created a role with the appropriate permissions, used STS to assume the role, tested the assumed role’s permissions, and cleaned up the resources. Make sure to replace the placeholders with your actual values before running the commands.
