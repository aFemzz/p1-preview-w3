package handler

import (
	"fmt"
	"log"
	"pagi/config"
	"pagi/entity"
)

func BooksGetList() {
	// Establish a connection to the database
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	// Define the query to retrieve books by author 'Jane Smith'
	query := `
		SELECT book_title, book_type
		FROM book
		JOIN author ON book.author_id = author.author_id
		WHERE author.author_name = 'Jane Smith';
	`

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	defer rows.Close()

	// Initialize a slice to store the list of books
	var listBooks []entity.Books

	// Iterate over the rows
	for rows.Next() {
		// Create a new Books struct to store book data
		var b entity.Books

		// Scan the values from the row into the Books struct fields
		if err := rows.Scan(&b.Book_Name, &b.Book_Type); err != nil {
			log.Fatal("Error scanning row:", err)
		}

		// Append the Books struct to the listBooks slice
		listBooks = append(listBooks, b)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Fatal("Error iterating over rows:", err)
	}

	// Print the list of books
	fmt.Println("=========================")
	fmt.Println("BOOKS LIST BY JANE SMITH")
	fmt.Println("=========================")
	for _, book := range listBooks {
		fmt.Printf("Book Title: %s, Book Type: %s\n", book.Book_Name, book.Book_Type)
	}

}
