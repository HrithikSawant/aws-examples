# Principle of Least Privilege (PoLP) Learning

The Principle of Least Privilege (PoLP) is a security best practice where users, applications, systems, and services are **granted only the permissions they need to perform their tasks nothing more**. This approach minimizes the potential damage from accidents, malicious actions, or security breaches by limiting access to resources.

**Important Note:** It’s a practice refined through trial, error and and iteration, whether in real environments or homelabs, not a gospel to follow strict steps to achieve PoLP.

## Why PoLP Matters

Implementing PoLP ensures:

- **Security:** Reduces the attack surface by limiting access.
- **Compliance:** Aligns with security practices and regulatory requirements.
- **Accountability:** Ensures users and applications have only the permissions necessary for their roles.

By combining **JEA** and **JIT**, you can achieve an effective PoLP implementation in AWS.

1. [Just-Enough-Access (JEA) Sample](https://github.com/HrithikSawant/aws-examples/blob/main/iam/polp/JEA/README.md)
2. [Just-In-Time (JIT) Sample](https://github.com/HrithikSawant/aws-examples/blob/main/iam/polp/JIT/README.md)

## Key Features of PoLP

**Minimal Permissions:** Assign the least number of actions (e.g., s3:GetObject) necessary for a task.

**Context-Based Access:** Allow access only to specific resources (e.g., one S3 bucket, not all buckets).

**Time-Bound Access:** Use temporary credentials to ensure access expires after a specified period.
Separation of Duties: Split permissions across multiple roles or users to prevent abuse.

### PoLP Use Cases

### 1. Data Access Management in S3

**Scenario:** A data analyst needs to analyze data stored in an S3 bucket for a project.

### Without PoLP

The user is granted the AmazonS3FullAccess policy, allowing them to list, read, write, and delete objects across all S3 buckets in the account.

### With PoLP

The user is granted a policy with permissions limited to the specific S3 bucket and action:

```bash
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::example-bucket/*"
    }
  ]
}

```

### 2. Temporary Developer Access

**Scenario:** A developer needs access to debug an AWS Lambda function for 2 hours.

### Without PoLP

The developer is given administrator-level access (AdministratorAccess policy), granting them permissions to modify any resource in the AWS account.

### With PoLP

The developer assumes a role using AWS STS, which provides temporary credentials valid for only 2 hours.
Permissions are scoped to Lambda-related actions:


```bash
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "lambda:GetFunction",
        "lambda:UpdateFunctionCode",
        "lambda:InvokeFunction"
      ],
      "Resource": "arn:aws:lambda:REGION:ACCOUNT_ID:function:FunctionName"
    }
  ]
}
```

The credentials expire automatically, and the developer cannot perform actions outside Lambda.

### 3. Controlled Access for Third-Party Applications

**Scenario:** A third-party application needs access to send logs to Amazon CloudWatch.

### Without PoLP

The application is provided the CloudWatchFullAccess policy, enabling it to read, write, and delete any CloudWatch logs, metrics, or alarms.

### With PoLP

The application is granted a custom policy allowing only the PutLogEvents action for a specific log group:

```bash
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "logs:PutLogEvents",
      "Resource": "arn:aws:logs:REGION:ACCOUNT_ID:log-group:/example-log-group:*"
    }
  ]
}

```
