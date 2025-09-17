# Vibe-Ecommerce Organization

## 🚨 **CRITICAL SECURITY WARNING**

**⚠️ All applications in this organization are intentionally vulnerable for educational purposes only!**

**DO NOT deploy any of these applications in production environments!**

## 📖 About This Organization

The **Vibe-Ecommerce** organization contains a series of intentionally vulnerable e-commerce applications designed for **educational and security training purposes**. These applications were generated using **Cursor AI** and demonstrate various levels of security vulnerabilities commonly found in real-world web applications.

### 🎯 **Educational Mission**
- **Security Training**: Learn about web application vulnerabilities
- **Penetration Testing Practice**: Hands-on exploitation experience
- **Security Awareness**: Understand real-world security risks
- **Academic Use**: Teaching cybersecurity concepts
- **Research**: Vulnerability analysis and security research

## 🏪 **Available Applications**

This organization contains **three progressively vulnerable versions** of the same e-commerce application, each demonstrating different aspects of web security vulnerabilities:

---

## 📦 **Version 1: Basic Vulnerabilities**
**Repository**: [`vibe-ecommerce`](https://github.com/abhijeth/Vibe-Ecommerce-Org/tree/main/vibe-ecommerce)  
**Port Range**: 8083-9000  
**Focus**: Fundamental security failures

### 🎯 **What You'll Learn**
- **Plain text password storage** vulnerabilities
- **PII exposure** in database storage
- **SQL injection** through string concatenation
- **No input validation** leading to XSS risks
- **Weak session management** vulnerabilities
- **Authorization bypass** in admin/owner functions

### 🚨 **Key Vulnerabilities**
| Vulnerability | Severity | Impact |
|---------------|----------|---------|
| Plain Text Passwords | 🔴 Critical | Complete account compromise |
| PII in Plain Text | 🔴 Critical | Identity theft, financial fraud |
| SQL Injection | 🔴 Critical | Database compromise |
| No Input Validation | 🟠 High | XSS, code injection |
| Weak Sessions | 🟠 High | Session hijacking |
| Auth Bypass | 🔴 Critical | Privilege escalation |

### 🎓 **Best For**
- **Beginners** in web security
- **Understanding fundamental** security failures
- **Learning basic exploitation** techniques
- **Academic courses** on web security

---

## 📦 **Version 2: Advanced Vulnerabilities**
**Repository**: [`vibe-ecommerce2`](https://github.com/abhijeth/Vibe-Ecommerce-Org/tree/main/vibe-ecommerce2)  
**Port Range**: 9001-10000  
**Focus**: Critical vulnerabilities with RCE capabilities

### 🎯 **What You'll Learn**
- **All Version 1 vulnerabilities** PLUS:
- **Command injection (RCE)** - Remote Code Execution
- **File exposure vulnerabilities** - System file access
- **IDOR attacks** - Insecure Direct Object Reference
- **Database direct access** - Complete database compromise
- **Secret Flag Challenge** - Educational CTF-style challenge

### 🚨 **Key Vulnerabilities**
| Vulnerability | Severity | Impact |
|---------------|----------|---------|
| Command Injection (RCE) | 🔴 Critical | Complete server takeover |
| File Exposure | 🔴 Critical | System file access |
| IDOR | 🟠 High | User data enumeration |
| Database Direct Access | 🔴 Critical | Complete database control |
| Authentication Bypass | 🔴 Critical | Privilege escalation |
| PII Exposure | 🔴 Critical | Identity theft |

### 🎓 **Best For**
- **Intermediate** security practitioners
- **Learning RCE** exploitation techniques
- **Understanding file system** vulnerabilities
- **Penetration testing** practice
- **CTF-style challenges**

### 🚩 **Secret Flag Challenge**
- **Flag**: `0b8dcaf09bee1fd3c2143ea0ffba4fa2`
- **Location**: `/secret` endpoint
- **Challenge**: Find the flag through file exposure vulnerability

---

## 📦 **Version 3: Maximum Vulnerabilities**
**Repository**: [`vibe-ecommerce3`](https://github.com/abhijeth/Vibe-Ecommerce-Org/tree/main/vibe-ecommerce3)  
**Port Range**: 10001-11000  
**Focus**: "Insecure by default" design with broken security features

### 🎯 **What You'll Learn**
- **All previous vulnerabilities** PLUS:
- **Complete authentication bypass** - No login required
- **Role bypass** - Any user can be admin/owner
- **Broken security features** - MFA and rate limiting that don't work
- **Maximum PII exposure** - All data visible to everyone
- **"Insecure by default"** design patterns

### 🚨 **Key Vulnerabilities**
| Vulnerability | Severity | Impact |
|---------------|----------|---------|
| Complete Auth Bypass | 🔴 Critical | No authentication required |
| Role Bypass | 🔴 Critical | Any user → Admin/Owner |
| Broken Security Features | 🟠 High | False sense of security |
| Maximum PII Exposure | 🔴 Critical | Complete identity theft |
| No Input Validation | 🟠 High | Multiple attack vectors |
| SQL Injection | 🔴 Critical | Database compromise |

### 🎓 **Best For**
- **Advanced** security practitioners
- **Understanding "insecure by default"** design
- **Learning bypass techniques** for security features
- **Security architecture** analysis
- **False security** recognition

---

## 🚀 **Quick Start Guide**

### 1. **Choose Your Version**
- **Beginner**: Start with [Version 1](https://github.com/abhijeth/Vibe-Ecommerce-Org/tree/main/vibe-ecommerce)
- **Intermediate**: Try [Version 2](https://github.com/abhijeth/Vibe-Ecommerce-Org/tree/main/vibe-ecommerce2)
- **Advanced**: Explore [Version 3](https://github.com/abhijeth/Vibe-Ecommerce-Org/tree/main/vibe-ecommerce3)

### 2. **Installation Steps**
```bash
# Clone the repository
git clone https://github.com/abhijeth/Vibe-Ecommerce-Org/[repository-name].git

# Navigate to directory
cd [repository-name]

# Install dependencies
go mod tidy

# Run the application
go run main.go
```

### 3. **Access in Browser**
- **Version 1**: `http://localhost:8083` (port range: 8083-9000)
- **Version 2**: `http://localhost:9001` (port range: 9001-10000)
- **Version 3**: `http://localhost:10001` (port range: 10001-11000)

### 4. **Test Accounts** (All Versions)
| Role | Email | Password |
|------|-------|----------|
| Customer | `alice@example.com` | `insecurepass1` |
| Customer | `bob@example.com` | `insecurepass2` |
| Admin | `admin@example.com` | `adminpass` |
| Owner | `owner@example.com` | `ownerpass` |

## 📊 **Version Comparison Matrix**

| Feature | Version 1 | Version 2 | Version 3 |
|---------|-----------|-----------|-----------|
| **Port Range** | 8083-9000 | 9001-10000 | 10001-11000 |
| **Password Hashing** | ❌ Plain text | ❌ Plain text | ✅ bcrypt (bypassed) |
| **Command Injection** | ❌ No | ✅ Yes | ❌ No |
| **File Exposure** | ❌ No | ✅ Yes | ❌ No |
| **IDOR** | ❌ No | ✅ Yes | ❌ No |
| **Auth Bypass** | ❌ No | ✅ Partial | ✅ Complete |
| **PII Exposure** | ✅ Basic | ✅ Enhanced | ✅ Maximum |
| **MFA** | ❌ No | ❌ No | ✅ Yes (broken) |
| **Rate Limiting** | ❌ No | ❌ No | ✅ Yes (bypassed) |
| **Secret Flag** | ❌ No | ✅ Yes | ❌ No |
| **Difficulty** | 🟢 Beginner | 🟡 Intermediate | 🔴 Advanced |

## 🎯 **Learning Path Recommendations**

### 🟢 **Beginner Path**
1. **Start with Version 1**
   - Learn basic vulnerabilities
   - Understand SQL injection
   - Practice authentication bypass
2. **Progress to Version 2**
   - Explore command injection
   - Try the Secret Flag Challenge
   - Learn file exposure techniques

### 🟡 **Intermediate Path**
1. **Begin with Version 2**
   - Master RCE techniques
   - Complete the Secret Flag Challenge
   - Understand IDOR attacks
2. **Advance to Version 3**
   - Learn bypass techniques
   - Understand false security
   - Analyze insecure design patterns

### 🔴 **Advanced Path**
1. **Start with Version 3**
   - Master complete authentication bypass
   - Learn role escalation techniques
   - Understand "insecure by default" design
2. **Compare with Versions 1 & 2**
   - Analyze vulnerability evolution
   - Understand security feature failures
   - Learn defense in depth principles

## 🛡️ **Security Best Practices Demonstrated**

### ❌ **What NOT to Do** (Demonstrated in these apps)
- Store passwords in plain text
- Expose PII without encryption
- Use string concatenation for SQL queries
- Skip input validation
- Implement weak session management
- Allow privilege escalation
- Create "insecure by default" designs

### ✅ **What TO Do** (Learn from fixing these apps)
- Use proper password hashing (bcrypt, scrypt, Argon2)
- Encrypt sensitive data at rest and in transit
- Use parameterized queries
- Implement comprehensive input validation
- Use secure session management
- Implement proper authorization checks
- Follow secure design principles

## 🧪 **Exploitation Examples**

### Version 1 - Basic Exploitation
```bash
# SQL Injection
curl -X POST http://localhost:8083/login \
  -d "email=admin@example.com' OR '1'='1" \
  -d "password=anything"

# Authorization Bypass
# Login as regular user, access admin functions
```

### Version 2 - Advanced Exploitation
```bash
# Command Injection
curl "http://localhost:9001/products?search=ls"

# File Exposure
curl "http://localhost:9001/secret?file=flag.txt"

# IDOR
curl "http://localhost:9001/profile?id=1"
```

### Version 3 - Maximum Exploitation
```bash
# Complete Auth Bypass
curl "http://localhost:10001/admin/insecure"

# Role Bypass
# Any logged-in user can access admin/owner functions
```

## 📚 **Educational Resources**

### 🎓 **Academic Use**
- **Cybersecurity Courses**: Perfect for hands-on labs
- **Web Security Classes**: Real-world vulnerability examples
- **Penetration Testing Training**: Practical exploitation practice
- **Security Research**: Vulnerability analysis and defense

### 📖 **Recommended Reading**
- OWASP Top 10 Web Application Security Risks
- Web Application Security Testing Guide
- Secure Coding Practices
- Penetration Testing Methodologies
- "Insecure by Default" Design Patterns

### 🔗 **Related Resources**
- [OWASP Foundation](https://owasp.org/)
- [Web Security Academy](https://portswigger.net/web-security)
- [OWASP Testing Guide](https://owasp.org/www-project-web-security-testing-guide/)
- [Secure Coding Practices](https://owasp.org/www-project-secure-coding-practices-quick-reference-guide/)

## 🤝 **Contributing**

### 🐛 **Reporting Issues**
- Use GitHub Issues for bug reports
- Provide detailed reproduction steps
- Include expected vs actual behavior
- Specify which version you're testing

### 🔧 **Adding Vulnerabilities**
1. Fork the repository
2. Add new vulnerable functionality
3. Document the vulnerability
4. Create exploitation examples
5. Submit a pull request

### 📝 **Improving Documentation**
- Enhance README files
- Add more exploitation examples
- Improve troubleshooting guides
- Create video tutorials

## ⚖️ **Legal and Ethical Use**

### ✅ **Permitted Uses**
- Educational and training purposes
- Security research in controlled environments
- Academic coursework and labs
- Penetration testing practice
- Security awareness training

### ❌ **Prohibited Uses**
- Production deployment
- Real-world attacks
- Unauthorized testing
- Malicious activities
- Commercial use without permission

### 🚨 **Responsible Disclosure**
- Report vulnerabilities responsibly
- Use in isolated, controlled environments
- Follow ethical hacking principles
- Respect others' systems and data

## 📞 **Support and Community**

### 💬 **Getting Help**
- Check individual repository README files
- Review troubleshooting sections
- Consult security documentation
- Use GitHub Discussions for questions

### 🌐 **Community**
- Join security training communities
- Participate in CTF competitions
- Share learning experiences
- Contribute to security education

## 📄 **License**

This organization and all repositories are for **educational purposes only**. Use at your own risk in controlled environments.

## 🏆 **Acknowledgments**

- **Cursor AI** for code generation assistance
- **Security Community** for vulnerability research
- **Educational Institutions** for security training
- **Open Source Community** for tools and resources

---

**Organization**: Vibe-Ecommerce  
**Purpose**: Educational Security Training  
**Last Updated**: September 2024  
**Security Classification**: Educational/Testing Only  
**Production Status**: ❌ NOT FOR PRODUCTION USE  

## 🎯 **Get Started Now**

Choose your learning path and start exploring web application security vulnerabilities:

- 🟢 **[Version 1 - Basic](https://github.com/abhijeth/Vibe-Ecommerce-Org/tree/main/vibe-ecommerce)** - Start here for fundamentals
- 🟡 **[Version 2 - Advanced](https://github.com/abhijeth/Vibe-Ecommerce-Org/tree/main/vibe-ecommerce2)** - Challenge yourself with RCE
- 🔴 **[Version 3 - Maximum](https://github.com/abhijeth/Vibe-Ecommerce-Org/tree/main/vibe-ecommerce3)** - Master "insecure by default"

**Happy Learning and Stay Secure!** 🛡️
