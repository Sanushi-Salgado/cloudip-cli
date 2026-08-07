# cloudip-cli

A lightweight command-line tool to identify whether an IP address belongs to a known cloud provider using offline CIDR database lookups.

Built on top of:

- https://github.com/rezmoss/go-cloudip

`cloudip-cli` converts the `go-cloudip` library into a simple terminal command for quick IP ownership checks during passive reconnaissance and asset enrichment.

---

## Features

- Offline cloud IP detection
- Fast CIDR lookup using Patricia trie
- No active scanning
- No connection to the target IP
- Supports IPv4 and IPv6 ranges
- Displays:
  - Cloud provider
  - Region
  - Service
  - Matched CIDR

Supported providers depend on the upstream database:

```
AWS
GCP
Cloudflare
Azure
DigitalOcean
Oracle Cloud
```

---

## Use Cases

Useful for:

- Passive reconnaissance
- Asset enrichment
- Cloud ownership identification
- Bug bounty recon workflows
- Security automation pipelines

---

# Installation

## Requirements

Install Go:

```bash
sudo apt install golang
```

Check:

```bash
go version
```

---

# Build From Source

Clone the repository:

```bash
git clone https://github.com/Sanushi-Salgado/cloudip-cli.git
cd cloudip-cli
```

Fix line endings if required:

```bash
dos2unix main.go
```

Build:

```bash
go build -o cloudip-cli
```

Test:

```bash
./cloudip-cli <IP address>
```

---

# Install Globally

Copy the binary into your system PATH:

```bash
sudo cp ~/cloudip-cli/cloudip-cli /usr/local/bin/
```

Make executable:

```bash
sudo chmod +x /usr/local/bin/cloudip-cli
```

Verify:

```bash
which cloudip-cli
```

Expected:

```
/usr/local/bin/cloudip-cli
```

Now run from anywhere:

```bash
cloudip-cli <IP address>
```

---

# Usage

Basic lookup:

```bash
cloudip-cli <IP address>
```

---

# Project structure

```
cloudip-cli/
|
├── main.go
├── go.mod
├── go.sum
└── cloudip-cli
```

Install dependency:

```bash
go get github.com/rezmoss/go-cloudip@latest
```

Build:

```bash
go build -o cloudip-cli
```

---

# Limitations

- Detection depends on the upstream cloud IP database.
- Not every hosting provider publishes complete IP ranges.
- Absence of a match does not mean the IP is not hosted in the cloud.

---

# Credits

Library:

- https://github.com/rezmoss/go-cloudip

Cloud IP data:

- https://github.com/rezmoss/cloudip-db

---

# ⚠️ Security Notice

This tool is intended for authorized security testing and educational purposes only.

- Do not use against systems without permission
- Follow responsible disclosure practices
- Respect program policies