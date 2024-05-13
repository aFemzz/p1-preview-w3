package handler

import (
	"fmt"
	"log"
	"pagi/config"
	"pagi/entity"
)

func CustomerGetList() {
	// Establish a connection to the database
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	// Define the query to retrieve books with more than one order '
	query := `
	SELECT customer_name, COUNT(DISTINCT book_id) AS num_books_ordered
	FROM bookstore_order
	GROUP BY customer_name
	HAVING num_books_ordered > 1;
	
	`

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	defer rows.Close()

	// Initialize a slice to store the data
	var customersBooks []entity.Customer

	// Iterate over the rows
	for rows.Next() {
		var c entity.Customer

		if err := rows.Scan(&c.Author, &c.Orders); err != nil {
			log.Fatal("Error scanning row:", err)
		}

		customersBooks = append(customersBooks, c)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Fatal("Error iterating over rows:", err)
	}

	// Print the list of books
	fmt.Println("=========================")
	fmt.Println("TOTAL BOOK ORDERS")
	fmt.Println("=========================")
	for _, book := range customersBooks {
		fmt.Printf("%s : %s\n", book.Author, book.Orders)

	}
	fmt.Println("")
}
