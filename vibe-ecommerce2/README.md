# Vibe E-commerce - Version 2 (Educational Vulnerable Application)

## 🚨 **CRITICAL SECURITY WARNING**

**⚠️ This application is intentionally vulnerable for educational purposes only!**

**DO NOT deploy in production environments!**

## 📖 Introduction

This vulnerable e-commerce application was **generated using Cursor AI** and is intentionally designed with advanced security vulnerabilities for educational purposes. Version 2 builds upon Version 1's vulnerabilities and adds critical Remote Code Execution (RCE) capabilities, file exposure, and IDOR attacks.

### ⚠️ **Important Warning**
- **This application contains intentional security vulnerabilities**
- **It includes command injection (RCE) capabilities**
- **It is designed for educational and testing purposes only**
- **Never use in production environments**
- **Never use real personal or financial data**
- **Always use in isolated, controlled environments**

## 🎯 Educational Purpose

This application is designed for:
- **Advanced Security Training**: Learning about RCE and file system vulnerabilities
- **Penetration Testing Practice**: Hands-on exploitation of critical vulnerabilities
- **Security Awareness**: Understanding the impact of command injection
- **Academic Use**: Teaching advanced cybersecurity concepts

## 🔧 Prerequisites

Before installing, ensure you have the following:

- **Go 1.21+** installed on your system
  - Download from: https://golang.org/dl/
  - Verify installation: `go version`
- **Git** for cloning the repository
  - Download from: https://git-scm.com/downloads
- **Web browser** (Chrome, Firefox, Safari, Edge)
- **Terminal/Command Prompt** access
- **curl** (optional) for command-line testing

## 📥 Steps to Install

### Method 1: Clone from GitHub (Recommended)
```bash
# Clone the repository
git clone https://github.com/yourusername/vibe-ecommerce2.git

# Navigate to the directory
cd vibe-ecommerce2

# Install Go dependencies
go mod tidy
```

### Method 2: Download ZIP
```bash
# Download and extract the ZIP file
# Navigate to the extracted directory
cd vibe-ecommerce2

# Install Go dependencies
go mod tidy
```

### Method 3: Using Cursor IDE
1. **Open Cursor IDE**
2. **Clone Repository**:
   - Press `Ctrl+Shift+P` (Windows/Linux) or `Cmd+Shift+P` (Mac)
   - Type "Git: Clone" and select it
   - Enter repository URL: `https://github.com/yourusername/vibe-ecommerce2.git`
   - Choose a local directory
3. **Open Terminal in Cursor**:
   - Press `Ctrl+`` (backtick) or `View > Terminal`
4. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

## 🚀 Steps to Run in Browser

### 1. Start the Application
```bash
# Navigate to the project directory
cd vibe-ecommerce2

# Run the application
go run main.go
```

### 2. Expected Output
You should see output similar to:
```
🚀 Vibe E-commerce Version 2 starting on http://localhost:9001
⚠️  WARNING: This is an intentionally vulnerable application for educational purposes!
🎯 Version 2 includes: Command Injection, File Exposure, IDOR, and Secret Flag Challenge
📚 Test accounts available - see home page for credentials
🚩 Secret Flag Challenge: /secret endpoint
```

### 3. Access in Browser
1. **Open your web browser**
2. **Navigate to the URL shown in terminal** (e.g., `http://localhost:9001`)
3. **You should see the Vibe Shop home page**

### 4. Test the Application
1. **Browse the home page** to see available features
2. **Use test accounts** to explore different user roles:
   - Customer: `alice@example.com` / `insecurepass1`
   - Customer: `bob@example.com` / `bobpass123`
   - Admin: `admin@example.com` / `adminpass`
   - Owner: `owner@example.com` / `ownerpass`
3. **Explore vulnerabilities** including the Secret Flag Challenge

## 🔧 Troubleshooting

### Common Issues and Solutions

#### ❌ **Port Already in Use**
**Error**: `bind: address already in use`

**Solutions**:
```bash
# Method 1: Kill existing processes
pkill -f "go run main.go"

# Method 2: Find and kill specific process
lsof -ti:9001 | xargs kill -9

# Method 3: Use different port
export PORT=9002
go run main.go
```

#### ❌ **Go Not Found**
**Error**: `go: command not found`

**Solutions**:
```bash
# Install Go from https://golang.org/dl/
# Add Go to PATH environment variable
# Restart terminal/command prompt
# Verify: go version
```

#### ❌ **Dependencies Not Found**
**Error**: `cannot find module`

**Solutions**:
```bash
# Clean module cache
go clean -modcache

# Reinstall dependencies
go mod tidy

# Verify go.mod file exists
ls -la go.mod
```

#### ❌ **Database Issues**
**Error**: Database connection problems

**Solutions**:
```bash
# Reset database
rm -f database/ecommerce.db
go run main.go

# Check database directory exists
mkdir -p database
```

#### ❌ **Permission Denied**
**Error**: Permission denied errors

**Solutions**:
```bash
# On Linux/Mac: Fix permissions
chmod +x main.go

# On Windows: Run as Administrator
# Or use PowerShell as Administrator
```

#### ❌ **Browser Can't Connect**
**Error**: Browser shows "This site can't be reached"

**Solutions**:
1. **Check if server is running**:
   ```bash
   # Look for "Server starting on" message
   # Check terminal for any error messages
   ```

2. **Verify port number**:
   ```bash
   # Check what port the server is using
   # Make sure you're using the correct URL
   ```

3. **Check firewall settings**:
   - Allow Go applications through firewall
   - Temporarily disable firewall for testing

4. **Try different browser**:
   - Test with Chrome, Firefox, Safari, or Edge
   - Clear browser cache

