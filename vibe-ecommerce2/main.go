/*
 * Vibe E-commerce - Version 2 (Educational Vulnerable Application)
 * 
 * ⚠️  SECURITY WARNING: This application is intentionally vulnerable for educational purposes.
 *     DO NOT use in production environments!
 * 
 * This application demonstrates advanced web security vulnerabilities including:
 * - All Version 1 vulnerabilities PLUS:
 * - Command injection (RCE)
 * - File exposure vulnerabilities
 * - IDOR (Insecure Direct Object Reference)
 * - Database direct access
 * - Authentication bypass
 * - Secret flag challenge
 * 
 * Author: Educational Security Project
 * Purpose: Security training and penetration testing practice
 */
package main

// Global application variables
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db        *sql.DB
	store     *sessions.CookieStore
	templates *template.Template
)

// findAvailablePort finds an available port in the range 9001-10000
// ============================================================================
// UTILITY FUNCTIONS
// ============================================================================

/*
 * findAvailablePort - Finds an available port in the range 9001-10000
 * This allows multiple instances to run simultaneously for testing
 */
func findAvailablePort() int {
	for port := 9001; port <= 10000; port++ {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			ln.Close()
			return port
		}
	}
	return 9001 // fallback
}

// User represents a user in the system
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsUser   bool   `json:"is_user"`
	IsAdmin  bool   `json:"is_admin"`
	IsOwner  bool   `json:"is_owner"`
}

// Product represents a product in the catalog
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
}

