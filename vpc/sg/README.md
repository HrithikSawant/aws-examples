# EC2 Instance Deployment with Apache Web Server (Security Group Demo)

This CloudFormation template is designed as a **demo for Security Group configuration**. It deploys an Amazon EC2 instance in a specified VPC and subnet with unrestricted network traffic for demonstration purposes only. The template also includes an IAM role for Systems Manager (SSM) access and a User Data script to set up an Apache web server with a sample webpage.

> **IMPORTANT**: This template is for **demonstration and testing purposes only**. It is not recommended for production environments due to its open Security Group settings.

## Template Overview

### Purpose

The main objective of this template is to demonstrate the configuration of a Security Group (`SG`) with unrestricted traffic (`0.0.0.0/0`).

### Features

- **EC2 Instance**: A single Amazon EC2 instance with a specified AMI and instance type.
- **IAM Role**: Grants the instance access to AWS Systems Manager for management tasks.
- **Security Group**: Configured with **unrestricted inbound and outbound traffic** for demonstration purposes.
- **User Data Script**: Sets up Apache web server and deploys a sample HTML webpage.

## Usage Instructions

### Prerequisites

1. Ensure you have:
   - AWS CLI or access to the AWS Management Console.
   - Necessary permissions to create resources in your AWS account.
2. Replace placeholder values in the parameters (`VpcId`, `ImageId`, and `SubnetId`) with actual IDs from your AWS environment.

### Deployment Steps

1. Save the CloudFormation template as `template.yaml`.
2. Deploy the stack using one of the following methods:

#### AWS Management Console

1. Open the **CloudFormation** service in the AWS Management Console.
2. Create a new stack and upload the `template.yaml` file.
3. Enter the required parameters (`VpcId`, `ImageId`, `SubnetId`) and follow the prompts.

#### AWS CLI

Run the following command:

```bash
aws cloudformation create-stack --stack-name security-group-demo   --template-body file://template.yaml
```

### Testing the Deployment

1. Retrieve the public IP address of the deployed EC2 instance:

   ```bash
   aws ec2 describe-instances --query "Reservations[*].Instances[*].PublicIpAddress" --output text
   ```

2. Open the public IP address in a web browser. You should see the sample webpage:

   ```
   <html><body><h1>Hello from Apache on Amazon Linux 2!</h1></body></html>
   ```

---

## Cleanup

To avoid incurring charges, delete the CloudFormation stack when you are finished:

```bash
aws cloudformation delete-stack --stack-name security-group-demo
```

---

## Security Considerations

1. **Unrestricted Traffic**: The Security Group (`SG`) in this template is configured with unrestricted traffic (`0.0.0.0/0`):
   - This is suitable for testing but poses a significant security risk.
   - For production environments, restrict inbound traffic to specific IP ranges or ports, such as SSH (`22`) or HTTP (`80`).
2. **Access Control**: Ensure that the IAM role (`SSMRole`) grants only the necessary permissions.
3. **Credentials**: Avoid embedding sensitive data in the template. Use AWS Secrets Manager or Parameter Store for secure data management.

---

## Notes

For more details on Security Group configurations, consult the [AWS Security Group Documentation](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html).
