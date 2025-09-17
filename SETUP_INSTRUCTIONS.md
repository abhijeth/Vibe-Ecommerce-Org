# Vibe-Ecommerce Organization Setup Instructions

## 🚀 **GitHub Organization Setup**

### 1. Create GitHub Organization
1. Go to https://github.com/organizations/new
2. Create organization named: **Vibe-Ecommerce**
3. Set visibility to **Public**
4. Add description: "Educational vulnerable e-commerce applications for security training"

### 2. Create Three Repositories

#### Repository 1: vibe-ecommerce (Version 1)
- **Name**: `vibe-ecommerce`
- **Description**: "Version 1: Basic vulnerabilities - Plain text passwords, SQL injection, PII exposure"
- **Visibility**: Public
- **Topics**: `security`, `vulnerable`, `educational`, `web-security`, `penetration-testing`, `sql-injection`, `authentication-bypass`

#### Repository 2: vibe-ecommerce2 (Version 2)
- **Name**: `vibe-ecommerce2`
- **Description**: "Version 2: Advanced vulnerabilities - Command injection (RCE), file exposure, IDOR, Secret Flag Challenge"
- **Visibility**: Public
- **Topics**: `security`, `vulnerable`, `educational`, `web-security`, `penetration-testing`, `command-injection`, `rce`, `file-exposure`, `idor`, `ctf`

#### Repository 3: vibe-ecommerce3 (Version 3)
- **Name**: `vibe-ecommerce3`
- **Description**: "Version 3: Maximum vulnerabilities - Complete auth bypass, role bypass, broken security features"
- **Visibility**: Public
- **Topics**: `security`, `vulnerable`, `educational`, `web-security`, `penetration-testing`, `authentication-bypass`, `insecure-by-default`, `broken-security`

### 3. Upload Files to Each Repository

For each repository, upload the following files:
- `main.go` (improved with comments)
- `README.md` (comprehensive documentation)
- `go.mod`
- `go.sum`
- `.gitignore`
- `templates/` directory
- `static/` directory

### 4. Set Organization README

Upload the `README.md` file from this directory to the organization's main page.

## 📁 **File Structure for Upload**

```
Vibe-Ecommerce-Org/
├── README.md                    # Organization README
├── SETUP_INSTRUCTIONS.md        # This file
├── vibe-ecommerce/              # Version 1 files
│   ├── main.go
│   ├── README.md
│   ├── go.mod
│   ├── go.sum
│   ├── .gitignore
│   ├── templates/
│   └── static/
├── vibe-ecommerce2/             # Version 2 files
│   ├── main.go
│   ├── README.md
│   ├── go.mod
│   ├── go.sum
│   ├── .gitignore
│   ├── templates/
│   └── static/
└── vibe-ecommerce3/             # Version 3 files
    ├── main.go
    ├── README.md
    ├── go.mod
    ├── go.sum
    ├── .gitignore
    ├── templates/
    └── static/
```

## 🎯 **Organization Features to Enable**

1. **Issues**: Enable for bug reports and feature requests
2. **Discussions**: Enable for community questions
3. **Projects**: Enable for tracking development
4. **Wiki**: Enable for additional documentation
5. **Security**: Enable for vulnerability reporting

## 🏷️ **Repository Topics to Add**

### Common Topics (All Repositories)
- `security`
- `vulnerable`
- `educational`
- `web-security`
- `penetration-testing`
- `go`
- `golang`
- `ecommerce`
- `web-application`
- `security-training`

### Version-Specific Topics
- **Version 1**: `sql-injection`, `authentication-bypass`, `pii-exposure`
- **Version 2**: `command-injection`, `rce`, `file-exposure`, `idor`, `ctf`
- **Version 3**: `authentication-bypass`, `insecure-by-default`, `broken-security`

## 📋 **Organization Settings**

### General Settings
- **Organization name**: Vibe-Ecommerce
- **Contact email**: [Your email]
- **Location**: [Your location]
- **Website**: https://github.com/Vibe-Ecommerce

### Member Privileges
- **Base permissions**: Read
- **Repository creation**: Member
- **Repository forking**: Allow

### Security Settings
- **Two-factor authentication**: Required
- **SSH certificate authorities**: Configure if needed

## 🚀 **Quick Setup Commands**

After creating the organization and repositories, use these commands to upload files:

```bash
# For each repository, clone and upload files
git clone https://github.com/Vibe-Ecommerce/vibe-ecommerce.git
cd vibe-ecommerce
# Copy files from local directories
git add .
git commit -m "Initial commit: Version 1 with basic vulnerabilities"
git push origin main

# Repeat for vibe-ecommerce2 and vibe-ecommerce3
```

## 📝 **Post-Setup Checklist**

- [ ] Organization created with proper name and description
- [ ] Three repositories created with appropriate descriptions
- [ ] All files uploaded to respective repositories
- [ ] Organization README uploaded
- [ ] Repository topics added
- [ ] Issues and Discussions enabled
- [ ] Security settings configured
- [ ] All repositories are public and accessible
- [ ] Test each application runs correctly
- [ ] Verify all links in README files work

## 🎯 **Success Criteria**

The organization setup is complete when:
1. All three repositories are accessible and functional
2. Each application runs without errors
3. All vulnerabilities are preserved and documented
4. README files provide clear instructions
5. Organization page clearly differentiates between versions
6. All security warnings are prominently displayed
