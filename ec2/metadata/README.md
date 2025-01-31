# EC2 Metadata Retrieval Script

## Overview

This script fetches metadata of an EC2 instance using the AWS EC2 Instance Metadata Service (IMDS). It provides key information such as the instance ID, public IPv4 address, availability zone, instance type, and IAM role details.

## Features

- Retrieves metadata like instance ID, availability zone, instance type, etc.
- Works only on EC2 instances (checks if it's running on EC2).
- Supports output in JSON format.
- Displays all metadata or selected fields.
- Uses IMDSv2

## Usage

```bash
# retrieve all metadata
./metadata --all

# retrieve metadata in json
./metadata -j

# retrive instance-id
./metadata -f instance-id

# help
./metadata -h
```

## Features

- Retrieve metadata EC2 Instance Family, Processor Info, Instance Size, Instance Profile, Instance Lifecycle State

## usage

```bash
# retrieve metadata commands
./ec2-metadata -f FAMILY
./ec2-metadata -p PROCESSOR
./ec2-metadata -s SIZE
./ec2-metadata -r PROFILE
```

## Features

- **EC2 Instance Console Screenshot**: Fetching and displaying EC2 instance metadata
- **EC2 Hostnames**: Fetching and displaying EC2 instance metadata
- **EC2 Default Username**: Fetching and displaying EC2 instance metadata
- **EC2 Burstable Instances**: Fetching and displaying EC2 instance metadata
- **EC2 Source & Destination Checks**: Fetching and displaying EC2 instance metadata
- **EC2 System Log**: Fetching and displaying EC2 instance metadata
- **EC2 Placement Groups**: Fetching and displaying EC2 instance metadata

## usage
  
```bash
# retrieve metadata 
./ec2-metadata-fetcher -s  Get EC2 instance console screenshot
./ec2-metadata-fetcher -h  Get EC2 private and public hostnames
./ec2-metadata-fetcher -u  Get default username based on AMI
./ec2-metadata-fetcher -b  Check if instance is burstable (T2, T3)
./ec2-metadata-fetcher -c  Check if source/destination check is enabled
./ec2-metadata-fetcher -l  Get EC2 system log
./ec2-metadata-fetcher -p  Get EC2 placement group information
./ec2-metadata-fetcher -a  Get all available metadata
```

## Features

- **EC2 AMI IMAGE**: creating ami image of an instance

## usage
  
```bash
./ec2-management -a
```