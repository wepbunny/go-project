package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Order represents an order in the database.
type Order struct {
	ID          int
	Title       string
	Image       string
	Price       float64
	Quantity    int
	Description string
	NumLikes    int
}

// InitializeDB initializes the database connection.
func InitializeDB() error {
	// Replace the connection details with your actual MySQL connection details.
	var err error
	db, err = sql.Open("mysql", "root:root_password@tcp(localhost:3306)/excelp")
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping the database: %v", err)
	}

	log.Println("Connected to the database")

	return nil
}

// CloseDB closes the database connection.
func CloseDB() {
	if db != nil {
		db.Close()
		log.Println("Closed the database connection")
	}
}

// GetOrders retrieves all orders from the database.
func GetOrders() ([]Order, error) {
	query := "SELECT id, title, image, price, quantity, description, 0 as numlikes FROM orders"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.ID,
			&order.Title,
			&order.Image,
			&order.Price,
			&order.Quantity,
			&order.Description,
			&order.NumLikes,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("encountered an error while iterating over rows: %v", err)
	}

	return orders, nil
}

// GetHighLikesOrders retrieves orders with amount > 100 and likes > 10 from the database.
func GetHighLikesOrders() ([]Order, error) {
	query := `
		SELECT o.id, o.title, o.image, o.price, o.quantity, o.description, count(l.id) as numlikes
		FROM orders o
		INNER JOIN likes l ON o.id = l.order_id
		GROUP BY o.id
		HAVING numlikes > 3
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.ID,
			&order.Title,
			&order.Image,
			&order.Price,
			&order.Quantity,
			&order.Description,
			&order.NumLikes,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("encountered an error while iterating over rows: %v", err)
	}

	return orders, nil
}

// GetLowLikesOrders retrieves orders with amount < 100 and likes < 3 from the database.
func GetLowLikesOrders() ([]Order, error) {
	query := `
		SELECT o.id, o.title, o.image, o.price, o.quantity, o.description, count(l.id) as numlikes
		FROM orders o
		LEFT JOIN likes l ON o.id = l.order_id
		GROUP BY o.id
		HAVING numlikes < 3
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.ID,
			&order.Title,
			&order.Image,
			&order.Price,
			&order.Quantity,
			&order.Description,
			&order.NumLikes,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("encountered an error while iterating over rows: %v", err)
	}

	return orders, nil
}
