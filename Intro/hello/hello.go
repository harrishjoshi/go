package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//var name string
	//name = "Your Name"
	//fmt.Println(name)

	// A slice of names.
	names := []string{"Harish", "Jhon", "Darrin"}
	messages, err := greetings.Hellos(names)

	// Request a greeting message.
	//message, err := greetings.Hello("Harish")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the greeting message
	// to the console.
	fmt.Println(messages)
}