// CartItem represents an item in the shopping cart
type CartItem struct {
	ProductID int     `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}

// Order represents a customer order
type Order struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	UserEmail string     `json:"user_email"`
	Items     []CartItem `json:"items"`
	Total     float64    `json:"total"`
	CreatedAt time.Time  `json:"created_at"`
}

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./database/ecommerce.db")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize session store
	store = sessions.NewCookieStore([]byte("insecure-secret-key"))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
	}

	// Initialize database tables
	initDatabase()

	// Parse templates with custom functions
	funcMap := template.FuncMap{
		"multiply": func(a, b interface{}) float64 {
			switch a := a.(type) {
			case float64:
				switch b := b.(type) {
				case float64:
					return a * b
				case int:
					return a * float64(b)
				}
			case int:
				switch b := b.(type) {
				case float64:
					return float64(a) * b
				case int:
					return float64(a * b)
				}
			}
			return 0
		},
	}

	// Parse templates individually to avoid conflicts
	templates = template.Must(template.New("").Funcs(funcMap).ParseFiles(
		"templates/base.html",
		"templates/index.html",
		"templates/login.html",
		"templates/register.html",
		"templates/products.html",
		"templates/cart.html",
		"templates/guest_cart.html",
		"templates/checkout.html",
		"templates/orders.html",
		"templates/order_detail.html",
		"templates/admin_dashboard.html",
		"templates/admin_products.html",
		"templates/admin_add_product.html",
		"templates/admin_edit_product.html",
		"templates/admin_orders.html",
		"templates/owner_dashboard.html",
		"templates/payments.html",
	))
}

func initDatabase() {
	// Create users table with PII
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL,
			phone TEXT NOT NULL,
			ssn TEXT NOT NULL,
			date_of_birth TEXT NOT NULL,
			address TEXT NOT NULL,
			city TEXT NOT NULL,
			state TEXT NOT NULL,
			zip_code TEXT NOT NULL,
			credit_card_number TEXT NOT NULL,
			credit_card_expiry TEXT NOT NULL,
			credit_card_cvv TEXT NOT NULL,
			is_user BOOLEAN DEFAULT 1,
			is_admin BOOLEAN DEFAULT 0,
			is_owner BOOLEAN DEFAULT 0
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create products table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT,
			price REAL NOT NULL,
			image_url TEXT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create orders table with PII
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			total REAL NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			shipping_address TEXT NOT NULL,
			billing_address TEXT NOT NULL,
			phone TEXT NOT NULL,
			email TEXT NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users (id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create order_items table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS order_items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			order_id INTEGER NOT NULL,
			product_id INTEGER NOT NULL,
			quantity INTEGER NOT NULL,
			price REAL NOT NULL,
			FOREIGN KEY (order_id) REFERENCES orders (id),
			FOREIGN KEY (product_id) REFERENCES products (id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create payments table with PII
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS payments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			order_id INTEGER NOT NULL,
			card_number TEXT NOT NULL,
			card_holder_name TEXT NOT NULL,
			expiration_month TEXT NOT NULL,
			expiration_year TEXT NOT NULL,
			cvv TEXT NOT NULL,
			billing_address TEXT NOT NULL,
			phone TEXT NOT NULL,
			ssn TEXT NOT NULL,
			driver_license TEXT NOT NULL,
			passport_number TEXT NOT NULL,
			FOREIGN KEY (order_id) REFERENCES orders (id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Seed initial data
	seedInitialData()
}

func seedInitialData() {
	// Check if users already exist
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		log.Println("Users already exist, skipping user creation")
		return
	}

	// Sample users with PII - Using ID scheme: 1000+ for admin, 0 for owner, 1+ for customers
	users := []struct {
		ID               int
		Email            string
		Password         string
		FirstName        string
		LastName         string
		Phone            string
		SSN              string
		DateOfBirth      string
		Address          string
		City             string
		State            string
		ZipCode          string
		CreditCardNumber string
		CreditCardExpiry string
		CreditCardCVV    string
		IsUser           bool
		IsAdmin          bool
		IsOwner          bool
	}{
		{
			ID:               1,
			Email:            "alice@example.com",
			Password:         "insecurepass1",
			FirstName:        "Alice",
			LastName:         "Johnson",
			Phone:            "555-123-4567",
			SSN:              "123-45-6789",
			DateOfBirth:      "1985-03-15",
			Address:          "123 Main Street",
			City:             "New York",
			State:            "NY",
			ZipCode:          "10001",
			CreditCardNumber: "4111111111111111",
			CreditCardExpiry: "12/25",
			CreditCardCVV:    "123",
			IsUser:           true,
			IsAdmin:          false,
			IsOwner:          false,
		},
		{
			ID:               2,
			Email:            "bob@example.com",
			Password:         "insecurepass2",
			FirstName:        "Bob",
			LastName:         "Smith",
			Phone:            "555-987-6543",
			SSN:              "987-65-4321",
			DateOfBirth:      "1990-07-22",
			Address:          "456 Oak Avenue",
			City:             "Los Angeles",
			State:            "CA",
			ZipCode:          "90210",
			CreditCardNumber: "5555555555554444",
			CreditCardExpiry: "08/26",
			CreditCardCVV:    "456",
			IsUser:           true,
			IsAdmin:          false,
			IsOwner:          false,
		},
		{
			ID:               3,
			Email:            "carol@example.com",
			Password:         "insecurepass3",
			FirstName:        "Carol",
			LastName:         "Davis",
			Phone:            "555-333-4444",
			SSN:              "333-44-5555",
			DateOfBirth:      "1988-11-30",
			Address:          "789 Pine Street",
			City:             "Houston",
			State:            "TX",
			ZipCode:          "77001",
			CreditCardNumber: "6011111111111117",
			CreditCardExpiry: "09/27",
			CreditCardCVV:    "789",
			IsUser:           true,
			IsAdmin:          false,
			IsOwner:          false,
		},
		{
			ID:               4,
			Email:            "david@example.com",
			Password:         "insecurepass4",
			FirstName:        "David",
			LastName:         "Wilson",
			Phone:            "555-555-6666",
			SSN:              "555-66-7777",
			DateOfBirth:      "1992-04-18",
			Address:          "321 Elm Court",
			City:             "Phoenix",
			State:            "AZ",
			ZipCode:          "85001",
			CreditCardNumber: "4222222222222222",
			CreditCardExpiry: "03/28",
			CreditCardCVV:    "012",
			IsUser:           true,
			IsAdmin:          false,
			IsOwner:          false,
		},
		{
			ID:               1000,
			Email:            "admin@example.com",
			Password:         "adminpass",
			FirstName:        "Admin",
			LastName:         "User",
			Phone:            "555-111-2222",
			SSN:              "111-22-3333",
			DateOfBirth:      "1980-01-10",
			Address:          "789 Admin Blvd",
			City:             "Chicago",
			State:            "IL",
			ZipCode:          "60601",
			CreditCardNumber: "378282246310005",
			CreditCardExpiry: "06/24",
			CreditCardCVV:    "789",
			IsUser:           true,
			IsAdmin:          true,
			IsOwner:          false,
		},
		{
			ID:               2000,
			Email:            "manager@example.com",
			Password:         "managerpass",
			FirstName:        "Manager",
			LastName:         "Supervisor",
			Phone:            "555-777-8888",
			SSN:              "777-88-9999",
			DateOfBirth:      "1982-06-25",
			Address:          "456 Manager Way",
			City:             "Denver",
			State:            "CO",
			ZipCode:          "80201",
			CreditCardNumber: "5105105105105100",
			CreditCardExpiry: "11/26",
			CreditCardCVV:    "345",
			IsUser:           true,
			IsAdmin:          true,
			IsOwner:          false,
		},
		{
			ID:               0,
			Email:            "owner@example.com",
			Password:         "ownerpass",
			FirstName:        "Owner",
			LastName:         "Manager",
			Phone:            "555-999-8888",
			SSN:              "999-88-7777",
			DateOfBirth:      "1975-12-05",
			Address:          "321 Owner Lane",
			City:             "Miami",
			State:            "FL",
			ZipCode:          "33101",
			CreditCardNumber: "6011111111111117",
			CreditCardExpiry: "03/27",
			CreditCardCVV:    "012",
			IsUser:           true,
			IsAdmin:          true,
			IsOwner:          true,
		},
		{
			Email:            "john.doe@example.com",
			Password:         "password123",
			FirstName:        "John",
			LastName:         "Doe",
			Phone:            "555-444-3333",
			SSN:              "444-33-2222",
			DateOfBirth:      "1992-09-18",
			Address:          "654 Pine Street",
			City:             "Seattle",
			State:            "WA",
			ZipCode:          "98101",
			CreditCardNumber: "4222222222222222",
			CreditCardExpiry: "11/25",
			CreditCardCVV:    "345",
			IsUser:           true,
			IsAdmin:          false,
			IsOwner:          false,
		},
		{
			Email:            "jane.wilson@example.com",
			Password:         "securepass",
			FirstName:        "Jane",
			LastName:         "Wilson",
			Phone:            "555-777-6666",
			SSN:              "777-66-5555",
			DateOfBirth:      "1988-04-30",
			Address:          "987 Elm Court",
			City:             "Austin",
			State:            "TX",
			ZipCode:          "73301",
			CreditCardNumber: "5105105105105100",
			CreditCardExpiry: "09/26",
			CreditCardCVV:    "678",
			IsUser:           true,
			IsAdmin:          false,
			IsOwner:          false,
		},
	}

	for _, user := range users {
		_, err := db.Exec(`
			INSERT INTO users (id, email, password, first_name, last_name, phone, ssn, date_of_birth, address, city, state, zip_code, credit_card_number, credit_card_expiry, credit_card_cvv, is_user, is_admin, is_owner)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`, user.ID, user.Email, user.Password, user.FirstName, user.LastName, user.Phone, user.SSN, user.DateOfBirth, user.Address, user.City, user.State, user.ZipCode, user.CreditCardNumber, user.CreditCardExpiry, user.CreditCardCVV, user.IsUser, user.IsAdmin, user.IsOwner)
		if err != nil {
			log.Printf("Error creating user %s: %v", user.Email, err)
		}
	}

	// Check if products already exist
	var productCount int
	err = db.QueryRow("SELECT COUNT(*) FROM products").Scan(&productCount)
	if err != nil {
		log.Fatal(err)
	}

	if productCount == 0 {
		// Create comprehensive sample products
		products := []Product{
			{Name: "GoLang Mug", Description: "A sturdy mug for your favorite Go developer.", Price: 15.99, ImageURL: "/static/images/go_mug.jpg"},
			{Name: "E-commerce Guidebook", Description: "Learn the basics of e-commerce development.", Price: 29.99, ImageURL: "/static/images/book.jpg"},
			{Name: "Secure Web Dev Poster", Description: "A poster illustrating common web vulnerabilities.", Price: 12.50, ImageURL: "/static/images/poster.jpg"},
			{Name: "Vintage CPU Keychain", Description: "A cool keychain made from recycled tech.", Price: 8.00, ImageURL: "/static/images/cpu_key.jpg"},
			{Name: "Custom Go T-Shirt", Description: "High-quality cotton t-shirt with Go gopher logo.", Price: 24.99, ImageURL: "/static/images/go_tshirt.jpg"},
			{Name: "Data Structures Card Set", Description: "Flashcards for common data structures.", Price: 19.99, ImageURL: "/static/images/flashcards.jpg"},
			{Name: "Debugging Duck", Description: "Your indispensable rubber debugging companion.", Price: 9.50, ImageURL: "/static/images/duck.jpg"},
			{Name: "Code Editor Stickers", Description: "Sticker pack for popular code editors.", Price: 7.00, ImageURL: "/static/images/stickers.jpg"},
			{Name: "Mechanical Keyboard", Description: "A clicky keyboard for serious coding.", Price: 120.00, ImageURL: "/static/images/keyboard.jpg"},
			{Name: "USB-C Hub", Description: "Multi-port hub for modern laptops.", Price: 45.00, ImageURL: "/static/images/usb_hub.jpg"},
			{Name: "Noise Cancelling Headphones", Description: "Immerse yourself in your code.", Price: 150.00, ImageURL: "/static/images/headphones.jpg"},
			{Name: "Smart Home Speaker", Description: "Voice-controlled assistant for your workspace.", Price: 75.00, ImageURL: "/static/images/speaker.jpg"},
		}

		for _, product := range products {
			_, err := db.Exec(`
				INSERT INTO products (name, description, price, image_url)
				VALUES (?, ?, ?, ?)
			`, product.Name, product.Description, product.Price, product.ImageURL)
			if err != nil {
				log.Printf("Error creating product %s: %v", product.Name, err)
			}
		}
	}

	// Check if orders already exist
	var orderCount int
	err = db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&orderCount)
	if err != nil {
		log.Fatal(err)
	}

	if orderCount == 0 {
		// Create sample orders with PII
		orders := []struct {
			UserID          int
			Total           float64
			CreatedAt       string
			ShippingAddress string
			BillingAddress  string
			Phone           string
			Email           string
		}{
			{1, 41.98, "2025-07-28T09:00:00Z", "123 Main Street, New York, NY 10001", "123 Main Street, New York, NY 10001", "555-123-4567", "alice@example.com"},
			{2, 24.99, "2025-07-29T13:45:00Z", "456 Oak Avenue, Los Angeles, CA 90210", "456 Oak Avenue, Los Angeles, CA 90210", "555-987-6543", "bob@example.com"},
			{1, 29.99, "2025-07-30T16:20:00Z", "123 Main Street, New York, NY 10001", "123 Main Street, New York, NY 10001", "555-123-4567", "alice@example.com"},
			{3, 54.97, "2025-07-31T10:00:00Z", "789 Admin Blvd, Chicago, IL 60601", "789 Admin Blvd, Chicago, IL 60601", "555-111-2222", "admin@example.com"},
			{5, 12.50, "2025-08-01T08:00:00Z", "654 Pine Street, Seattle, WA 98101", "654 Pine Street, Seattle, WA 98101", "555-444-3333", "john.doe@example.com"},
			{6, 9.50, "2025-08-01T12:00:00Z", "987 Elm Court, Austin, TX 73301", "987 Elm Court, Austin, TX 73301", "555-777-6666", "jane.wilson@example.com"},
			{1, 203.00, "2025-08-02T09:00:00Z", "123 Main Street, New York, NY 10001", "123 Main Street, New York, NY 10001", "555-123-4567", "alice@example.com"},
			{2, 15.99, "2025-08-02T09:30:00Z", "456 Oak Avenue, Los Angeles, CA 90210", "456 Oak Avenue, Los Angeles, CA 90210", "555-987-6543", "bob@example.com"},
		}

		for _, order := range orders {
			_, err := db.Exec(`
				INSERT INTO orders (user_id, total, created_at, shipping_address, billing_address, phone, email)
				VALUES (?, ?, ?, ?, ?, ?, ?)
			`, order.UserID, order.Total, order.CreatedAt, order.ShippingAddress, order.BillingAddress, order.Phone, order.Email)
			if err != nil {
				log.Printf("Error creating order: %v", err)
			}
		}
	}

	// Check if order_items already exist
	var orderItemCount int
	err = db.QueryRow("SELECT COUNT(*) FROM order_items").Scan(&orderItemCount)
	if err != nil {
		log.Fatal(err)
	}

	if orderItemCount == 0 {
		// Create sample order items
		orderItems := []struct {
			OrderID         int
			ProductID       int
			Quantity        int
			PriceAtPurchase float64
		}{
			{1, 1, 1, 15.99},
			{1, 2, 1, 29.99},
			{2, 5, 1, 24.99},
			{3, 2, 1, 29.99},
			{4, 6, 1, 19.99},
			{4, 8, 5, 7.00},
			{5, 3, 1, 12.50},
			{6, 7, 1, 9.50},
			{7, 11, 1, 150.00},
			{7, 10, 1, 45.00},
			{7, 4, 1, 8.00},
			{8, 1, 1, 15.99},
		}

		for _, item := range orderItems {
			_, err := db.Exec(`
				INSERT INTO order_items (order_id, product_id, quantity, price)
				VALUES (?, ?, ?, ?)
			`, item.OrderID, item.ProductID, item.Quantity, item.PriceAtPurchase)
			if err != nil {
				log.Printf("Error creating order item: %v", err)
			}
		}
	}

	// Check if payments already exist
	var paymentCount int
	err = db.QueryRow("SELECT COUNT(*) FROM payments").Scan(&paymentCount)
	if err != nil {
		log.Fatal(err)
	}

	if paymentCount == 0 {
		// Create sample payments with PII
		payments := []struct {
			OrderID         int
			CardNumber      string
			CardHolderName  string
			ExpirationMonth string
			ExpirationYear  string
			CVV             string
			BillingAddress  string
			Phone           string
			SSN             string
			DriverLicense   string
			PassportNumber  string
		}{
			{1, "4111111111111111", "Alice Johnson", "12", "2025", "123", "123 Main Street, New York, NY 10001", "555-123-4567", "123-45-6789", "DL123456789", "P123456789"},
			{2, "5555555555554444", "Bob Smith", "08", "2026", "456", "456 Oak Avenue, Los Angeles, CA 90210", "555-987-6543", "987-65-4321", "DL987654321", "P987654321"},
			{3, "378282246310005", "Admin User", "06", "2024", "789", "789 Admin Blvd, Chicago, IL 60601", "555-111-2222", "111-22-3333", "DL111222333", "P111222333"},
			{4, "6011111111111117", "Owner Manager", "03", "2027", "012", "321 Owner Lane, Miami, FL 33101", "555-999-8888", "999-88-7777", "DL999888777", "P999888777"},
			{5, "4222222222222222", "John Doe", "11", "2025", "345", "654 Pine Street, Seattle, WA 98101", "555-444-3333", "444-33-2222", "DL444333222", "P444333222"},
			{6, "5105105105105100", "Jane Wilson", "09", "2026", "678", "987 Elm Court, Austin, TX 73301", "555-777-6666", "777-66-5555", "DL777666555", "P777666555"},
			{7, "4111111111111111", "Alice Johnson", "12", "2025", "123", "123 Main Street, New York, NY 10001", "555-123-4567", "123-45-6789", "DL123456789", "P123456789"},
			{8, "5555555555554444", "Bob Smith", "08", "2026", "456", "456 Oak Avenue, Los Angeles, CA 90210", "555-987-6543", "987-65-4321", "DL987654321", "P987654321"},
		}

		for _, payment := range payments {
			_, err := db.Exec(`
				INSERT INTO payments (order_id, card_number, card_holder_name, expiration_month, expiration_year, cvv, billing_address, phone, ssn, driver_license, passport_number)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
			`, payment.OrderID, payment.CardNumber, payment.CardHolderName, payment.ExpirationMonth, payment.ExpirationYear, payment.CVV, payment.BillingAddress, payment.Phone, payment.SSN, payment.DriverLicense, payment.PassportNumber)
			if err != nil {
				log.Printf("Error creating payment: %v", err)
			}
		}
	}
}

// ============================================================================
// MAIN APPLICATION
// ============================================================================

/*
 * main - Application entry point
 * Sets up routes and starts the HTTP server
 * Includes additional vulnerable endpoints for Version 2
 */
func main() {
	r := mux.NewRouter()

	// Static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Public routes
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/login", loginHandler).Methods("GET", "POST")
	r.HandleFunc("/register", registerHandler).Methods("GET", "POST")
	r.HandleFunc("/logout", logoutHandler).Methods("GET")

	// Product routes
	r.HandleFunc("/products", productsHandler).Methods("GET")
	r.HandleFunc("/product/{id}", productDetailHandler).Methods("GET")

	// Cart routes
	r.HandleFunc("/cart", cartHandler).Methods("GET")
	r.HandleFunc("/cart/add", addToCartHandler).Methods("POST")
	r.HandleFunc("/cart/remove", removeFromCartHandler).Methods("POST")

	// Checkout routes
	r.HandleFunc("/checkout", checkoutHandler).Methods("GET", "POST")

	// Order routes
	r.HandleFunc("/orders", ordersHandler).Methods("GET")
	r.HandleFunc("/order/{id}", orderDetailHandler).Methods("GET")

	// Profile route (VULNERABLE: IDOR - no authentication required)
	r.HandleFunc("/profile", profileHandler).Methods("GET")

	// Admin routes (INSECURE: any logged-in user can access)
	r.HandleFunc("/admin", adminMiddleware(adminHandler)).Methods("GET")
	r.HandleFunc("/admin/products", adminMiddleware(adminProductsHandler)).Methods("GET")
	r.HandleFunc("/admin/products/add", adminMiddleware(adminAddProductHandler)).Methods("GET", "POST")
	r.HandleFunc("/admin/products/edit/{id}", adminMiddleware(adminEditProductHandler)).Methods("GET", "POST")
	r.HandleFunc("/admin/products/delete/{id}", adminMiddleware(adminDeleteProductHandler)).Methods("POST")
	r.HandleFunc("/admin/orders", adminMiddleware(adminOrdersHandler)).Methods("GET")

	// Owner routes (INSECURE: any logged-in user can access)
	r.HandleFunc("/owner", ownerMiddleware(ownerHandler)).Methods("GET")
	r.HandleFunc("/payments", ownerMiddleware(paymentsHandler)).Methods("GET")
	r.HandleFunc("/owner/debug", debugHandler).Methods("GET")
	r.HandleFunc("/owner/database", databaseHandler).Methods("GET")

	// INSECURE: Admin routes without any authentication (major vulnerability!)
	r.HandleFunc("/admin/insecure", insecureAdminHandler).Methods("GET")
	r.HandleFunc("/admin/products/insecure", insecureAdminProductsHandler).Methods("GET")
	r.HandleFunc("/payments/insecure", insecurePaymentsHandler).Methods("GET")

	// INSECURE: Owner routes without any authentication (major vulnerability!)
	r.HandleFunc("/owner/payments/insecure", insecureOwnerPaymentsHandler).Methods("GET")

	// Secret flag challenge route
	r.HandleFunc("/secret", secretFlagHandler).Methods("GET")

	// Find available port
	port := findAvailablePort()
	fmt.Printf("🚀 Vibe E-commerce Version 2 starting on http://localhost:%d\n", port)
	fmt.Printf("⚠️  WARNING: This is an intentionally vulnerable application for educational purposes!\n")
	fmt.Printf("🎯 Version 2 includes: Command Injection, File Exposure, IDOR, and Secret Flag Challenge\n")
	fmt.Printf("📚 Test accounts available - see home page for credentials\n")
	fmt.Printf("🚩 Secret Flag Challenge: /secret endpoint\n")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

// Authentication middleware
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		userID, ok := session.Values["user_id"].(int)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		// Store user info in request context
		r.Header.Set("User-ID", strconv.Itoa(userID))
		next.ServeHTTP(w, r)
	}
}

// Admin middleware - INSECURE: allows any logged-in user to access admin functions
func adminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		_, ok := session.Values["user_id"].(int)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// INSECURE: Allow any logged-in user to access admin functions
		// No check for is_admin - this is a major security vulnerability
		next.ServeHTTP(w, r)
	}
}

// Owner middleware - INSECURE: allows any logged-in user to access owner functions
func ownerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		_, ok := session.Values["user_id"].(int)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// INSECURE: Allow any logged-in user to access owner functions
		// No check for is_owner - this is a major security vulnerability
		next.ServeHTTP(w, r)
	}
}

// Handlers
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Check if user is logged in
	session, _ := store.Get(r, "session")
	_, isLoggedIn := session.Values["user_id"].(int)
	userEmail, _ := session.Values["user_email"].(string)

	// Return a simple HTML response for now
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Vibe E-commerce</title>
			<style>
				body { font-family: Arial, sans-serif; margin: 40px; }
				.nav { background: #333; color: white; padding: 10px; }
				.nav a { color: white; text-decoration: none; margin-right: 20px; }
				.container { max-width: 800px; margin: 0 auto; }
				.btn { background: #007bff; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px; }
				.guest-welcome { background: #f8f9fa; padding: 20px; border-radius: 8px; margin: 20px 0; }
				.user-welcome { background: #d4edda; padding: 20px; border-radius: 8px; margin: 20px 0; }
			</style>
		</head>
		<body>
			<div class="nav">
				<a href="/">Vibe Shop</a>
				<a href="/products">Products</a>
				<a href="/cart">Cart</a>
				` + (func() string {
		if isLoggedIn {
			return `<a href="/orders">Orders</a>`
		}
		return `<a href="/login">Login</a>`
	})() + `
				<a href="/register">Register</a>
				<a href="/profile?id=1">Profile</a>
			</div>
			<div class="container">
				<h1>Welcome to Vibe Shop</h1>
				` + (func() string {
		if isLoggedIn {
			return `<div class="user-welcome">
							<h3>Welcome back, ` + userEmail + `!</h3>
							<p>You're logged in and can access all features.</p>
							<p><a href="/products" class="btn">Browse Products</a> | <a href="/cart" class="btn">View Cart</a></p>
						</div>`
		}
		return `<div class="guest-welcome">
						<h3>👋 Welcome Guest!</h3>
						<p>You can browse products without logging in, but you'll need an account to make purchases.</p>
						<p><a href="/products" class="btn">Browse Products</a> | <a href="/login" class="btn">Login</a> | <a href="/register" class="btn">Register</a></p>
					</div>`
	})() + `
				<p><strong>Test Credentials:</strong></p>
				<ul>
					<li>Customer: alice@example.com / insecurepass1</li>
					<li>Admin: admin@example.com / adminpass</li>
					<li>Owner: owner@example.com / ownerpass</li>
				</ul>
			</div>
		</body>
		</html>
	`))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		var user User
		err := db.QueryRow("SELECT id, email, password, is_user, is_admin, is_owner FROM users WHERE email = '"+email+"'").Scan(
			&user.ID, &user.Email, &user.Password, &user.IsUser, &user.IsAdmin, &user.IsOwner)

		if err != nil {
			http.Redirect(w, r, "/login?error=invalid", http.StatusSeeOther)
			return
		}

		if password == user.Password { // Insecure comparison
			session, _ := store.Get(r, "session")
			session.Values["user_id"] = user.ID
			session.Values["user_email"] = user.Email
			session.Values["is_admin"] = user.IsAdmin
			session.Values["is_owner"] = user.IsOwner
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/login?error=invalid", http.StatusSeeOther)
		}
	} else {
		w.Header().Set("Content-Type", "text/html")
		html := `<!DOCTYPE html>
<html>
<head>
    <title>Login - Vibe E-commerce</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .nav { background: #333; color: white; padding: 10px; }
        .nav a { color: white; text-decoration: none; margin-right: 20px; }
        .container { max-width: 400px; margin: 0 auto; }
        .form-group { margin-bottom: 1rem; }
        .form-group label { display: block; margin-bottom: 0.5rem; font-weight: bold; }
        .form-group input { width: 100%; padding: 0.5rem; border: 1px solid #ddd; border-radius: 4px; }
        .btn { background: #007bff; color: white; padding: 10px 20px; border: none; border-radius: 5px; cursor: pointer; width: 100%; }
        .btn:hover { background: #0056b3; }
        .error { color: red; margin-bottom: 1rem; }
        .test-credentials { background: #f8f9fa; padding: 1rem; border-radius: 4px; margin-top: 2rem; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">Vibe Shop</a>
        <a href="/products">Products</a>
        <a href="/cart">Cart</a>
        <a href="/login">Login</a>
        <a href="/register">Register</a>
        <a href="/profile?id=1">Profile</a>
    </div>
    <div class="container">
        <h2>Login</h2>
        <form method="POST" action="/login">
            <div class="form-group">
                <label for="email">Email:</label>
                <input type="email" id="email" name="email">
            </div>
            <div class="form-group">
                <label for="password">Password:</label>
                <input type="password" id="password" name="password">
            </div>
            <button type="submit" class="btn">Login</button>
        </form>
        
        <div class="test-credentials">
            <h3>Test Credentials:</h3>
            <p><strong>Customer:</strong> alice@example.com / insecurepass1</p>
            <p><strong>Admin:</strong> admin@example.com / adminpass</p>
            <p><strong>Owner:</strong> owner@example.com / ownerpass</p>
        </div>
        
        <p style="text-align: center; margin-top: 2rem;">
            <a href="/register">Don't have an account? Register here</a>
        </p>
    </div>
</body>
</html>`
		w.Write([]byte(html))
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		_, err := db.Exec("INSERT INTO users (email, password, is_user) VALUES ('" + email + "', '" + password + "', 1)")

		if err != nil {
			http.Redirect(w, r, "/register?error=failed", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		w.Header().Set("Content-Type", "text/html")
		html := `<!DOCTYPE html>
<html>
<head>
    <title>Register - Vibe E-commerce</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .nav { background: #333; color: white; padding: 10px; }
        .nav a { color: white; text-decoration: none; margin-right: 20px; }
        .container { max-width: 400px; margin: 0 auto; }
        .form-group { margin-bottom: 1rem; }
        .form-group label { display: block; margin-bottom: 0.5rem; font-weight: bold; }
        .form-group input { width: 100%; padding: 0.5rem; border: 1px solid #ddd; border-radius: 4px; }
        .btn { background: #28a745; color: white; padding: 10px 20px; border: none; border-radius: 5px; cursor: pointer; width: 100%; }
        .btn:hover { background: #218838; }
        .error { color: red; margin-bottom: 1rem; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">Vibe Shop</a>
        <a href="/products">Products</a>
        <a href="/cart">Cart</a>
        <a href="/login">Login</a>
        <a href="/register">Register</a>
        <a href="/profile?id=1">Profile</a>
    </div>
    <div class="container">
        <h2>Register</h2>
        <form method="POST" action="/register">
            <div class="form-group">
                <label for="email">Email:</label>
                <input type="email" id="email" name="email">
            </div>
            <div class="form-group">
                <label for="password">Password:</label>
                <input type="password" id="password" name="password">
            </div>
            <button type="submit" class="btn">Register</button>
        </form>
        
        <p style="text-align: center; margin-top: 2rem;">
            <a href="/login">Already have an account? Login here</a>
        </p>
    </div>
</body>
</html>`
		w.Write([]byte(html))
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Values["user_id"] = nil
	session.Values["user_email"] = nil
	session.Values["is_admin"] = nil
	session.Values["is_owner"] = nil
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("search")
	var searchResults string
	var filteredProducts []Product

	// Get all products first
	rows, err := db.Query("SELECT id, name, description, price, image_url FROM products")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var allProducts []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.ImageURL)
		if err != nil {
			continue
		}
		allProducts = append(allProducts, p)
	}

	if searchQuery != "" {
		// Input validation and sanitization
		searchQuery = strings.TrimSpace(searchQuery)

		// Validate search query length
		if len(searchQuery) < 1 || len(searchQuery) > 100 {
			searchResults = "❌ Error: Search query must be between 1 and 100 characters."
		} else {
			// Check for malicious patterns and reject them
			maliciousPatterns := []string{
				";", "|", "&", "`", "$(", "&&", "||", ">", "<", "(", ")", "{", "}", "[", "]",
				"ls", "whoami", "pwd", "id", "cmd", "system", "file", "process", "network",
				"user", "disk", "memory", "environment", "database", "config", "log",
				"exec", "eval", "shell", "bash", "sh", "nc", "netcat", "wget", "curl",
				"rm", "del", "format", "fdisk", "dd", "cat", "head", "tail", "grep",
				"find", "locate", "which", "whereis", "type", "command", "builtin",
			}

			isMalicious := false
			searchQueryLower := strings.ToLower(searchQuery)

			for _, pattern := range maliciousPatterns {
				if strings.Contains(searchQueryLower, pattern) {
					isMalicious = true
					break
				}
			}

			// Additional validation for suspicious command-like patterns
			if strings.Contains(searchQueryLower, "info") &&
				(strings.Contains(searchQueryLower, "system") ||
					strings.Contains(searchQueryLower, "file") ||
					strings.Contains(searchQueryLower, "process") ||
					strings.Contains(searchQueryLower, "network") ||
					strings.Contains(searchQueryLower, "user") ||
					strings.Contains(searchQueryLower, "disk") ||
					strings.Contains(searchQueryLower, "memory") ||
					strings.Contains(searchQueryLower, "environment") ||
					strings.Contains(searchQueryLower, "database") ||
					strings.Contains(searchQueryLower, "config") ||
					strings.Contains(searchQueryLower, "log")) {
				isMalicious = true
			}

			if isMalicious {
				searchResults = "❌ Error: Invalid search query detected. Please use only alphanumeric characters and common words for product searches."
			} else {
				// Perform secure product search
				for _, product := range allProducts {
					if strings.Contains(strings.ToLower(product.Name), strings.ToLower(searchQuery)) ||
						strings.Contains(strings.ToLower(product.Description), strings.ToLower(searchQuery)) {
						filteredProducts = append(filteredProducts, product)
					}
				}
				if len(filteredProducts) == 0 {
					searchResults = "No products found matching your search."
				}
			}
		}
	} else {
		filteredProducts = allProducts
	}

	// Use filtered products or all products
	products := filteredProducts
	if len(products) == 0 {
		products = allProducts
	}

	// Check if user is logged in
	session, _ := store.Get(r, "session")
	_, isLoggedIn := session.Values["user_id"].(int)

	// Create simple HTML response for products
	w.Header().Set("Content-Type", "text/html")

	html := `<!DOCTYPE html>
<html>
<head>
    <title>Products - Vibe E-commerce</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .nav { background: #333; color: white; padding: 10px; }
        .nav a { color: white; text-decoration: none; margin-right: 20px; }
        .container { max-width: 1200px; margin: 0 auto; }
        .btn { background: #007bff; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px; }
        .guest-notice { background-color: #fff3cd; border: 1px solid #ffeaa7; color: #856404; padding: 1rem; border-radius: 4px; margin-bottom: 2rem; }
        .products-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(300px, 1fr)); gap: 2rem; margin-top: 2rem; }
        .product-card { background: white; border-radius: 8px; overflow: hidden; box-shadow: 0 2px 4px rgba(0,0,0,0.1); transition: transform 0.3s; }
        .product-card:hover { transform: translateY(-5px); }
        .product-image-placeholder { width: 100%; height: 200px; background-color: #ecf0f1; display: flex; align-items: center; justify-content: center; color: #7f8c8d; }
        .product-info { padding: 1.5rem; }
        .price { font-size: 1.2rem; font-weight: bold; color: #27ae60; margin: 1rem 0; }
        .guest-action { margin-top: 1rem; }
        .search-form { background: #f8f9fa; padding: 1rem; border-radius: 4px; margin-bottom: 2rem; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">Vibe Shop</a>
        <a href="/products">Products</a>
        <a href="/cart">Cart</a>
        <a href="/login">Login</a>
        <a href="/register">Register</a>
        <a href="/profile?id=1">Profile</a>
    </div>
    <div class="container">
        <h2>Our Products</h2>
        
        <div class="search-form">
            <h3>🔍 Search Products</h3>
            <form method="GET" action="/products">
                <input type="text" name="search" placeholder="Enter search term..." value="` + searchQuery + `" style="width: 300px; padding: 0.5rem; margin-right: 1rem;">
                <button type="submit" class="btn">Search</button>
            </form>
            <p><em>Secure product search with input validation and sanitization</em></p>
        </div>`

	if searchResults != "" && searchResults != "No products found matching your search." {
		if strings.Contains(searchResults, "❌ Error:") {
			html += `<div class="search-results">
                <h4>🚫 Validation Error:</h4>
                <p style="color: #dc3545; font-weight: bold; background: #f8d7da; padding: 1rem; border-radius: 4px; border-left: 4px solid #dc3545;">` + searchResults + `</p>
            </div>`
		} else {
			html += `<div class="search-results">
                <h4>🔍 Search Results:</h4>
                <p style="color: #666; font-style: italic;">` + searchResults + `</p>
            </div>`
		}
	}

	if !isLoggedIn {
		html += `<div class="guest-notice">
            <p><strong>👋 Welcome Guest!</strong> You can browse products and add items to your cart. Guest checkout is available!</p>
        </div>`
	}

	html += `<div class="products-grid">`

	for _, product := range products {
		html += `<div class="product-card">
            <div class="product-image-placeholder">No Image</div>
            <div class="product-info">
                <h3>` + product.Name + `</h3>
                <p>` + product.Description + `</p>
                <p class="price">$` + fmt.Sprintf("%.2f", product.Price) + `</p>`

		html += `<form method="POST" action="/cart/add">
                <input type="hidden" name="product_id" value="` + strconv.Itoa(product.ID) + `">
                <input type="number" name="quantity" value="1" min="1" style="width: 60px; padding: 0.5rem; margin-right: 1rem; border: 1px solid #ddd; border-radius: 4px;">
                <button type="submit" class="btn">Add to Cart</button>
            </form>`

		html += `</div>
        </div>`
	}

	html += `</div>
    </div>
</body>
</html>`

	w.Write([]byte(html))
}

func productDetailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var product Product
	err := db.QueryRow("SELECT id, name, description, price, image_url FROM products WHERE id = ?", id).Scan(
		&product.ID, &product.Name, &product.Description, &product.Price, &product.ImageURL)

	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	templates.ExecuteTemplate(w, "product_detail.html", product)
}

func cartHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, isLoggedIn := session.Values["user_id"].(int)

	// Get cart items from session
	cartJSON := session.Values["cart"]
	var cart []CartItem
	if cartJSON != nil {
		json.Unmarshal([]byte(cartJSON.(string)), &cart)
	}

	w.Header().Set("Content-Type", "text/html")

	if !isLoggedIn {
		// Show guest cart with actual items
		html := `<!DOCTYPE html>
<html>
<head>
    <title>Cart - Vibe E-commerce</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .nav { background: #333; color: white; padding: 10px; }
        .nav a { color: white; text-decoration: none; margin-right: 20px; }
        .container { max-width: 800px; margin: 0 auto; }
        .btn { background: #007bff; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px; }
        .cart-item { background: white; padding: 1rem; margin: 1rem 0; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        .cart-total { font-size: 1.2rem; font-weight: bold; color: #28a745; margin: 1rem 0; }
        .guest-notice { background: #fff3cd; border: 1px solid #ffeaa7; color: #856404; padding: 1rem; border-radius: 4px; margin-bottom: 2rem; }
        .empty-cart { text-align: center; padding: 2rem; color: #6c757d; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">Vibe Shop</a>
        <a href="/products">Products</a>
        <a href="/cart">Cart</a>
        <a href="/login">Login</a>
        <a href="/register">Register</a>
        <a href="/profile?id=1">Profile</a>
    </div>
    <div class="container">
        <h2>Shopping Cart</h2>`

		if len(cart) == 0 {
			html += `<div class="empty-cart">
                <h3>Your cart is empty</h3>
                <p>Add some products to get started!</p>
                <a href="/products" class="btn">Browse Products</a>
            </div>`
		} else {
			var total float64
			for _, item := range cart {
				total += item.Price * float64(item.Quantity)
				html += fmt.Sprintf(`<div class="cart-item">
                    <h3>%s</h3>
                    <p>Price: $%.2f x %d = $%.2f</p>
                    <form method="POST" action="/cart/remove" style="display: inline;">
                        <input type="hidden" name="index" value="%d">
                        <button type="submit" style="background: #dc3545; color: white; border: none; padding: 5px 10px; border-radius: 3px; cursor: pointer;">Remove</button>
                    </form>
                </div>`, item.Name, item.Price, item.Quantity, item.Price*float64(item.Quantity), 0) // Note: index calculation would need to be more complex
			}

			html += fmt.Sprintf(`<div class="cart-total">Total: $%.2f</div>
            <div class="guest-notice">
                <p><strong>Guest Checkout Available!</strong> You can complete your purchase as a guest.</p>
                <a href="/checkout" class="btn">Proceed to Checkout</a>
            </div>`)
		}

		html += `<div style="margin-top: 2rem;">
            <a href="/products" class="btn" style="background: #95a5a6;">Continue Shopping</a>
        </div>
    </div>
</body>
</html>`

		w.Write([]byte(html))
	} else {
		// For logged-in users, show their cart
		html := `<!DOCTYPE html>
<html>
<head>
    <title>Cart - Vibe E-commerce</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .nav { background: #333; color: white; padding: 10px; }
        .nav a { color: white; text-decoration: none; margin-right: 20px; }
        .container { max-width: 800px; margin: 0 auto; }
        .btn { background: #007bff; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px; }
        .cart-item { background: white; padding: 1rem; margin: 1rem 0; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        .cart-total { font-size: 1.2rem; font-weight: bold; color: #28a745; margin: 1rem 0; }
        .empty-cart { text-align: center; padding: 2rem; color: #6c757d; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">Vibe Shop</a>
        <a href="/products">Products</a>
        <a href="/cart">Cart</a>
        <a href="/orders">Orders</a>
        <a href="/profile?id=1">Profile</a>
        <a href="/logout">Logout</a>
    </div>
    <div class="container">
        <h2>Your Shopping Cart</h2>`

		if len(cart) == 0 {
			html += `<div class="empty-cart">
                <h3>Your cart is empty</h3>
                <p>Add some products to get started!</p>
                <a href="/products" class="btn">Browse Products</a>
            </div>`
		} else {
			var total float64
			for _, item := range cart {
				total += item.Price * float64(item.Quantity)
				html += fmt.Sprintf(`<div class="cart-item">
                    <h3>%s</h3>
                    <p>Price: $%.2f x %d = $%.2f</p>
                    <form method="POST" action="/cart/remove" style="display: inline;">
                        <input type="hidden" name="index" value="%d">
                        <button type="submit" style="background: #dc3545; color: white; border: none; padding: 5px 10px; border-radius: 3px; cursor: pointer;">Remove</button>
                    </form>
                </div>`, item.Name, item.Price, item.Quantity, item.Price*float64(item.Quantity), 0) // Note: index calculation would need to be more complex
			}

			html += fmt.Sprintf(`<div class="cart-total">Total: $%.2f</div>
            <a href="/checkout" class="btn">Proceed to Checkout</a>`)
		}

		html += `</div>
</body>
</html>`

		w.Write([]byte(html))
	}
}

func addToCartHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	productID, _ := strconv.Atoi(r.FormValue("product_id"))
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))

	if quantity <= 0 {
		quantity = 1
	}

	var product Product
	err := db.QueryRow("SELECT id, name, price FROM products WHERE id = "+strconv.Itoa(productID)).Scan(
		&product.ID, &product.Name, &product.Price)

	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Get current cart from session
	cartJSON := session.Values["cart"]
	var cart []CartItem
	if cartJSON != nil {
		json.Unmarshal([]byte(cartJSON.(string)), &cart)
	}

	// Check if product already in cart
	found := false
	for i, item := range cart {
		if item.ProductID == productID {
			cart[i].Quantity += quantity
			found = true
			break
		}
	}

	if !found {
		cart = append(cart, CartItem{
			ProductID: productID,
			Name:      product.Name,
			Price:     product.Price,
			Quantity:  quantity,
		})
	}

	// Store cart as JSON string
	cartBytes, _ := json.Marshal(cart)
	session.Values["cart"] = string(cartBytes)
	session.Save(r, w)

	// Debug: log the cart after adding
	log.Printf("Added product %d to cart. Cart now has %d items", productID, len(cart))

	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func removeFromCartHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	index, _ := strconv.Atoi(r.FormValue("index"))

	cartItems := session.Values["cart"]
	if cartItems != nil {
		cart := cartItems.([]CartItem)
		if index >= 0 && index < len(cart) {
			cart = append(cart[:index], cart[index+1:]...)
			session.Values["cart"] = cart
			session.Save(r, w)
		}
	}

	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, isLoggedIn := session.Values["user_id"].(int)

	if r.Method == "POST" {
		// Process checkout
		cartJSON := session.Values["cart"]
		if cartJSON == nil {
			http.Redirect(w, r, "/cart", http.StatusSeeOther)
			return
		}

		var cart []CartItem
		json.Unmarshal([]byte(cartJSON.(string)), &cart)

		if len(cart) == 0 {
			http.Redirect(w, r, "/cart", http.StatusSeeOther)
			return
		}

		// Calculate total
		var total float64
		for _, item := range cart {
			total += item.Price * float64(item.Quantity)
		}

		// For guest checkout, use a default user ID (1) or create a guest user
		userID := 1 // Default to first user for guest orders
		if isLoggedIn {
			userID, _ = session.Values["user_id"].(int)
		}

		// Create order
		result, err := db.Exec("INSERT INTO orders (user_id, total) VALUES (" + strconv.Itoa(userID) + ", " + fmt.Sprintf("%f", total) + ")")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		orderID, _ := result.LastInsertId()

		// Add order items
		for _, item := range cart {
			_, err := db.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES (" + strconv.FormatInt(orderID, 10) + ", " + strconv.Itoa(item.ProductID) + ", " + strconv.Itoa(item.Quantity) + ", " + fmt.Sprintf("%f", item.Price) + ")")
			if err != nil {
				log.Printf("Error adding order item: %v", err)
			}
		}

		// Store payment info for educational purposes
		cardNumber := r.FormValue("card_number")
		cardHolder := r.FormValue("card_holder")
		expMonth := r.FormValue("exp_month")
		expYear := r.FormValue("exp_year")
		cvv := r.FormValue("cvv")

		_, err = db.Exec(`
			INSERT INTO payments (order_id, card_number, card_holder_name, expiration_month, expiration_year, cvv)
			VALUES (` + strconv.FormatInt(orderID, 10) + `, '` + cardNumber + `', '` + cardHolder + `', '` + expMonth + `', '` + expYear + `', '` + cvv + `')
		`)

		// Clear cart
		session.Values["cart"] = ""
		session.Save(r, w)

		http.Redirect(w, r, fmt.Sprintf("/order/%d", orderID), http.StatusSeeOther)
	} else {
		cartJSON := session.Values["cart"]
		if cartJSON == nil {
			cartJSON = ""
		}

		var cart []CartItem
		if cartJSON != "" {
			json.Unmarshal([]byte(cartJSON.(string)), &cart)
		}

		var total float64
		for _, item := range cart {
			total += item.Price * float64(item.Quantity)
		}

		w.Header().Set("Content-Type", "text/html")
		html := `<!DOCTYPE html>
<html>
<head>
    <title>Checkout - Vibe E-commerce</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .nav { background: #333; color: white; padding: 10px; }
        .nav a { color: white; text-decoration: none; margin-right: 20px; }
        .container { max-width: 600px; margin: 0 auto; }
        .form-group { margin-bottom: 1rem; }
        .form-group label { display: block; margin-bottom: 0.5rem; font-weight: bold; }
        .form-group input { width: 100%; padding: 0.5rem; border: 1px solid #ddd; border-radius: 4px; }
        .btn { background: #28a745; color: white; padding: 10px 20px; border: none; border-radius: 5px; cursor: pointer; width: 100%; }
        .btn:hover { background: #218838; }
        .cart-summary { background: #f8f9fa; padding: 1rem; border-radius: 4px; margin: 1rem 0; }
        .total { font-size: 1.2rem; font-weight: bold; color: #28a745; }
        .guest-notice { background: #fff3cd; border: 1px solid #ffeaa7; color: #856404; padding: 1rem; border-radius: 4px; margin: 1rem 0; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">Vibe Shop</a>
        <a href="/products">Products</a>
        <a href="/cart">Cart</a>
        <a href="/login">Login</a>
        <a href="/register">Register</a>
        <a href="/profile?id=1">Profile</a>
    </div>
    <div class="container">
        <h2>Checkout</h2>`

		if !isLoggedIn {
			html += `<div class="guest-notice">
                <p><strong>Guest Checkout</strong> - You're checking out as a guest. Your order will be processed normally.</p>
            </div>`
		}

		html += `<div class="cart-summary">
            <h3>Order Summary</h3>`

		for _, item := range cart {
			html += fmt.Sprintf(`<div style="margin: 0.5rem 0;">
                <strong>%s</strong> - $%.2f x %d = $%.2f
            </div>`, item.Name, item.Price, item.Quantity, item.Price*float64(item.Quantity))
		}

		html += fmt.Sprintf(`<div class="total">Total: $%.2f</div>
        </div>
        
        <form method="POST" action="/checkout">
            <div class="form-group">
                <label for="card_number">Card Number:</label>
                <input type="text" id="card_number" name="card_number" placeholder="1234 5678 9012 3456">
            </div>
            <div class="form-group">
                <label for="card_holder">Card Holder Name:</label>
                <input type="text" id="card_holder" name="card_holder">
            </div>
            <div style="display: flex; gap: 1rem;">
                <div class="form-group" style="flex: 1;">
                    <label for="exp_month">Expiry Month:</label>
                    <input type="text" id="exp_month" name="exp_month" placeholder="MM">
                </div>
                <div class="form-group" style="flex: 1;">
                    <label for="exp_year">Expiry Year:</label>
                    <input type="text" id="exp_year" name="exp_year" placeholder="YYYY">
                </div>
                <div class="form-group" style="flex: 1;">
                    <label for="cvv">CVV:</label>
                    <input type="text" id="cvv" name="cvv" placeholder="123">
                </div>
            </div>
            <button type="submit" class="btn">Complete Purchase</button>
        </form>
    </div>
</body>
</html>`, total)

		w.Write([]byte(html))
	}
}

