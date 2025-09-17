# Vibe E-commerce - Version 3 (Educational Vulnerable Application)

## 🚨 **CRITICAL SECURITY WARNING**

**⚠️ This application is intentionally vulnerable for educational purposes only!**

**DO NOT deploy in production environments!**

## 📖 Introduction

This vulnerable e-commerce application was **generated using Cursor AI** and is intentionally designed with "insecure by default" principles for educational purposes. Version 3 represents the worst-case scenario of what happens when security is an afterthought and security features are implemented incorrectly.

### ⚠️ **Important Warning**
- **This application contains intentional security vulnerabilities**
- **It demonstrates "insecure by default" design patterns**
- **It includes broken security features that provide false security**
- **It is designed for educational and testing purposes only**
- **Never use in production environments**
- **Never use real personal or financial data**
- **Always use in isolated, controlled environments**

## 🎯 Educational Purpose

This application is designed for:
- **Advanced Security Training**: Understanding "insecure by default" design
- **Penetration Testing Practice**: Exploiting broken security implementations
- **Security Awareness**: Learning about false security and bypass techniques
- **Academic Use**: Teaching cybersecurity concepts and secure design principles

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
git clone https://github.com/yourusername/vibe-ecommerce3.git

# Navigate to the directory
cd vibe-ecommerce3

# Install Go dependencies
go mod tidy
```

### Method 2: Download ZIP
```bash
# Download and extract the ZIP file
# Navigate to the extracted directory
cd vibe-ecommerce3

# Install Go dependencies
go mod tidy
```

### Method 3: Using Cursor IDE
1. **Open Cursor IDE**
2. **Clone Repository**:
   - Press `Ctrl+Shift+P` (Windows/Linux) or `Cmd+Shift+P` (Mac)
   - Type "Git: Clone" and select it
   - Enter repository URL: `https://github.com/yourusername/vibe-ecommerce3.git`
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
cd vibe-ecommerce3

# Run the application
go run main.go
```

### 2. Expected Output
You should see output similar to:
```
🚀 Vibe E-commerce Version 3 starting on http://localhost:10001
⚠️  WARNING: This is an intentionally vulnerable application for educational purposes!
💥 Version 3: Insecure by Default - Maximum vulnerabilities with broken security features
🔓 Complete authentication bypass - any user can access admin/owner functions
📚 Test accounts available - see home page for credentials
```

### 3. Access in Browser
1. **Open your web browser**
2. **Navigate to the URL shown in terminal** (e.g., `http://localhost:10001`)
3. **You should see the Vibe Shop home page**

### 4. Test the Application
1. **Browse the home page** to see available features
2. **Use test accounts** to explore different user roles:
   - Customer: `alice@example.com` / `insecurepass1`
   - Customer: `bob@example.com` / `insecurepass2`
   - Admin: `admin@example.com` / `adminpass`
   - Owner: `owner@example.com` / `ownerpass`
3. **Explore vulnerabilities** including complete authentication bypass

## 🔧 Troubleshooting

### Common Issues and Solutions

#### ❌ **Port Already in Use**
**Error**: `bind: address already in use`

