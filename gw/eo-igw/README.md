# Egress-Only Internet Gateway (EO-IGW) Demo

## Overview
Simple Demo for an **Egress-Only Internet Gateway (EO-IGW)** in AWS. EO-IGWs are used to allow outbound IPv6 traffic from resources in your VPC, while blocking inbound IPv6 traffic.

## Requirements

Before running this script, ensure that you have:

- AWS CLI installed and configured with necessary credentials.
- Appropriate IAM permissions to create and manage VPCs, subnets, Internet Gateways, EC2 instances, and other related resources.
- A key pair (e.g., `MytestKey`) for SSH access to EC2 instances.

## Steps to Run the Script

1. **Set Up the VPC**: The script creates a VPC with both IPv4 and IPv6 CIDR blocks, which allows dual-stack (IPv4 and IPv6) communication.

2. **Subnet Creation**: A subnet with both IPv4 and IPv6 CIDR blocks is created in the VPC. DNS64 is enabled for IPv6-only clients to resolve IPv4 addresses.

3. **Egress-Only Internet Gateway**: An Egress-Only Internet Gateway is created for outbound IPv6 traffic, enabling IPv6 clients to access the internet for outbound connections.

4. **Internet Gateway**: A standard Internet Gateway is created and attached to the VPC for outbound IPv4 traffic.

5. **NAT Gateway**: A NAT Gateway is created in the subnet with an Elastic IP address, enabling instances in private subnets to access the internet for IPv4 traffic.

6. **Route Table**: The script creates a route table with routes for both IPv4 and IPv6 traffic:
    - Routes all IPv6 traffic (`::/0`) to the Egress-Only Internet Gateway.
    - Routes IPv6 NAT64 traffic (`64:ff9b::/96`) to the NAT Gateway. 
        - The **64:ff9b::/96** address block is used by NAT64 to allow IPv6-only clients to communicate with IPv4 servers. It maps IPv6 addresses to IPv4 addresses by encoding the IPv4 address in the lower 32 bits of the IPv6 address.
    - Routes all IPv4 traffic (`0.0.0.0/0`) to the Internet Gateway.

7. **Security Group**: A security group is created and configured to allow:
    - SSH access on port 22 from any IP.
    - ICMP (ping) traffic from any IP.

8. **Launch EC2 Instance**: Finally, the script launches an EC2 instance with the following features:
    - An attached public IPv4 address.
    - An assigned IPv6 address.
    - An associated security group for network access control.
    - Tags for identification.
    - Custom metadata and DNS name options.


## Steps to Set Up EO-IGW

1. **Run the Script**
   Update and Execute the script.
   ```bash
   chmod u+x eo-igw
   ./eo-igw
   ````

2. **Check the Status**
   After running the script, verify that the Egress-Only Internet Gateway is properly attached and configured by logging into your AWS Console and checking the resources in the **VPC** and **Internet Gateway** sections.

## Testing

### Connect to the EC2 instance using SSH (replace <Your EC2 IP> with the actual public IP of your instance)
```bash
ssh -i "MyPemtestfile.pem" ec2-user@<Your EC2 IP>
```

```bash
# Copy the script to instance using scp
./ipv6_test
```

### Test from external source
```bash
./verify_ipv6_reachability
```

## Cleanup
Once you've completed your testing or demo, you can delete the resources created by the script to avoid unnecessary charges:
- Delete the **Egress-Only Internet Gateway**.
- Delete the **VPC** and its associated resources.

```bash
./clean-up
```

## Additional Information
- For further details on Egress-Only Internet Gateways, refer to the [AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/egress-only-internet-gateway.html).
  