func ordersHandler(w http.ResponseWriter, r *http.Request) {
	// INSECURE: No authentication required - anyone can view orders
	rows, err := db.Query(`
		SELECT o.id, o.total, o.created_at, o.shipping_address, o.billing_address, o.phone, o.email as order_email, 
		       u.email as customer_email, u.first_name, u.last_name, u.ssn, u.date_of_birth, u.address, u.city, u.state, u.zip_code,
		       u.credit_card_number, u.credit_card_expiry, u.credit_card_cvv
		FROM orders o 
		JOIN users u ON o.user_id = u.id 
		ORDER BY o.created_at DESC
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type OrderWithPII struct {
		ID               int
		Total            float64
		CreatedAt        time.Time
		ShippingAddress  string
		BillingAddress   string
		Phone            string
		OrderEmail       string
		CustomerEmail    string
		FirstName        string
		LastName         string
		SSN              string
		DateOfBirth      string
		Address          string
		City             string
		State            string
		ZipCode          string
		CreditCardNumber string
		CreditCardExpiry string
		CreditCardCVV    string
	}

	var orders []OrderWithPII
	for rows.Next() {
		var order OrderWithPII
		err := rows.Scan(&order.ID, &order.Total, &order.CreatedAt, &order.ShippingAddress, &order.BillingAddress, &order.Phone, &order.OrderEmail, &order.CustomerEmail, &order.FirstName, &order.LastName, &order.SSN, &order.DateOfBirth, &order.Address, &order.City, &order.State, &order.ZipCode, &order.CreditCardNumber, &order.CreditCardExpiry, &order.CreditCardCVV)
		if err != nil {
			continue
		}
		orders = append(orders, order)
	}

	w.Header().Set("Content-Type", "text/html")
	html := `<!DOCTYPE html>
<html>
<head>
    <title>Orders - Vibe E-commerce</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .nav { background: #333; color: white; padding: 10px; }
        .nav a { color: white; text-decoration: none; margin-right: 20px; }
        .container { max-width: 1200px; margin: 0 auto; }
        .order-item { background: white; padding: 1rem; margin: 1rem 0; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        .order-id { font-weight: bold; color: #007bff; }
        .order-total { font-size: 1.2rem; font-weight: bold; color: #28a745; }
        .order-date { color: #6c757d; font-size: 0.9rem; }
        .no-orders { text-align: center; padding: 2rem; color: #6c757d; }
        .pii-data { background: #f8f9fa; padding: 10px; margin: 10px 0; border-radius: 5px; }
        .pii-label { font-weight: bold; color: #dc3545; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">Vibe Shop</a>
        <a href="/products">Products</a>
        <a href="/cart">Cart</a>
        <a href="/orders">Orders</a>
        <a href="/profile?id=1">Profile</a>
        <a href="/logout">Logout</a>
    </div>
    <div class="container">
        <h2>All Orders (INSECURE: Shows PII Data)</h2>`

	if len(orders) == 0 {
		html += `<div class="no-orders">
            <h3>No orders yet</h3>
            <p>Start shopping to see your orders here!</p>
            <a href="/products" style="background: #007bff; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px;">Browse Products</a>
        </div>`
	} else {
		for _, order := range orders {
			html += fmt.Sprintf(`<div class="order-item">
                <div class="order-id">Order #%d</div>
                <div class="order-total">$%.2f</div>
                <div class="order-date">%s</div>
                <div class="pii-data">
                    <div><span class="pii-label">Customer Name:</span> %s %s</div>
                    <div><span class="pii-label">Customer Email:</span> %s</div>
                    <div><span class="pii-label">Order Email:</span> %s</div>
                    <div><span class="pii-label">Phone:</span> %s</div>
                    <div><span class="pii-label">SSN:</span> %s</div>
                    <div><span class="pii-label">Date of Birth:</span> %s</div>
                    <div><span class="pii-label">Full Address:</span> %s, %s, %s %s</div>
                    <div><span class="pii-label">Shipping Address:</span> %s</div>
                    <div><span class="pii-label">Billing Address:</span> %s</div>
                    <div><span class="pii-label">Credit Card Number:</span> %s</div>
                    <div><span class="pii-label">Credit Card Expiry:</span> %s</div>
                    <div><span class="pii-label">Credit Card CVV:</span> %s</div>
                </div>
                <a href="/order/%d" style="color: #007bff; text-decoration: none;">View Details</a>
            </div>`, order.ID, order.Total, order.CreatedAt.Format("January 2, 2006 at 3:04 PM"), order.FirstName, order.LastName, order.CustomerEmail, order.OrderEmail, order.Phone, order.SSN, order.DateOfBirth, order.Address, order.City, order.State, order.ZipCode, order.ShippingAddress, order.BillingAddress, order.CreditCardNumber, order.CreditCardExpiry, order.CreditCardCVV, order.ID)
		}
	}

	html += `</div>
</body>
</html>`

	w.Write([]byte(html))
}

func orderDetailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	// INSECURE: No authentication required - anyone can view any order
	// Get order details with PII
	type OrderDetailWithPII struct {
		ID               int
		UserID           int
		Total            float64
		CreatedAt        time.Time
		UserEmail        string
		ShippingAddress  string
		BillingAddress   string
		Phone            string
		OrderEmail       string
		FirstName        string
		LastName         string
		SSN              string
		DateOfBirth      string
		Address          string
		City             string
		State            string
		ZipCode          string
		CreditCardNumber string
		CreditCardExpiry string
		CreditCardCVV    string
		Items            []CartItem
	}

	var order OrderDetailWithPII
	err := db.QueryRow(`
		SELECT o.id, o.user_id, o.total, o.created_at, u.email, o.shipping_address, o.billing_address, o.phone, o.email as order_email,
		       u.first_name, u.last_name, u.ssn, u.date_of_birth, u.address, u.city, u.state, u.zip_code,
		       u.credit_card_number, u.credit_card_expiry, u.credit_card_cvv
		FROM orders o 
		JOIN users u ON o.user_id = u.id 
		WHERE o.id = `+orderID+`
	`).Scan(&order.ID, &order.UserID, &order.Total, &order.CreatedAt, &order.UserEmail, &order.ShippingAddress, &order.BillingAddress, &order.Phone, &order.OrderEmail, &order.FirstName, &order.LastName, &order.SSN, &order.DateOfBirth, &order.Address, &order.City, &order.State, &order.ZipCode, &order.CreditCardNumber, &order.CreditCardExpiry, &order.CreditCardCVV)

	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	// Get order items
	rows, err := db.Query(`
		SELECT oi.product_id, p.name, oi.quantity, oi.price 
		FROM order_items oi 
		JOIN products p ON oi.product_id = p.id 
		WHERE oi.order_id = ` + orderID + `
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var item CartItem
		err := rows.Scan(&item.ProductID, &item.Name, &item.Quantity, &item.Price)
		if err != nil {
			continue
		}
		order.Items = append(order.Items, item)
	}

	// Generate HTML response with PII data
	w.Header().Set("Content-Type", "text/html")
	html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>Order Details - Vibe E-commerce</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .nav { background: #333; color: white; padding: 10px; }
        .nav a { color: white; text-decoration: none; margin-right: 20px; }
        .container { max-width: 1200px; margin: 0 auto; }
        .order-details { background: white; padding: 2rem; margin: 1rem 0; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        .order-id { font-size: 1.5rem; font-weight: bold; color: #007bff; }
        .order-total { font-size: 1.2rem; font-weight: bold; color: #28a745; }
        .order-date { color: #6c757d; font-size: 0.9rem; }
        .pii-data { background: #f8f9fa; padding: 15px; margin: 15px 0; border-radius: 5px; border-left: 4px solid #dc3545; }
        .pii-label { font-weight: bold; color: #dc3545; }
        .item { padding: 10px; border-bottom: 1px solid #eee; }
        .item:last-child { border-bottom: none; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">Vibe Shop</a>
        <a href="/products">Products</a>
        <a href="/cart">Cart</a>
        <a href="/orders">Orders</a>
        <a href="/logout">Logout</a>
    </div>
    <div class="container">
        <div class="order-details">
            <div class="order-id">Order #%d</div>
            <div class="order-total">$%.2f</div>
            <div class="order-date">%s</div>
            
            <div class="pii-data">
                <h3>PII Data (INSECURE: Exposed)</h3>
                <div><span class="pii-label">Customer Name:</span> %s %s</div>
                <div><span class="pii-label">Customer Email:</span> %s</div>
                <div><span class="pii-label">Order Email:</span> %s</div>
                <div><span class="pii-label">Phone:</span> %s</div>
                <div><span class="pii-label">SSN:</span> %s</div>
                <div><span class="pii-label">Date of Birth:</span> %s</div>
                <div><span class="pii-label">Full Address:</span> %s, %s, %s %s</div>
                <div><span class="pii-label">Shipping Address:</span> %s</div>
                <div><span class="pii-label">Billing Address:</span> %s</div>
                <div><span class="pii-label">Credit Card Number:</span> %s</div>
                <div><span class="pii-label">Credit Card Expiry:</span> %s</div>
                <div><span class="pii-label">Credit Card CVV:</span> %s</div>
            </div>
            
            <h3>Order Items:</h3>`, order.ID, order.Total, order.CreatedAt.Format("January 2, 2006 at 3:04 PM"), order.FirstName, order.LastName, order.UserEmail, order.OrderEmail, order.Phone, order.SSN, order.DateOfBirth, order.Address, order.City, order.State, order.ZipCode, order.ShippingAddress, order.BillingAddress, order.CreditCardNumber, order.CreditCardExpiry, order.CreditCardCVV)

	for _, item := range order.Items {
		html += fmt.Sprintf(`<div class="item">
                <strong>%s</strong> - Quantity: %d - Price: $%.2f - Total: $%.2f
            </div>`, item.Name, item.Quantity, item.Price, float64(item.Quantity)*item.Price)
	}

	html += `</div>
    </div>
</body>
</html>`

	w.Write([]byte(html))
}

// Admin handlers
func adminHandler(w http.ResponseWriter, r *http.Request) {
	adminMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		html := `<!DOCTYPE html>
<html>
<head>
    <title>Admin Dashboard - Vibe E-commerce</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .nav { background: #333; color: white; padding: 10px; }
        .nav a { color: white; text-decoration: none; margin-right: 20px; }
        .container { max-width: 800px; margin: 0 auto; }
        .admin-section { background: white; padding: 1.5rem; margin: 1rem 0; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        .btn { background: #007bff; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px; display: inline-block; margin: 0.5rem; }
        .btn:hover { background: #0056b3; }
        .btn-danger { background: #dc3545; }
        .btn-danger:hover { background: #c82333; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">Vibe Shop</a>
        <a href="/admin">Admin Dashboard</a>
        <a href="/admin/products">Manage Products</a>
        <a href="/admin/orders">View Orders</a>
        <a href="/logout">Logout</a>
    </div>
    <div class="container">
        <h2>Admin Dashboard</h2>
        
        <div class="admin-section">
            <h3>Product Management</h3>
            <p>Add, edit, and delete products from the catalog.</p>
            <a href="/admin/products" class="btn">Manage Products</a>
            <a href="/admin/products/add" class="btn">Add New Product</a>
        </div>
        
        <div class="admin-section">
            <h3>Order Management</h3>
            <p>View and manage customer orders.</p>
            <a href="/admin/orders" class="btn">View All Orders</a>
        </div>
        
        <div class="admin-section">
            <h3>Quick Actions</h3>
            <a href="/products" class="btn">View Products</a>
            <a href="/" class="btn">Go to Home</a>
        </div>
    </div>
</body>
</html>`
		w.Write([]byte(html))
	})(w, r)
}

func adminProductsHandler(w http.ResponseWriter, r *http.Request) {
	adminMiddleware(func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, description, price, image_url FROM products ORDER BY id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var products []Product
		for rows.Next() {
			var p Product
			err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.ImageURL)
			if err != nil {
				continue
			}
			products = append(products, p)
		}

		templates.ExecuteTemplate(w, "admin_products.html", map[string]interface{}{
			"Products": products,
		})
	})(w, r)
}

func adminAddProductHandler(w http.ResponseWriter, r *http.Request) {
	adminMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			name := r.FormValue("name")
			description := r.FormValue("description")
			price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
			imageURL := r.FormValue("image_url")

			_, err := db.Exec("INSERT INTO products (name, description, price, image_url) VALUES ('" + name + "', '" + description + "', " + fmt.Sprintf("%f", price) + ", '" + imageURL + "')")

			if err != nil {
				templates.ExecuteTemplate(w, "admin_add_product.html", map[string]interface{}{
					"Error": "Failed to add product",
				})
				return
			}

			http.Redirect(w, r, "/admin/products", http.StatusSeeOther)
		} else {
			templates.ExecuteTemplate(w, "admin_add_product.html", nil)
		}
	})(w, r)
}

func adminEditProductHandler(w http.ResponseWriter, r *http.Request) {
	adminMiddleware(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if r.Method == "POST" {
			name := r.FormValue("name")
			description := r.FormValue("description")
			price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
			imageURL := r.FormValue("image_url")

			_, err := db.Exec("UPDATE products SET name = '" + name + "', description = '" + description + "', price = " + fmt.Sprintf("%f", price) + ", image_url = '" + imageURL + "' WHERE id = " + id)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/admin/products", http.StatusSeeOther)
		} else {
			var product Product
			err := db.QueryRow("SELECT id, name, description, price, image_url FROM products WHERE id = "+id).Scan(
				&product.ID, &product.Name, &product.Description, &product.Price, &product.ImageURL)

			if err != nil {
				http.Error(w, "Product not found", http.StatusNotFound)
				return
			}

			templates.ExecuteTemplate(w, "admin_edit_product.html", product)
		}
	})(w, r)
}

func adminDeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	adminMiddleware(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		_, err := db.Exec("DELETE FROM products WHERE id = " + id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/products", http.StatusSeeOther)
	})(w, r)
}

func adminOrdersHandler(w http.ResponseWriter, r *http.Request) {
	adminMiddleware(func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
			SELECT o.id, o.total, o.created_at, u.email 
			FROM orders o 
			JOIN users u ON o.user_id = u.id 
			ORDER BY o.created_at DESC
		`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var orders []Order
		for rows.Next() {
			var order Order
			err := rows.Scan(&order.ID, &order.Total, &order.CreatedAt, &order.UserEmail)
			if err != nil {
				continue
			}
			orders = append(orders, order)
		}

		templates.ExecuteTemplate(w, "admin_orders.html", map[string]interface{}{
			"Orders": orders,
		})
	})(w, r)
}

// Owner handlers
func ownerHandler(w http.ResponseWriter, r *http.Request) {
	ownerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		// Get system statistics
		var stats struct {
			TotalUsers    int
			TotalProducts int
			TotalOrders   int
			TotalRevenue  float64
		}

		db.QueryRow("SELECT COUNT(*) FROM users").Scan(&stats.TotalUsers)
		db.QueryRow("SELECT COUNT(*) FROM products").Scan(&stats.TotalProducts)
		db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&stats.TotalOrders)
		db.QueryRow("SELECT COALESCE(SUM(total), 0) FROM orders").Scan(&stats.TotalRevenue)

		templates.ExecuteTemplate(w, "owner_dashboard.html", stats)
	})(w, r)
}

func debugHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	cartItems := session.Values["cart"]

	// Test setting a simple string value
	session.Values["simple_test"] = "hello world"
	session.Save(r, w)

	// VULNERABLE: Command injection through config parameter
	configCmd := r.URL.Query().Get("config")
	var configOutput string

	if configCmd != "" {
		// CRITICAL VULNERABILITY: Direct command execution from debug interface
		// ⚠️ CRITICAL VULNERABILITY: Command injection through exec.Command
		// This allows arbitrary command execution on the server
		cmd := exec.Command("sh", "-c", "echo 'Server Config: "+configCmd+"' cmd := exec.Command("sh", "-c", "echo 'Server Config: "+configCmd+"' && cat /etc/passwd")cmd := exec.Command("sh", "-c", "echo 'Server Config: "+configCmd+"' && cat /etc/passwd") cat /etc/passwd")
		output, err := cmd.CombinedOutput()
		if err != nil {
			configOutput = "Error executing config command: " + err.Error()
		} else {
			configOutput = string(output)
		}
	}

	w.Header().Set("Content-Type", "text/html")
	html := `<!DOCTYPE html>
<html>
<head>
    <title>Debug Mode - Vibe E-commerce</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; background: #f0f0f0; }
        .debug-container { max-width: 1200px; margin: 0 auto; background: white; padding: 2rem; border-radius: 8px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        .debug-header { background: #dc3545; color: white; padding: 1rem; border-radius: 4px; margin-bottom: 2rem; }
        .debug-section { background: #f8f9fa; padding: 1rem; border-radius: 4px; margin-bottom: 1rem; }
        .config-form { background: #fff3cd; border: 1px solid #ffeaa7; padding: 1rem; border-radius: 4px; margin-bottom: 2rem; }
        .config-output { background: #dc3545; color: white; padding: 1rem; border-radius: 4px; font-family: monospace; white-space: pre-wrap; margin-top: 1rem; }
        .vulnerability-warning { background: #dc3545; color: white; padding: 1rem; border-radius: 4px; margin-bottom: 1rem; }
        .btn { background: #007bff; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px; border: none; cursor: pointer; }
        .btn-danger { background: #dc3545; }
        .nav { background: #333; color: white; padding: 10px; margin-bottom: 2rem; }
        .nav a { color: white; text-decoration: none; margin-right: 20px; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">Vibe Shop</a>
        <a href="/products">Products</a>
        <a href="/debug">Debug Mode</a>
    </div>
    
    <div class="debug-container">
        <div class="debug-header">
            <h1>🔧 DEBUG MODE</h1>
            <p>System configuration and diagnostic tools</p>
        </div>

        <div class="debug-section">
            <h3>📊 Session Information</h3>
            <p><strong>Cart Items:</strong> ` + fmt.Sprintf("%+v", cartItems) + `</p>
            <p><strong>Simple Test:</strong> ` + fmt.Sprintf("%+v", session.Values["simple_test"]) + `</p>
        </div>

        <div class="config-form">
            <h3>🔧 Server Configuration</h3>
            <form method="GET" action="/debug">
                <input type="text" name="config" placeholder="Enter configuration command..." value="` + configCmd + `" style="width: 400px; padding: 0.5rem; margin-right: 1rem;">
                <button type="submit" class="btn btn-danger">Execute Command</button>
            </form>
            
            <h4>💡 Common Commands:</h4>
            <ul>
                <li><code>whoami</code> - Check current user</li>
                <li><code>ls -la</code> - List files</li>
                <li><code>cat /etc/passwd</code> - Read system users</li>
                <li><code>ps aux</code> - List processes</li>
                <li><code>netstat -an</code> - Network connections</li>
            </ul>
        </div>`

	if configOutput != "" {
		html += `<div class="config-output">
            <h4>🔴 CONFIG COMMAND OUTPUT:</h4>
            ` + configOutput + `
        </div>`
	}

	html += `
        <div class="debug-section">
            <h3>🔗 Quick Access Links</h3>
            <p><a href="/debug?config=whoami" class="btn btn-danger">Check Current User</a></p>
            <p><a href="/debug?config=ls%20-la" class="btn btn-danger">List Files</a></p>
            <p><a href="/debug?config=cat%20/etc/passwd" class="btn btn-danger">Read System Users</a></p>
            <p><a href="/debug?config=ps%20aux" class="btn btn-danger">List Processes</a></p>
            <p><a href="/debug?config=netstat%20-an" class="btn btn-danger">Network Connections</a></p>
        </div>
    </div>
</body>
</html>`

	w.Write([]byte(html))
}

func databaseHandler(w http.ResponseWriter, r *http.Request) {
	// VULNERABLE: Direct database access without authentication
	query := r.URL.Query().Get("query")
	var result string

	if query != "" {
		// CRITICAL VULNERABILITY: Direct SQL execution without sanitization
		rows, err := db.Query(query)
		if err != nil {
			result = "Error executing query: " + err.Error()
		} else {
			defer rows.Close()

			// Get column names
			columns, err := rows.Columns()
			if err != nil {
				result = "Error getting columns: " + err.Error()
			} else {
				// Create slice to hold values
				values := make([]interface{}, len(columns))
				valuePtrs := make([]interface{}, len(columns))
				for i := range values {
					valuePtrs[i] = &values[i]
				}

				// Build result string
				result = "Query: " + query + "\n\nResults:\n"
				result += "Columns: " + fmt.Sprintf("%v", columns) + "\n\n"

				rowNum := 1
				for rows.Next() {
					err := rows.Scan(valuePtrs...)
					if err != nil {
						result += "Error scanning row: " + err.Error() + "\n"
					} else {
						result += fmt.Sprintf("Row %d: %v\n", rowNum, values)
						rowNum++
					}
				}
			}
		}
	}

	w.Header().Set("Content-Type", "text/html")
	html := `<!DOCTYPE html>
<html>
<head>
    <title>Database Interface - Vibe E-commerce</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; background: #f0f0f0; }
        .db-container { max-width: 1200px; margin: 0 auto; background: white; padding: 2rem; border-radius: 8px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        .db-header { background: #28a745; color: white; padding: 1rem; border-radius: 4px; margin-bottom: 2rem; }
        .db-section { background: #f8f9fa; padding: 1rem; border-radius: 4px; margin-bottom: 1rem; }
        .query-form { background: #fff3cd; border: 1px solid #ffeaa7; padding: 1rem; border-radius: 4px; margin-bottom: 2rem; }
        .query-output { background: #dc3545; color: white; padding: 1rem; border-radius: 4px; font-family: monospace; white-space: pre-wrap; margin-top: 1rem; }
        .btn { background: #007bff; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px; border: none; cursor: pointer; }
        .btn-success { background: #28a745; }
        .btn-danger { background: #dc3545; }
        .nav { background: #333; color: white; padding: 10px; margin-bottom: 2rem; }
        .nav a { color: white; text-decoration: none; margin-right: 20px; }
        .table-info { background: #e9ecef; padding: 1rem; border-radius: 4px; margin-bottom: 1rem; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">Vibe Shop</a>
        <a href="/products">Products</a>
        <a href="/owner/debug">Debug</a>
        <a href="/database">Database</a>
    </div>
    
    <div class="db-container">
        <div class="db-header">
            <h1>🗄️ Database Interface</h1>
            <p>Direct SQL query execution and database management</p>
        </div>

        <div class="db-section">
            <h3>📊 Database Tables</h3>
            <div class="table-info">
                <p><strong>Available Tables:</strong></p>
                <ul>
                    <li><code>users</code> - User accounts and PII data</li>
                    <li><code>products</code> - Product catalog</li>
                    <li><code>orders</code> - Customer orders</li>
                    <li><code>order_items</code> - Order line items</li>
                    <li><code>payments</code> - Payment information</li>
                </ul>
            </div>
        </div>

        <div class="query-form">
            <h3>🔍 SQL Query Interface</h3>
            <form method="GET" action="/database">
                <textarea name="query" placeholder="Enter SQL query..." style="width: 100%; height: 100px; padding: 0.5rem; margin-bottom: 1rem; font-family: monospace;">` + query + `</textarea>
                <button type="submit" class="btn btn-success">Execute Query</button>
            </form>
            
            <h4>💡 Example Queries:</h4>
            <ul>
                <li><code>SELECT * FROM users</code> - View all users and PII</li>
                <li><code>SELECT * FROM orders</code> - View all orders</li>
                <li><code>SELECT * FROM payments</code> - View payment data</li>
                <li><code>SELECT name, price FROM products</code> - View products</li>
                <li><code>DELETE FROM users WHERE id = 1</code> - Delete user (DANGEROUS!)</li>
                <li><code>DROP TABLE users</code> - Drop table (DANGEROUS!)</li>
            </ul>
        </div>`

	if result != "" {
		html += `<div class="query-output">
            <h4>🔴 QUERY RESULTS:</h4>
            ` + result + `
        </div>`
	}

	html += `
        <div class="db-section">
            <h3>🔗 Quick Access Links</h3>
            <p><a href="/database?query=SELECT%20*%20FROM%20users" class="btn btn-danger">View All Users</a></p>
            <p><a href="/database?query=SELECT%20*%20FROM%20orders" class="btn btn-danger">View All Orders</a></p>
            <p><a href="/database?query=SELECT%20*%20FROM%20payments" class="btn btn-danger">View All Payments</a></p>
            <p><a href="/database?query=SELECT%20*%20FROM%20products" class="btn btn-danger">View All Products</a></p>
        </div>
    </div>
</body>
</html>`

	w.Write([]byte(html))
}

// INSECURE: Admin handlers without authentication (major vulnerability!)
func insecureAdminHandler(w http.ResponseWriter, r *http.Request) {
	// VULNERABLE: No authentication required
	w.Header().Set("Content-Type", "text/html")
	html := `<!DOCTYPE html>
<html>
<head>
    <title>INSECURE Admin Dashboard - Vibe E-commerce</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .nav { background: #dc3545; color: white; padding: 10px; }
        .nav a { color: white; text-decoration: none; margin-right: 20px; }
        .container { max-width: 800px; margin: 0 auto; }
        .admin-section { background: white; padding: 1.5rem; margin: 1rem 0; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        .btn { background: #007bff; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px; display: inline-block; margin: 0.5rem; }
        .btn:hover { background: #0056b3; }
        .btn-danger { background: #dc3545; }
        .btn-danger:hover { background: #c82333; }
        .vulnerability-warning { background: #dc3545; color: white; padding: 1rem; border-radius: 4px; margin-bottom: 1rem; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">Vibe Shop</a>
        <a href="/admin/insecure">INSECURE Admin Dashboard</a>
        <a href="/admin/products/insecure">INSECURE Manage Products</a>
        <a href="/payments/insecure">INSECURE Payments</a>
    </div>
    <div class="container">
        <div class="vulnerability-warning">
            <h3>🚨 CRITICAL VULNERABILITY: NO AUTHENTICATION REQUIRED</h3>
            <p>This admin interface is accessible without any authentication!</p>
        </div>
        
        <h2>INSECURE Admin Dashboard</h2>
        
        <div class="admin-section">
            <h3>Product Management</h3>
            <p>Add, edit, and delete products from the catalog (NO AUTH REQUIRED).</p>
            <a href="/admin/products/insecure" class="btn btn-danger">INSECURE Manage Products</a>
        </div>
        
        <div class="admin-section">
            <h3>Payment Management</h3>
            <p>View and manage payment information (NO AUTH REQUIRED).</p>
            <a href="/payments/insecure" class="btn btn-danger">INSECURE View Payments</a>
        </div>
        
        <div class="admin-section">
            <h3>Quick Actions</h3>
            <a href="/products" class="btn">View Products</a>
            <a href="/" class="btn">Go to Home</a>
        </div>
    </div>
</body>
</html>`
	w.Write([]byte(html))
}

func insecureAdminProductsHandler(w http.ResponseWriter, r *http.Request) {
	// VULNERABLE: No authentication required
	rows, err := db.Query("SELECT id, name, description, price, image_url FROM products ORDER BY id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.ImageURL)
		if err != nil {
			continue
		}
		products = append(products, p)
	}

	w.Header().Set("Content-Type", "text/html")
	html := `<!DOCTYPE html>
<html>
<head>
    <title>INSECURE Product Management - Vibe E-commerce</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .nav { background: #dc3545; color: white; padding: 10px; }
        .nav a { color: white; text-decoration: none; margin-right: 20px; }
        .container { max-width: 1200px; margin: 0 auto; }
        .vulnerability-warning { background: #dc3545; color: white; padding: 1rem; border-radius: 4px; margin-bottom: 1rem; }
        .product-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(300px, 1fr)); gap: 1rem; }
        .product-card { background: white; padding: 1rem; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        .btn { background: #007bff; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px; }
        .btn-danger { background: #dc3545; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">Vibe Shop</a>
        <a href="/admin/insecure">INSECURE Admin Dashboard</a>
        <a href="/admin/products/insecure">INSECURE Manage Products</a>
    </div>
    <div class="container">
        <div class="vulnerability-warning">
            <h3>🚨 CRITICAL VULNERABILITY: NO AUTHENTICATION REQUIRED</h3>
            <p>This product management interface is accessible without any authentication!</p>
        </div>
        
        <h2>INSECURE Product Management</h2>
        <p>Total Products: ` + strconv.Itoa(len(products)) + `</p>
        
        <div class="product-grid">`

	for _, product := range products {
		html += `<div class="product-card">
            <h3>` + product.Name + `</h3>
            <p>` + product.Description + `</p>
            <p><strong>Price:</strong> $` + fmt.Sprintf("%.2f", product.Price) + `</p>
            <p><strong>ID:</strong> ` + strconv.Itoa(product.ID) + `</p>
        </div>`
	}

	html += `
        </div>
    </div>
</body>
</html>`
	w.Write([]byte(html))
}

func insecurePaymentsHandler(w http.ResponseWriter, r *http.Request) {
	// VULNERABLE: No authentication required - direct access to payment data
	rows, err := db.Query(`
		SELECT p.id, p.order_id, p.card_number, p.card_holder_name, 
			   p.expiration_month, p.expiration_year, p.cvv,
			   p.billing_address, p.phone, p.ssn, p.driver_license, p.passport_number,
			   o.total, u.email
		FROM payments p
		JOIN orders o ON p.order_id = o.id
		JOIN users u ON o.user_id = u.id
		ORDER BY p.id DESC
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type PaymentWithPII struct {
		ID              int
		OrderID         int
		CardNumber      string
		CardHolderName  string
		ExpirationMonth string
		ExpirationYear  string
		CVV             string
		BillingAddress  string
		Phone           string
		SSN             string
		DriverLicense   string
		PassportNumber  string
		Total           float64
		UserEmail       string
	}

	var payments []PaymentWithPII
	for rows.Next() {
		var payment PaymentWithPII
		err := rows.Scan(&payment.ID, &payment.OrderID, &payment.CardNumber, &payment.CardHolderName,
			&payment.ExpirationMonth, &payment.ExpirationYear, &payment.CVV,
			&payment.BillingAddress, &payment.Phone, &payment.SSN, &payment.DriverLicense, &payment.PassportNumber,
			&payment.Total, &payment.UserEmail)
		if err != nil {
			continue
		}
		payments = append(payments, payment)
	}

	w.Header().Set("Content-Type", "text/html")
	html := `<!DOCTYPE html>
<html>
<head>
    <title>INSECURE Payment Data - Vibe E-commerce</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .nav { background: #dc3545; color: white; padding: 10px; }
        .nav a { color: white; text-decoration: none; margin-right: 20px; }
        .container { max-width: 1200px; margin: 0 auto; }
        .vulnerability-warning { background: #dc3545; color: white; padding: 1rem; border-radius: 4px; margin-bottom: 1rem; }
        .payment-card { background: white; padding: 1rem; margin: 1rem 0; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        .pii-data { background: #fff3cd; border: 1px solid #ffeaa7; padding: 1rem; border-radius: 4px; margin: 1rem 0; }
        .btn { background: #007bff; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">Vibe Shop</a>
        <a href="/admin/insecure">INSECURE Admin Dashboard</a>
        <a href="/admin/products/insecure">INSECURE Manage Products</a>
        <a href="/payments/insecure">INSECURE Payments</a>
    </div>
    <div class="container">
        <div class="vulnerability-warning">
            <h3>🚨 CRITICAL VULNERABILITY: NO AUTHENTICATION REQUIRED</h3>
            <p>This payment data interface is accessible without any authentication!</p>
            <p>Complete payment information including credit card numbers, CVV, SSN, and more are exposed!</p>
        </div>
        
        <h2>INSECURE Payment Data Exposure</h2>
        <p>Total Payment Records: ` + strconv.Itoa(len(payments)) + `</p>`

	for _, payment := range payments {
		html += `<div class="payment-card">
            <h3>Payment ID: ` + strconv.Itoa(payment.ID) + `</h3>
            <div class="pii-data">
                <h4>🔴 EXPOSED PII DATA:</h4>
                <p><strong>Order ID:</strong> ` + strconv.Itoa(payment.OrderID) + `</p>
                <p><strong>Card Number:</strong> ` + payment.CardNumber + `</p>
                <p><strong>Card Holder:</strong> ` + payment.CardHolderName + `</p>
                <p><strong>Expiry:</strong> ` + payment.ExpirationMonth + `/` + payment.ExpirationYear + `</p>
                <p><strong>CVV:</strong> ` + payment.CVV + `</p>
                <p><strong>Billing Address:</strong> ` + payment.BillingAddress + `</p>
                <p><strong>Phone:</strong> ` + payment.Phone + `</p>
                <p><strong>SSN:</strong> ` + payment.SSN + `</p>
                <p><strong>Driver License:</strong> ` + payment.DriverLicense + `</p>
                <p><strong>Passport Number:</strong> ` + payment.PassportNumber + `</p>
                <p><strong>Total:</strong> $` + fmt.Sprintf("%.2f", payment.Total) + `</p>
                <p><strong>User Email:</strong> ` + payment.UserEmail + `</p>
            </div>
        </div>`
	}

	html += `
    </div>
</body>
</html>`
	w.Write([]byte(html))
}

// INSECURE: Owner payments handler without authentication (major vulnerability!)
func insecureOwnerPaymentsHandler(w http.ResponseWriter, r *http.Request) {
	// VULNERABLE: No authentication required - direct access to owner payment data
	rows, err := db.Query(`
		SELECT p.id, p.order_id, p.card_number, p.card_holder_name, 
			   p.expiration_month, p.expiration_year, p.cvv,
			   p.billing_address, p.phone, p.ssn, p.driver_license, p.passport_number,
			   o.total, u.email
		FROM payments p
		JOIN orders o ON p.order_id = o.id
		JOIN users u ON o.user_id = u.id
		ORDER BY p.id DESC
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type PaymentWithPII struct {
		ID              int
		OrderID         int
		CardNumber      string
		CardHolderName  string
		ExpirationMonth string
		ExpirationYear  string
		CVV             string
		BillingAddress  string
		Phone           string
		SSN             string
		DriverLicense   string
		PassportNumber  string
		Total           float64
		UserEmail       string
	}

	var payments []PaymentWithPII
	for rows.Next() {
		var payment PaymentWithPII
		err := rows.Scan(&payment.ID, &payment.OrderID, &payment.CardNumber, &payment.CardHolderName,
			&payment.ExpirationMonth, &payment.ExpirationYear, &payment.CVV,
			&payment.BillingAddress, &payment.Phone, &payment.SSN, &payment.DriverLicense, &payment.PassportNumber,
			&payment.Total, &payment.UserEmail)
		if err != nil {
			continue
		}
		payments = append(payments, payment)
	}

	w.Header().Set("Content-Type", "text/html")
	html := `<!DOCTYPE html>
<html>
<head>
    <title>INSECURE Owner Payment Data - Vibe E-commerce</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .nav { background: #dc3545; color: white; padding: 10px; }
        .nav a { color: white; text-decoration: none; margin-right: 20px; }
        .container { max-width: 1200px; margin: 0 auto; }
        .vulnerability-warning { background: #dc3545; color: white; padding: 1rem; border-radius: 4px; margin-bottom: 1rem; }
        .payment-card { background: white; padding: 1rem; margin: 1rem 0; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        .pii-data { background: #fff3cd; border: 1px solid #ffeaa7; padding: 1rem; border-radius: 4px; margin: 1rem 0; }
        .btn { background: #007bff; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px; }
        .owner-section { background: #e3f2fd; border: 1px solid #2196f3; padding: 1rem; border-radius: 4px; margin: 1rem 0; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">Vibe Shop</a>
        <a href="/admin/insecure">INSECURE Admin Dashboard</a>
        <a href="/owner/payments/insecure">INSECURE Owner Payments</a>
        <a href="/owner/debug">Owner Debug</a>
        <a href="/owner/database">Owner Database</a>
    </div>
    <div class="container">
        <div class="vulnerability-warning">
            <h3>🚨 CRITICAL VULNERABILITY: OWNER ACCESS WITHOUT AUTHENTICATION</h3>
            <p>This owner payment data interface is accessible without any authentication!</p>
            <p>Complete payment information including credit card numbers, CVV, SSN, and more are exposed!</p>
        </div>
        
        <div class="owner-section">
            <h3>👑 OWNER-LEVEL ACCESS GRANTED</h3>
            <p>This interface provides owner-level access to all payment data without any authentication requirements.</p>
            <p>This demonstrates a complete privilege escalation vulnerability.</p>
        </div>
        
        <h2>INSECURE Owner Payment Data Exposure</h2>
        <p>Total Payment Records: ` + strconv.Itoa(len(payments)) + `</p>`

	for _, payment := range payments {
		html += `<div class="payment-card">
            <h3>Payment ID: ` + strconv.Itoa(payment.ID) + `</h3>
            <div class="pii-data">
                <h4>🔴 EXPOSED OWNER PII DATA:</h4>
                <p><strong>Order ID:</strong> ` + strconv.Itoa(payment.OrderID) + `</p>
                <p><strong>Card Number:</strong> ` + payment.CardNumber + `</p>
                <p><strong>Card Holder:</strong> ` + payment.CardHolderName + `</p>
                <p><strong>Expiry:</strong> ` + payment.ExpirationMonth + `/` + payment.ExpirationYear + `</p>
                <p><strong>CVV:</strong> ` + payment.CVV + `</p>
                <p><strong>Billing Address:</strong> ` + payment.BillingAddress + `</p>
                <p><strong>Phone:</strong> ` + payment.Phone + `</p>
                <p><strong>SSN:</strong> ` + payment.SSN + `</p>
                <p><strong>Driver License:</strong> ` + payment.DriverLicense + `</p>
                <p><strong>Passport Number:</strong> ` + payment.PassportNumber + `</p>
                <p><strong>Total:</strong> $` + fmt.Sprintf("%.2f", payment.Total) + `</p>
                <p><strong>User Email:</strong> ` + payment.UserEmail + `</p>
            </div>
        </div>`
	}

	html += `
    </div>
</body>
</html>`
	w.Write([]byte(html))
}

func paymentsHandler(w http.ResponseWriter, r *http.Request) {
	ownerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		// Get all payment information with PII (for educational purposes - this is intentionally insecure!)
		rows, err := db.Query(`
			SELECT p.id, p.order_id, p.card_number, p.card_holder_name, 
			   p.expiration_month, p.expiration_year, p.cvv,
				   p.billing_address, p.phone, p.ssn, p.driver_license, p.passport_number,
				   o.total, u.email
			FROM payments p
			JOIN orders o ON p.order_id = o.id
			JOIN users u ON o.user_id = u.id
			ORDER BY p.id DESC
		`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		type PaymentWithPII struct {
			ID              int
			OrderID         int
			CardNumber      string
			CardHolderName  string
			ExpirationMonth string
			ExpirationYear  string
			CVV             string
			BillingAddress  string
			Phone           string
			SSN             string
			DriverLicense   string
			PassportNumber  string
			Total           float64
			UserEmail       string
		}

		var payments []PaymentWithPII
		for rows.Next() {
			var payment PaymentWithPII
			err := rows.Scan(&payment.ID, &payment.OrderID, &payment.CardNumber, &payment.CardHolderName,
				&payment.ExpirationMonth, &payment.ExpirationYear, &payment.CVV,
				&payment.BillingAddress, &payment.Phone, &payment.SSN, &payment.DriverLicense, &payment.PassportNumber,
				&payment.Total, &payment.UserEmail)
			if err != nil {
				continue
			}
			payments = append(payments, payment)
		}

		// Generate HTML response with PII data
		w.Header().Set("Content-Type", "text/html")
		html := `<!DOCTYPE html>
<html>
<head>
    <title>Payment Information - Vibe E-commerce</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .nav { background: #333; color: white; padding: 10px; }
        .nav a { color: white; text-decoration: none; margin-right: 20px; }
        .container { max-width: 1400px; margin: 0 auto; }
        .payment-item { background: white; padding: 1rem; margin: 1rem 0; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        .payment-id { font-weight: bold; color: #007bff; }
        .payment-total { font-size: 1.2rem; font-weight: bold; color: #28a745; }
        .payment-date { color: #6c757d; font-size: 0.9rem; }
        .pii-data { background: #f8f9fa; padding: 10px; margin: 10px 0; border-radius: 5px; border-left: 4px solid #dc3545; }
        .pii-label { font-weight: bold; color: #dc3545; }
        .card-data { background: #fff3cd; padding: 10px; margin: 10px 0; border-radius: 5px; border-left: 4px solid #ffc107; }
        .card-label { font-weight: bold; color: #856404; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">Vibe Shop</a>
        <a href="/products">Products</a>
        <a href="/cart">Cart</a>
        <a href="/orders">Orders</a>
        <a href="/logout">Logout</a>
    </div>
    <div class="container">
        <h2>Payment Information (INSECURE: Shows PII Data)</h2>`

		for _, payment := range payments {
			html += fmt.Sprintf(`<div class="payment-item">
                <div class="payment-id">Payment #%d (Order #%d)</div>
                <div class="payment-total">$%.2f</div>
                <div class="card-data">
                    <div><span class="card-label">Card Number:</span> %s</div>
                    <div><span class="card-label">Card Holder:</span> %s</div>
                    <div><span class="card-label">Expiry:</span> %s/%s</div>
                    <div><span class="card-label">CVV:</span> %s</div>
                </div>
                <div class="pii-data">
                    <div><span class="pii-label">Customer Email:</span> %s</div>
                    <div><span class="pii-label">Billing Address:</span> %s</div>
                    <div><span class="pii-label">Phone:</span> %s</div>
                    <div><span class="pii-label">SSN:</span> %s</div>
                    <div><span class="pii-label">Driver License:</span> %s</div>
                    <div><span class="pii-label">Passport Number:</span> %s</div>
                </div>
            </div>`, payment.ID, payment.OrderID, payment.Total, payment.CardNumber, payment.CardHolderName, payment.ExpirationMonth, payment.ExpirationYear, payment.CVV, payment.UserEmail, payment.BillingAddress, payment.Phone, payment.SSN, payment.DriverLicense, payment.PassportNumber)
		}

		html += `</div>
</body>
</html>`

		w.Write([]byte(html))
	})(w, r)
}

// VULNERABLE: Insecure Direct Object Reference (IDOR) - Profile Page
// Allows access to any user's profile by changing the ID parameter
func profileHandler(w http.ResponseWriter, r *http.Request) {
	// VULNERABLE: No authentication check - anyone can access any profile
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "User ID required", http.StatusBadRequest)
		return
	}

	// VULNERABLE: Direct ID usage without authorization check
	rows, err := db.Query(`
		SELECT id, email, first_name, last_name, phone, ssn, date_of_birth, 
		       address, city, state, zip_code, credit_card_number, credit_card_expiry, 
		       credit_card_cvv, is_user, is_admin, is_owner
		FROM users 
		WHERE id = ?
	`, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type UserProfile struct {
		ID               int
		Email            string
		FirstName        string
		LastName         string
		Phone            string
		SSN              string
		DateOfBirth      string
		Address          string
		City             string
		State            string
		ZipCode          string
		CreditCardNumber string
		CreditCardExpiry string
		CreditCardCVV    string
		IsUser           bool
		IsAdmin          bool
		IsOwner          bool
	}

	var user UserProfile
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Phone, &user.SSN, &user.DateOfBirth, &user.Address, &user.City, &user.State, &user.ZipCode, &user.CreditCardNumber, &user.CreditCardExpiry, &user.CreditCardCVV, &user.IsUser, &user.IsAdmin, &user.IsOwner)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Determine role based on ID scheme
	var role string
	if user.ID == 0 {
		role = "Owner"
	} else if user.ID >= 1000 {
		role = "Admin"
	} else {
		role = "Customer"
	}

	w.Header().Set("Content-Type", "text/html")
	html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>User Profile - Vibe E-commerce</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .nav { background: #333; color: white; padding: 10px; }
        .nav a { color: white; text-decoration: none; margin-right: 20px; }
        .container { max-width: 800px; margin: 0 auto; }
        .profile-card { background: white; padding: 2rem; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        .profile-header { border-bottom: 2px solid #007bff; padding-bottom: 1rem; margin-bottom: 1rem; }
        .profile-name { font-size: 2rem; font-weight: bold; color: #007bff; }
        .profile-role { font-size: 1.2rem; color: #6c757d; margin-bottom: 1rem; }
        .profile-section { margin: 1rem 0; }
        .section-title { font-weight: bold; color: #495057; border-bottom: 1px solid #dee2e6; padding-bottom: 0.5rem; }
        .pii-data { background: #fff3cd; border: 1px solid #ffeaa7; padding: 1rem; border-radius: 4px; margin: 1rem 0; }
        .pii-label { font-weight: bold; color: #856404; }
        .financial-data { background: #f8d7da; border: 1px solid #f5c6cb; padding: 1rem; border-radius: 4px; margin: 1rem 0; }
        .financial-label { font-weight: bold; color: #721c24; }

    </style>
</head>
<body>
    <div class="nav">
        <a href="/">Vibe Shop</a>
        <a href="/products">Products</a>
        <a href="/profile?id=1">My Profile</a>
        <a href="/logout">Logout</a>
    </div>
    <div class="container">
        <div class="profile-card">
            <div class="profile-header">
                <div class="profile-name">%s %s</div>
                <div class="profile-role">%s (ID: %d)</div>
                <div>Email: %s</div>
            </div>

            <div class="profile-section">
                <div class="section-title">Personal Information</div>
                <div class="pii-data">
                    <div><span class="pii-label">Phone:</span> %s</div>
                    <div><span class="pii-label">SSN:</span> %s</div>
                    <div><span class="pii-label">Date of Birth:</span> %s</div>
                    <div><span class="pii-label">Address:</span> %s</div>
                    <div><span class="pii-label">City:</span> %s</div>
                    <div><span class="pii-label">State:</span> %s</div>
                    <div><span class="pii-label">Zip Code:</span> %s</div>
                </div>
            </div>

            <div class="profile-section">
                <div class="section-title">Financial Information</div>
                <div class="financial-data">
                    <div><span class="financial-label">Credit Card Number:</span> %s</div>
                    <div><span class="financial-label">Expiry Date:</span> %s</div>
                    <div><span class="financial-label">CVV:</span> %s</div>
                </div>
            </div>

            <div class="profile-section">
                <div class="section-title">Account Information</div>
                <div>
                    <div><strong>User ID:</strong> %d</div>
                    <div><strong>Is User:</strong> %t</div>
                    <div><strong>Is Admin:</strong> %t</div>
                    <div><strong>Is Owner:</strong> %t</div>
                </div>
            </div>
        </div>
    </div>
</body>
</html>`, user.FirstName, user.LastName, role, user.ID, user.Email, user.Phone, user.SSN, user.DateOfBirth, user.Address, user.City, user.State, user.ZipCode, user.CreditCardNumber, user.CreditCardExpiry, user.CreditCardCVV, user.ID, user.IsUser, user.IsAdmin, user.IsOwner)

	w.Write([]byte(html))
}

// Secret flag handler - exposes sensitive files for the challenge
func secretFlagHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("file")

	if filePath == "" {
		// Show available files for the challenge
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>System Files</title>
			<style>
				body { font-family: Arial, sans-serif; margin: 40px; }
				.file-list { background: #f5f5f5; padding: 20px; border-radius: 5px; }
				.file-item { margin: 10px 0; }
				.file-link { color: #0066cc; text-decoration: none; }
				.file-link:hover { text-decoration: underline; }
				.note { background: #fff3cd; padding: 10px; border-radius: 5px; margin: 20px 0; }
			</style>
		</head>
		<body>
			<h1>🔍 System Files Explorer</h1>
			<div class="note">
				<strong>Challenge:</strong> Find the secret flag by exploring system files. 
				The flag is the MD5 hash of "Vib3Sec".
			</div>
			<div class="file-list">
				<h3>Available Files:</h3>
				<div class="file-item">
					<a href="/secret?file=config.txt" class="file-link">📄 config.txt</a>
				</div>
				<div class="file-item">
					<a href="/secret?file=secrets.txt" class="file-link">🔐 secrets.txt</a>
				</div>
				<div class="file-item">
					<a href="/secret?file=flag.txt" class="file-link">🚩 flag.txt</a>
				</div>
				<div class="file-item">
					<a href="/secret?file=backup.txt" class="file-link">💾 backup.txt</a>
				</div>
			</div>
		</body>
		</html>`
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
		return
	}

	// Simulate file content based on requested file
	var content string
	switch filePath {
	case "config.txt":
		content = `# Application Configuration
DEBUG_MODE=true
DATABASE_URL=sqlite:///database/ecommerce.db
SECRET_KEY=insecure-secret-key
ADMIN_EMAIL=admin@example.com
FLAG_HINT=Check the secrets file for more information`
	case "secrets.txt":
		content = `# Secret Information
API_KEY=sk-1234567890abcdef
DATABASE_PASSWORD=admin123
ENCRYPTION_KEY=weak-key-123
FLAG_LOCATION=flag.txt
SECRET_MESSAGE=The flag is hidden in plain sight`
	case "flag.txt":
		content = `🎉 CONGRATULATIONS! 🎉
You found the secret flag!
Flag: 0b8dcaf09bee1fd3c2143ea0ffba4fa2

This is the MD5 hash of "Vib3Sec".
You successfully exploited the file exposure vulnerability!`
	case "backup.txt":
		content = `# Backup Information
BACKUP_DATE=2024-01-15
BACKUP_LOCATION=/backups/
BACKUP_SIZE=1.2GB
FLAG_HINT=The flag is not here, try flag.txt`
	default:
		content = "File not found: " + filePath
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(content))
}