**Solutions**:
```bash
# Method 1: Kill existing processes
pkill -f "go run main.go"

# Method 2: Find and kill specific process
lsof -ti:10001 | xargs kill -9

# Method 3: Use different port
export PORT=10002
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

#### ❌ **Authentication Bypass Not Working**
**Error**: Admin/Owner functions still require login

**Solutions**:
1. **Check insecure endpoints**:
   ```bash
   # Test with curl
   curl "http://localhost:10001/admin/insecure"
   curl "http://localhost:10001/owner/debug"
   ```

2. **Verify route configuration**:
   - Ensure insecure routes are properly registered
   - Check for any middleware blocking access

3. **Test with different browsers**:
   - Some browsers may cache authentication states

#### ❌ **MFA/Rate Limiting Issues**
**Error**: Security features not working as expected

**Solutions**:
1. **This is intentional** - the security features are broken by design
2. **Test bypass techniques**:
   - Try different input formats
   - Test with various user agents
   - Check for timing attacks

3. **Review code comments**:
   - Look for bypass techniques documented in code

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
- [ ] Authentication bypass works (admin/owner without login)

## 🚨 **Major Security Vulnerabilities**

### 1. **Complete Authentication Bypass** 🔓
- **Location**: All admin/owner endpoints
- **Impact**: Complete system compromise
- **Severity**: Critical

### 2. **Role Bypass** 👑
- **Location**: Authorization middleware
- **Impact**: Privilege escalation
- **Severity**: Critical

### 3. **Broken Security Features** 🛡️
- **Location**: MFA, rate limiting implementations
- **Impact**: False sense of security
- **Severity**: High

### 4. **Maximum PII Exposure** 💳
- **Location**: All user data endpoints
- **Impact**: Complete identity theft
- **Severity**: Critical

### 5. **No Input Validation** 🚫
- **Location**: All forms and inputs
- **Impact**: Multiple attack vectors
- **Severity**: High

### 6. **SQL Injection** 💉
- **Location**: Database queries
- **Impact**: Database compromise
- **Severity**: Critical

### 7. **Plain Text Passwords** 🔐
- **Location**: User authentication
- **Impact**: Account compromise
- **Severity**: Critical

## 🧪 **Exploitation Examples**

### Complete Authentication Bypass
```bash
# Access admin functions without login
curl http://localhost:10001/admin/insecure

# Access owner functions without login
curl http://localhost:10001/owner/debug

# Access payment data without login
curl http://localhost:10001/payments/insecure
```

### Role Bypass Testing
```bash
# Login as regular user, then access admin functions
# 1. Login as alice@example.com
# 2. Access admin dashboard: http://localhost:10001/admin
# 3. Access owner dashboard: http://localhost:10001/owner
# 4. Any logged-in user can access these functions
```

### PII Exposure Testing
```bash
# Access orders without login (should show PII)
curl http://localhost:10001/orders

# Access specific order details
curl http://localhost:10001/order/1

# Access payment data
curl http://localhost:10001/payments/insecure
```

### SQL Injection Testing
```bash
# Test SQL injection in login
curl -X POST http://localhost:10001/login \
  -d "email=alice@example.com' OR '1'='1" \
  -d "password=anything"
```

## 📚 **Learning Resources**

### Security Concepts Demonstrated
- **OWASP Top 10**: Multiple vulnerabilities from the OWASP Top 10
- **"Insecure by Default"**: Design patterns that lead to vulnerabilities
- **False Security**: Security features that don't work
- **Authentication Bypass**: Complete authentication failures
- **Authorization Flaws**: Missing access controls
- **Input Validation**: Insufficient input validation
- **Data Exposure**: Maximum sensitive data disclosure
- **SQL Injection**: Database query vulnerabilities

### Recommended Reading
- OWASP Top 10 Web Application Security Risks
- Web Application Security Testing Guide
- Secure Coding Practices
- Penetration Testing Methodologies
- "Insecure by Default" Design Patterns
- Security Feature Implementation Best Practices

## 💡 **Key Learning Points**

### "Insecure by Default" Design
- Security features implemented but bypassable
- False sense of security
- Multiple attack vectors
- Complete system compromise possible

### Real-World Implications
- This represents what happens when security is an afterthought
- Common in poorly designed applications
- Demonstrates the importance of secure design principles
- Shows why defense in depth is crucial

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

**Version**: 3.0  
**Last Updated**: September 2024  
**Security Classification**: Educational/Testing Only  
**Production Status**: ❌ NOT FOR PRODUCTION USE  

## �� **Next Steps**

After exploring Version 3, check out:
- **Version 1**: Basic vulnerabilities with fundamental security failures
- **Version 2**: Advanced vulnerabilities with command injection and file exposure

Each version demonstrates different aspects of web application security vulnerabilities.