#### ❌ **Command Injection Not Working**
**Error**: Command injection attempts fail

**Solutions**:
1. **Check search functionality**:
   - Navigate to `/products` page
   - Try searching for `ls` or `whoami`
   - Check browser developer tools for errors

2. **Verify endpoint access**:
   ```bash
   # Test with curl
   curl "http://localhost:9001/products?search=ls"
   ```

3. **Check system permissions**:
   - Ensure Go has permission to execute commands
   - Check if antivirus is blocking command execution

#### ❌ **Secret Flag Challenge Not Accessible**
**Error**: `/secret` endpoint not working

**Solutions**:
1. **Verify endpoint exists**:
   ```bash
   # Test with curl
   curl "http://localhost:9001/secret"
   ```

2. **Check route configuration**:
   - Ensure the route is properly registered
   - Check for any middleware blocking access

3. **Try different browsers**:
   - Some browsers may block certain content types

### 🔍 **Debugging Tips**

#### Enable Verbose Logging
```bash
# Run with verbose output
go run main.go -v

# Or add debug prints in code
```

#### Check System Resources
```bash
# Check available ports
netstat -an | grep LISTEN

# Check system resources
top  # or htop on Linux
```

#### Verify Installation
```bash
# Check Go installation
go version

# Check Go environment
go env

# Check project structure
tree .  # or ls -la
```

## 🎯 **Quick Start Checklist**

- [ ] Go 1.21+ installed and in PATH
- [ ] Project cloned/downloaded
- [ ] Dependencies installed (`go mod tidy`)
- [ ] Application started (`go run main.go`)
- [ ] Browser opened to correct URL
- [ ] Home page loads successfully
- [ ] Test accounts work for login
- [ ] Secret Flag Challenge accessible at `/secret`

## 🚨 **Major Security Vulnerabilities**

### 1. **Command Injection (RCE)** 💥
- **Location**: Search functionality, debug endpoints
- **Impact**: Complete server takeover
- **Severity**: Critical

### 2. **File Exposure Vulnerability** 📁
- **Location**: `/secret` endpoint
- **Impact**: System file access, configuration disclosure
- **Severity**: Critical

### 3. **IDOR (Insecure Direct Object Reference)** 🔍
- **Location**: Profile access endpoints
- **Impact**: User data enumeration, privacy violation
- **Severity**: High

### 4. **Database Direct Access** 🗄️
- **Location**: `/owner/database` endpoint
- **Impact**: Complete database compromise
- **Severity**: Critical

### 5. **Authentication Bypass** 🔓
- **Location**: Multiple admin/owner endpoints
- **Impact**: Privilege escalation
- **Severity**: Critical

### 6. **PII Exposure** 💳
- **Location**: Orders and payment endpoints
- **Impact**: Identity theft, financial fraud
- **Severity**: Critical

## 🧪 **Exploitation Examples**

### Command Injection Testing
```bash
# Test basic command injection
curl "http://localhost:9001/products?search=ls"

# Test system information
curl "http://localhost:9001/products?search=whoami"

# Test file system access
curl "http://localhost:9001/products?search=cat%20/etc/passwd"
```

### File Exposure Testing
```bash
# Access the file explorer
curl "http://localhost:9001/secret"

# Direct access to flag
curl "http://localhost:9001/secret?file=flag.txt"
```

### IDOR Testing
```bash
# Access customer profile
curl "http://localhost:9001/profile?id=1"

# Access admin profile
curl "http://localhost:9001/profile?id=1000"
```

## 🚩 **Secret Flag Challenge**

### Challenge Details
- **Flag**: `0b8dcaf09bee1fd3c2143ea0ffba4fa2`
- **Source**: MD5 hash of "Vib3Sec"
- **Location**: Hidden in file exposure vulnerability

### Access the Challenge
```
http://localhost:9001/secret
```

### Exploitation Steps
1. Navigate to `/secret`
2. Explore available files
3. Find the flag in `flag.txt`

## 📚 **Learning Resources**

### Security Concepts Demonstrated
- **OWASP Top 10**: Multiple vulnerabilities from the OWASP Top 10
- **Command Injection**: OS command execution vulnerabilities
- **File Exposure**: System file access vulnerabilities
- **IDOR**: Insecure Direct Object Reference
- **Authentication Bypass**: Weak authentication mechanisms
- **Authorization Flaws**: Missing access controls
- **Input Validation**: Insufficient input validation
- **Data Exposure**: Sensitive data disclosure

### Recommended Reading
- OWASP Top 10 Web Application Security Risks
- Web Application Security Testing Guide
- Secure Coding Practices
- Penetration Testing Methodologies
- Command Injection Prevention
- File System Security

## 🤝 **Contributing**

### Adding Vulnerabilities
1. Fork the repository
2. Add new vulnerable functionality
3. Document the vulnerability
4. Create exploitation examples
5. Submit a pull request

### Reporting Issues
- Use GitHub Issues for bug reports
- Provide detailed reproduction steps
- Include expected vs actual behavior

## 📄 **License**

This project is for educational purposes only. Use at your own risk in controlled environments.

## 📞 **Support**

For questions or issues:
- Check the troubleshooting section above
- Review the code comments
- Consult security documentation
- Use in isolated testing environments only

---

**Version**: 2.0  
**Last Updated**: September 2024  
**Security Classification**: Educational/Testing Only  
**Production Status**: ❌ NOT FOR PRODUCTION USE  

## 🎯 **Next Steps**

After exploring Version 2, check out:
- **Version 1**: Basic vulnerabilities with fundamental security failures
- **Version 3**: Maximum vulnerability with "insecure by default" design

Each version demonstrates different aspects of web application security vulnerabilities.
