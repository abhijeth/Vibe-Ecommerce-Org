# Vibe E-commerce - Version 1 (Educational Vulnerable Application)

## 🚨 **CRITICAL SECURITY WARNING**

**⚠️ This application is intentionally vulnerable for educational purposes only!**

**DO NOT deploy in production environments!**

## 📖 Introduction

This vulnerable e-commerce application was **generated using Cursor AI** and is intentionally designed with multiple security vulnerabilities for educational purposes. The code demonstrates common web application security flaws that are frequently found in real-world applications.

### ⚠️ **Important Warning**
- **This application contains intentional security vulnerabilities**
- **It is designed for educational and testing purposes only**
- **Never use in production environments**
- **Never use real personal or financial data**
- **Always use in isolated, controlled environments**

## 🎯 Educational Purpose

This application is designed for:
- **Security Training**: Learning about common web vulnerabilities
- **Penetration Testing Practice**: Hands-on exploitation experience
- **Security Awareness**: Understanding real-world security risks
- **Academic Use**: Teaching cybersecurity concepts

## 🔧 Prerequisites

Before installing, ensure you have the following:

- **Go 1.21+** installed on your system
  - Download from: https://golang.org/dl/
  - Verify installation: `go version`
- **Git** for cloning the repository
  - Download from: https://git-scm.com/downloads
- **Web browser** (Chrome, Firefox, Safari, Edge)
- **Terminal/Command Prompt** access

## 📥 Steps to Install

### Method 1: Clone from GitHub (Recommended)
```bash
# Clone the repository
git clone https://github.com/yourusername/vibe-ecommerce.git

# Navigate to the directory
cd vibe-ecommerce

# Install Go dependencies
go mod tidy
```

### Method 2: Download ZIP
```bash
# Download and extract the ZIP file
# Navigate to the extracted directory
cd vibe-ecommerce

# Install Go dependencies
go mod tidy
```

### Method 3: Using Cursor IDE
1. **Open Cursor IDE**
2. **Clone Repository**:
   - Press `Ctrl+Shift+P` (Windows/Linux) or `Cmd+Shift+P` (Mac)
   - Type "Git: Clone" and select it
   - Enter repository URL: `https://github.com/yourusername/vibe-ecommerce.git`
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
cd vibe-ecommerce

# Run the application
go run main.go
```

### 2. Expected Output
You should see output similar to:
```
🚀 Vibe E-commerce Version 1 starting on http://localhost:8083
⚠️  WARNING: This is an intentionally vulnerable application for educational purposes!
📚 Test accounts available - see home page for credentials
```

### 3. Access in Browser
1. **Open your web browser**
2. **Navigate to the URL shown in terminal** (e.g., `http://localhost:8083`)
3. **You should see the Vibe Shop home page**

### 4. Test the Application
1. **Browse the home page** to see available features
2. **Use test accounts** to explore different user roles:
   - Customer: `alice@example.com` / `insecurepass1`
   - Admin: `admin@example.com` / `adminpass`
   - Owner: `owner@example.com` / `ownerpass`
3. **Explore vulnerabilities** as documented in the security sections

## 🔧 Troubleshooting

### Common Issues and Solutions

#### ❌ **Port Already in Use**
**Error**: `bind: address already in use`

**Solutions**:
```bash
# Method 1: Kill existing processes
pkill -f "go run main.go"

# Method 2: Find and kill specific process
lsof -ti:8083 | xargs kill -9

# Method 3: Use different port
export PORT=8084
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

#### ❌ **Templates Not Found**
**Error**: Template loading errors

**Solutions**:
```bash
# Verify templates directory exists
ls -la templates/

# Check file permissions
chmod -R 755 templates/

# Ensure you're running from project root
pwd  # Should show path ending with vibe-ecommerce
```

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

## 🚨 **Security Vulnerabilities**

### 1. **Plain Text Password Storage** 🔓
- **Location**: User authentication system
- **Impact**: Complete account compromise if database is breached
- **Severity**: Critical

### 2. **PII Exposure in Plain Text** 💳
- **Location**: Database storage
- **Impact**: Identity theft, financial fraud
- **Severity**: Critical

### 3. **SQL Injection Vulnerabilities** 💉
- **Location**: Login and other database queries
- **Impact**: Database compromise, data theft
- **Severity**: Critical

### 4. **No Input Validation** ��
- **Location**: All user input fields
- **Impact**: XSS, code injection, data corruption
- **Severity**: High

### 5. **Weak Session Management** 🍪
- **Location**: Session handling
- **Impact**: Session hijacking, account takeover
- **Severity**: High

### 6. **Authorization Bypass** 👑
- **Location**: Admin/Owner middleware
- **Impact**: Privilege escalation
- **Severity**: Critical

## 🧪 **Exploitation Examples**

### SQL Injection Testing
```bash
# Test SQL injection in login
curl -X POST http://localhost:8083/login \
  -d "email=admin@example.com' OR '1'='1" \
  -d "password=anything"
```

### Authorization Bypass Testing
```bash
# Login as regular user, then access admin functions
# 1. Login as alice@example.com
# 2. Access admin dashboard: http://localhost:8083/admin
# 3. Access owner dashboard: http://localhost:8083/owner
```

## 📚 **Learning Resources**

### Security Concepts Demonstrated
- **OWASP Top 10**: Multiple vulnerabilities from the OWASP Top 10
- **Authentication Failures**: Weak password handling
- **Authorization Flaws**: Missing access controls
- **Input Validation**: Insufficient input validation
- **Data Exposure**: Sensitive data disclosure
- **SQL Injection**: Database query vulnerabilities
- **Session Management**: Weak session handling

### Recommended Reading
- OWASP Top 10 Web Application Security Risks
- Web Application Security Testing Guide
- Secure Coding Practices
- Penetration Testing Methodologies

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

**Version**: 1.0  
**Last Updated**: September 2024  
**Security Classification**: Educational/Testing Only  
**Production Status**: ❌ NOT FOR PRODUCTION USE  

## 🎯 **Next Steps**

After exploring Version 1, check out:
- **Version 2**: Enhanced vulnerabilities with command injection and file exposure
- **Version 3**: Maximum vulnerability with "insecure by default" design

Each version demonstrates different aspects of web application security vulnerabilities.
