package handler

import (
	"fmt"
	"log"
	"pagi/config"
	"pagi/entity"
)

func SalesGetList() {
	// Establish a connection to the database
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	// Define the query to retrieve sales
	query := `
	SELECT book.book_type, SUM(bookstore_order.price) AS total_sales
	FROM book
	JOIN bookstore_order ON book.book_id = bookstore_order.book_id
	GROUP BY book.book_type;
	`

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	defer rows.Close()

	// Initialize a slice to store the list of sales
	var listSales []entity.Sales

	// Iterate over the rows
	for rows.Next() {
		// Create a new sales struct to store book data
		var s entity.Sales

		// Scan the values from the row into the sales struct fields
		if err := rows.Scan(&s.Physical, &s.Ebook); err != nil {
			log.Fatal("Error scanning row:", err)
		}

		// Append the sales struct to the listSales slice
		listSales = append(listSales, s)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Fatal("Error iterating over rows:", err)
	}

	// Print the list of sales
	fmt.Println("=========================")
	fmt.Println("TOTAL SALES FOR EACH BOOK TYPE")
	fmt.Println("=========================")
	for _, sales := range listSales {
		fmt.Printf("%s : %s\n", sales.Physical, sales.Ebook)

	}
	fmt.Println("")
}
