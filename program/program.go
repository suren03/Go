package main

import (
	"fmt"
	"log"
	"suren/greetings"
)

// go mod tidy

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("Main: ")
	log.SetFlags(0)

	fmt.Println("Hello Worls")
	// fmt.Println(quote.Go())
	greetings.Hello("hgfsgfdd")
	message, err := greetings.Hello("")

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err) // you use the log package's Fatal function to print the error and stop the program. // exit status 1
	}
}
