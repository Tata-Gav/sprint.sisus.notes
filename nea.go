package main

import (
    "bufio"    // For reading input with a buffered reader 
    "fmt"      // For formatted I/O (e.g., printing to the console)
    "os"       // For access to OS functions like standard input/output
    "strconv"  // For converting strings to other types (like int) 
    "strings"  // For string manipulation functions
)

// Nea
func main() {

	// check if the number of arguments are correct, or if the user asks for help
	if len(os.Args) != 2 || os.Args[1] == "help" {
		fmt.Println("Description:\nThis command-line tool allows you to read, write and delete notes efficiently.")
		fmt.Println()
		fmt.Println("Usage: ./notestool [COLLECTION_NAME]")
		fmt.Println()
		fmt.Println("Arguments:\nCOLLECTION_NAME: The name of the collection of notes you want to create or access.")
		fmt.Println()
		fmt.Println("Example:\nType: ./notestool my_notes\n to create a new collection named my_notes, or open an existing one named my_notes")
		fmt.Println()
		fmt.Println("Note: Only one argument is allowed. If no arguments or two or more arguments are provided, this help message is displayd.")
		return // exit if the arguments are invalid
	}

	collectionName := os.Args[1] + ".txt" //Name of the file based on the argument

	//create or load the collection file
	file, err := os.OpenFile(collectionName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil{
		fmt.Println("Error loading or creating the collection:", err)
		return
	}
	defer file.Close() //Ensure file is closed when done

	// Load the existing notes (empty if new file)
	notes := loadNotes(collectionName)

	fmt.Println("Welcome to the notestool!")

	menuLoop(notes, collectionName)
}

//Nea
// Function to load notes
func loadNotes(filename string) []string{
	file, err := os.Open(filename)
	//check if there was an error opening the file
	if err != nil {
		return []string{} //return empty slice if there's an error (for example the file does not exist)
	}
	//Ensure the file is closed after the function completes
	defer file.Close()

	// Initialize a slice to hold notes
	var notes []string

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	// Loop through the lines of the file
	for scanner.Scan() {
		// Append each line to the notes slice
		notes = append(notes, scanner.Text())
	}

	// Return the slice of notes loaded from the file
	return notes
}