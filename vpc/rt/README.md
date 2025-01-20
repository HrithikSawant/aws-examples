
# Route Table Demo README

## Overview

This demo provides scripts create sample route table 

## Pre-requisites

Ensure that you have the following before running the scripts:

- AWS CLI configured with proper permissions.
- A valid AWS account.
- Appropriate IAM roles/permissions to create/delete VPC, subnets, route tables, and manage associations.
- A configured VPC with necessary subnet, IGW, and other resources.

## Example Workflow

Here’s an example of a typical workflow using the scripts:

1. **Create a VPC**:  This script is reused from vpc/basics
   ```
   ./create_vpc
   VPC_ID: vpc-xxxxxxxxxxxxxxxxx
   IGW_ID: igw-xxxxxxxxxxxxxxxxx
   SUBNET_ID_1: subnet-xxxxxxxxxxxxxxxxx
   SUBNET_ID_2: subnet-xxxxxxxxxxxxxxxxx
   RT_ID: rtb-xxxxxxxxxxxxxxxxx
   ASSOC_ID_1: rtbassoc-xxxxxxxxxxxxxxxxx
   ASSOC_ID_2: rtbassoc-xxxxxxxxxxxxxxxxx
   {
       "Return": true
   }
   ./delete_vpc vpc-xxxxxxxxxxxxxxxxx igw-xxxxxxxxxxxxxxxxx rtbassoc-xxxxxxxxxxxxxxxxx rtbassoc-xxxxxxxxxxxxxxxxx subnet-xxxxxxxxxxxxxxxxx subnet-xxxxxxxxxxxxxxxxx rtb-xxxxxxxxxxxxxxxxx

   ```
   - Creates the VPC, IGW, and subnets, and associates the route tables.

2. **Create a new subnet without association**:
   ```
   ./create_subnets_without_association vpc-xxxxxxxxxxxx test-subnet-rt 10.1.0.0/20
   No availability zone provided. Defaulting to: xx-region
   Creating subnet test-subnet-rt in VPC vpc-xxxxxxxxxxxx with CIDR 10.1.0.0/20 in Availability Zone xx-region...
   Subnet test-subnet-rt created with ID: subnet-xxxxxxxxxxxx
   Subnet creation complete. The subnet test-subnet-rt is not associated with any route table.

   ```
   - Creates a subnet in the VPC without associating it to any route table.

3. **Create a custom route table and associate it with a subnet**:
   ```
   ./create_route_table
   Enter route table name: test-custom-rt
   Enter VPC ID (e.g., vpc-xxxxxxxx): vpc-xxxxxxxx
   Enter at least one Subnet ID (e.g., subnet-xxxxxxxx): subnet-xxxxxxxx
   Enter destination CIDR block for route (e.g., 0.0.0.0/0) [optional]: 
   Enter gateway ID for route (e.g., igw-xxxxxxxx) [optional]: igw-xxxxxxxx
   Creating route table 'test-custom-rt' in VPC vpc-xxxxxxxx...
   Route table created with ID: rtb-xxxxxxxx
   Associating route table with Subnet subnet-xxxxxxxx...
   {
       "AssociationId": "rtbassoc-xxxxxxxx",
       "AssociationState": {
           "State": "associated"
       }
   }
   Route table successfully associated with subnet subnet-xxxxxxxx.
   No route added.
   Script execution complete.

   ```
   - Creates a new route table and associates it with the subnet.

4. **Delete resources**:
   - Delete a subnet:  
     ```
     ./delete_subnet
     Enter VPC ID (e.g., vpc-xxxxxxxx): vpc-xxxxxxxx
     Enter Subnet ID (e.g., subnet-xxxxxxxx): subnet-xxxxxxxx
     Checking if subnet subnet-xxxxxxxx exists in VPC vpc-xxxxxxxx...
     Deleting subnet subnet-xxxxxxxx from VPC vpc-xxxxxxxx...
     Subnet subnet-xxxxxxxx has been successfully deleted.
     Removing tags from subnet subnet-xxxxxxxx...
     Tags removed from subnet subnet-xxxxxxxx (if any existed).
     Subnet deletion complete.

     ```
   - Delete a route table:  
     ```
     ./delete_route_table 
      Enter route table ID (e.g., rtb-xxxxxxxx): rtb-xxxxxxxx
      Enter VPC ID (e.g., vpc-xxxxxxxx): vpc-xxxxxxxx
      Enter at least one Subnet ID (e.g., subnet-xxxxxxxx): subnet-xxxxxxxx
      Checking if route table rtb-xxxxxxxx exists in VPC vpc-xxxxxxxxxxxxxx...
      Disassociating route table from subnet subnet-xxxxxxxxxxxx...
      No association found for subnet subnet-xxxxxxxxxxxx.
      Deleting route table rtb-xxxxxxxxxxxx...
      Route table rtb-xxxxxxxxxxxx has been successfully deleted.
      Removing tags from route table rtb-xxxxxxxxxxxx...
      Tags removed from route table rtb-xxxxxxxxxxxx (if any existed).
      Route table cleanup complete.
     ```

2. **delete_vpc**  
   This script deletes the VPC along with all its associated resources, such as subnets, route tables, and internet gateways.

   - Delete the VPC: This script is reused from vpc/basics
   ```
   ./delete_vpc <VPC_ID> <IGW_ID> <ASSOC_ID_1> <ASSOC_ID_2> <SUBNET_ID_1> <SUBNET_ID_2> <RT_ID>
   ```

   **Output Example**: 
   ```bash
   Detaching Internet Gateway igw-xxxxxxxxxxxx from VPC vpc-xxxxxxxxxxxx...
   Disassociating route table subnet-xxxxxxxxxxxx from subnet subnet-xxxxxxxxxxxx with Association ID: rtbassoc-xxxxxxxxxxxx
   Deleting subnet subnet-xxxxxxxxxxxx...
   Deleting subnet subnet-xxxxxxxxxxxx...
   Deleting VPC vpc-xxxxxxxxxxxx...
   ```
