package main

import (
	"fmt"
	"os"
	"pagi/handler"
)

func main() {
	// Check if a command is provided
	if len(os.Args) < 2 {
		fmt.Println("Please specify a command")
		return // Exit the program if no command is provided
	}

	command := os.Args[1]

	switch command {
	case "books":
		handler.BooksGetList()

	case "sales":

		fmt.Println("This is sales")
	case "customers":

		fmt.Println("This is customers")
	case "topauthor":

		fmt.Println("This is topauthor")
	default:
		fmt.Println("Unknown Command:", command)
	}
}
