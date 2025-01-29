# EC2 Metadata Retrieval Script

## Overview

This script fetches metadata of an EC2 instance using the AWS EC2 Instance Metadata Service (IMDS). It provides key information such as the instance ID, public IPv4 address, availability zone, instance type, and IAM role details.

## Features:

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
