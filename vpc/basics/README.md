# AWS VPC Setup and Deletion Scripts

This repository contains two scripts that automate the process of creating and cleaning up AWS Virtual Private Cloud (VPC) resources:

- **VPC Setup Script**: Automates the creation and configuration of a VPC, subnets, internet gateway (IGW), and route tables.
- **VPC Deletion Script**: Removes the resources created by the setup script (or other VPC resources) by detaching the internet gateway, disassociating the route tables, and deleting the subnets and VPC.

## Overview of Actions

### VPC Setup Script

1. **Create a VPC**: A VPC with the CIDR block `your-range/16` is created and tagged with `Name=my-vpc-01`.
2. **Enable DNS Hostnames**: DNS hostnames are enabled for the VPC.
3. **Create an Internet Gateway (IGW)**: An IGW is created and attached to the VPC to enable internet access.
4. **Create Two Subnets**: Two subnets with CIDR blocks `your-range/20` and `your-range/20` are created in the VPC.
5. **Enable Public IP on Subnet Launch**: Auto-assign public IP addresses for instances launched within the subnets.
6. **Create and Associate a Route Table**: A default route table is created and associated with the subnets.
7. **Create a Route to IGW**: A default route to the IGW is added to the route table.
8. **Cleanup Command**: A command to delete all the created resources is displayed for cleanup purposes.

### VPC Deletion Script

1. **Detach Internet Gateway (IGW)**: The specified IGW is detached from the VPC.
2. **Disassociate Route Table**: The route table is disassociated from the provided subnets.
3. **Delete Subnets**: The specified subnets are deleted.
4. **Delete VPC**: The VPC specified by the VPC ID is deleted.
5. **Cleanup**: Resources such as subnets, VPC, and route table associations are cleaned up.

## Usage

### Running the Script

To execute the script, simply run it in your terminal:

```bash
chmod u+x create_vpc
./create_vpc
```

This will initiate the VPC creation process and display the generated resource IDs (VPC, IGW, subnets, etc.) as the script progresses.

```bash
chmod u+x delete_vpc
./delete_vpc <VPC_ID> <IGW_ID> <ASSOC_ID_1> <ASSOC_ID_2> <SUBNET_ID_1> <SUBNET_ID_2> <RT_ID>
```