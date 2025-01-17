# README

## Overview

Demo of **Network Access Control List (NACL)** ingress rules in an AWS environment. It shows how to launch a simple EC2 instance, configure NACL ingress rules, and manage VPC resources using AWS CloudFormation and custom scripts.

The steps include creating and deleting a VPC, deploying the EC2 instance using CloudFormation, and configuring NACL ingress rules to control traffic.

### Files Included

1. **CloudFormation Template (`template.yaml`)**:
   - Deploys an EC2 instance with an Apache HTTP server configured via UserData.
   - Configures IAM roles, Security Groups, and Network Interfaces for EC2 instance deployment.

2. **`fetch_ami_id`**:
   - Fetches the latest Amazon Linux 2 AMI ID for use in the CloudFormation template.

3. **`nacl-rule`**:
   - A script to create ingress/egress rules in a specified NACL to control traffic to your instances.

4. **`create_stack`**:
   - A script to create a CloudFormation stack using the provided `template.yaml`.

5. **`create_vpc`** and **`delete_vpc`**:
   - Scripts to create and delete a VPC for testing and deployment.

---

## Steps for Deployment and Configuration

### 1. **Set Up the Environment**

Ensure that you have the AWS CLI installed and configured with the correct IAM permissions to create and delete resources.

### 2. **Fetch Latest AMI ID**

Run the `fetch_lastest_ami` script to get the latest Amazon Linux 2 AMI ID:

```bash
./fetch_lastest_ami
```

This will output the `AMI_ID`, which will be used in the CloudFormation template.

### 3. **Create VPC (Optional)**

If you need to create a new VPC, use the provided `create_vpc` script:

```bash
cd aws-examples/vpc/basics/bin
./create_vpc
```

This will create a VPC for the CloudFormation stack. Make sure to note the VPC and Subnet IDs for later.

Fetch **ACL_ID**

```bash
aws ec2 describe-network-acls --filters "Name=vpc-id,Values=<YOUR_VPC_ID>" --query "NetworkAcls[0].NetworkAclId" --output text
```

### 4. **Update CloudFormation Template**

In the `template.yaml` file, update the following fields:

- **VPC_ID**: Replace `<YOUR-VPC>` with your VPC ID.
- **AMI_ID**: Replace the default AMI ID with the ID you fetched from the `fetch_ami_id` script.
- **Subnet_ID**: Replace `<YOUR-SUBNET>` with the correct subnet ID.


### 5. **Launch CloudFormation Stack**

To deploy the CloudFormation stack, use the `create_stack` script. This will create an EC2 instance and related resources defined in the `template.yaml` file:

```bash
./create_stack template.yaml MyIngressStack-fh904353
```

This command will initiate the creation of the CloudFormation stack and all related resources.


### 6. **Test the Apache Web Server**

Once the EC2 instance is up and running, you can test the Apache HTTP server by accessing the instance’s public IP through a web browser.

- Find the public IP of the EC2 instance (you can check the EC2 console or run `aws ec2 describe-instances`).
- Open a web browser and go to: `http://<EC2_PUBLIC_IP>`

You should see a page displaying the message:

<html><body><h1>Hello from Apache on Amazon Linux 2!</h1></body></html>

### 7. **Create NACL Ingress Rule**

To create an ingress rule in your NACL, run the `create_nacl_entry` script. For example, to deny traffic from a specific IP (`your-public-ip/32`) on all ports: Mention the ip from where the http request is initiated

```bash
./nacl_rule --ingress  <YOUR_NACL_ID> 90 -1 From=0,To=65535 <YOUR_PUBLIC_IP>/32 deny
```

This command creates a rule in the NACL with the following parameters:
- **ACL ID**: Replace `acl-xxxxxxxx` with your actual NACL ID.
- **Rule number**: 90 (you can adjust the rule number).
- **Protocol**: -1 for all protocols.
- **Port range**: From 0 to 65535.
- **CIDR block**: `<YOUR-Public-IP>/32` (the IP you want to block).
- **Action**: Deny.

### 8. **Test NACL Ingress Rule**

To test the NACL ingress rule you created in Step 6, you can try to access the Apache web server from  the IP you specified in the rule. The access should be denied based on the NACL rule you applied.

To verify the ingress rule:

- Attempt to access the Apache server from a browser or `curl` from the IP `<YOUR-IP>` (or any other blocked IP).
- The connection should be denied, confirming that the NACL rule is working as expected.

curl: (28) Failed to connect to <YOUR-PUBLIC-IP> port 80 after 133521 ms: Connection timed out

### 9. **Delete CloudFormation Stack**

To delete the CloudFormation stack when done, run:

```bash
aws cloudformation delete-stack --stack-name MyIngressStack-fh904353
```

### 10. **Delete VPC (Optional)**

If you created a new VPC using the `create_vpc` script, you can delete it with:

```bash
cd aws-examples/vpc/basics/bin
./delete_vpc
```

This will clean up the VPC and any associated resources.

---
