
# CLI Installation Scripts

Scripts to install commonly used CLI tools: AWS CLI, PowerShell CLI, and Terraform CLI. Each script is designed to automate the installation process on a Linux-based system.

---

## Table of Contents
1. [AWS CLI Installation](#aws-cli-installation)
2. [PowerShell CLI Installation](#powershell-cli-installation)
3. [Terraform CLI Installation](#terraform-cli-installation)
4. [How to Run](#how-to-run)
5. [Prerequisites](#prerequisites)

---

## AWS CLI Installation

### Description:
Installs the AWS CLI (version 2) on your system.

### Commands:
```bash
./aws_cli_install
```

---

## PowerShell CLI Installation

### Description:
Installs PowerShell CLI and prepares it for AWS Tools integration (optional).

### Commands:
```bash
./powershell_cli_install
```

---

## Terraform CLI Installation

### Description:
Installs the latest version of Terraform CLI on your system.

### Commands:
```bash
./terraform_cli_install
```

---

## How to Run

### Step 2: Make the Scripts Executable
```bash
chmod u+x aws_cli_install powershell_cli_install terraform_cli_install
```

### Step 3: Run the Scripts
To install AWS CLI:
```bash
./aws_cli_install
```

To install PowerShell CLI:
```bash
./powershell_cli_install
```

To install Terraform CLI:
```bash
./terraform_cli_install
```

## Notes

- For PowerShell, after installation, you can start it using:
  ```bash
  pwsh
  ```
- For AWS CLI, configure it after installation:
  ```bash
  aws configure
  ```
- For Terraform, verify the installation with:
  ```bash
  terraform --version
  ```