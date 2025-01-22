# Egress-Only Internet Gateway (EO-IGW) Demo

## Overview
This guide walks you through the process of setting up an **Egress-Only Internet Gateway (EO-IGW)** in AWS. EO-IGWs are used to allow outbound IPv6 traffic from resources in your VPC, while blocking inbound IPv6 traffic. They are typically used in IPv6-only VPCs to allow internet access without allowing inbound connections.

### Key Concepts:
- **VPC (Virtual Private Cloud)**: A logically isolated network in the AWS cloud.
- **Subnet**: A range of IP addresses within your VPC.
- **Egress-Only Internet Gateway (EO-IGW)**: A gateway used to send IPv6 traffic from the VPC to the internet while blocking incoming IPv6 traffic.

## Prerequisites
Before proceeding, ensure that you have:
- An **AWS account**.
- AWS CLI installed and configured on your local machine.
- Permissions to create VPCs, subnets, and internet gateways.

## Steps to Set Up EO-IGW

2. **Run the Script**
   Update and Execute the script in your local terminal to automatically create the necessary resources in AWS. The script will:
   - Create a VPC with IPv6 CIDR blocks.
   - Create a subnet with both IPv4 and IPv6 addresses.
   - Enable DNS64 for the subnet.
   - Set up and attach an **Egress-Only Internet Gateway** to the VPC.

3. **Check the Status**
   After running the script, verify that the Egress-Only Internet Gateway is properly attached and configured by logging into your AWS Console and checking the resources in the **VPC** and **Internet Gateway** sections.

## Additional Information
- For further details on Egress-Only Internet Gateways, refer to the [AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Egress-Only_Internet_Gateway.html).
  
## Cleanup
Once you've completed your testing or demo, you can delete the resources created by the script to avoid unnecessary charges:
- Delete the **Egress-Only Internet Gateway**.
- Delete the **VPC** and its associated resources.
