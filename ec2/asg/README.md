# AWS Auto Scaling Group (ASG) Creation Script

## Overview
This script automates the creation of an **AWS Auto Scaling Group (ASG)** using the AWS CLI. It sets up a launch template and an ASG to dynamically scale EC2 instances based on demand.

## Prerequisites
Before running this script, ensure you have:
- AWS CLI installed and configured with proper credentials.
- A valid **AMI ID** (Amazon Machine Image) for your EC2 instances.
- At least two subnets in different Availability Zones.

## Usage
Run the script with optional parameters to customize the ASG settings:

```bash
./asg --asg-name "MyASG" --instance-type "t3.micro" --min-size 1 --max-size 3 --desired 2
```

### Options:
- `-a, --asg-name`      : Set the Auto Scaling Group name (default: `MyAutoScalingGroup`)
- `-t, --template-name` : Set the Launch Template name (default: `MyLaunchTemplate`)
- `-i, --instance-type` : Set the EC2 instance type (default: `t2.micro`)
- `-m, --min-size`      : Set the minimum number of instances (default: `2`)
- `-x, --max-size`      : Set the maximum number of instances (default: `5`)
- `-d, --desired`       : Set the desired number of instances (default: `2`)
- `-s, --subnets`       : Set the subnets for the ASG (default: predefined values)
- `-h, --help`          : Show this help message

## How It Works
1. **Creates a Launch Template**
   - Defines the AMI, instance type, and tags.
   - This ensures all instances created by the ASG follow the same configuration.
2. **Creates an Auto Scaling Group (ASG)**
   - Automatically scales EC2 instances up or down based on demand.
   - Ensures high availability and cost-efficiency.

## Benefits of ASG
- **Automatic Scaling**: Adjusts EC2 instances based on traffic.
- **High Availability**: Distributes instances across multiple AZs.
- **Cost Optimization**: Reduces costs by running only required instances.

## Conclusion
This script helps deploy an ASG quickly, ensuring scalability and efficiency. You can modify it further to integrate with **ALB (Application Load Balancer)** for better traffic distribution.

For more details, check the [AWS ASG Documentation](https://docs.aws.amazon.com/autoscaling/ec2/userguide/what-is-amazon-ec2-auto-scaling.html).

