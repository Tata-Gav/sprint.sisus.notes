package main

import (
	"bufio"   // For reading input with a buffered reader
	"fmt"     // For formatted I/O (e.g., printing to the console)
	"os"      // For access to OS functions like standard input/output
	"strconv" // For converting strings to other types (like int)
	"strings" // For string manipulation functions
)

// Function to display notes
func showNotes(notes []string) {
    if len(notes) == 0 { // Check if there are any notes
        fmt.Println("No notes available.")
        return
    }

    fmt.Println("Your notes:")
    for i, note := range notes { // Display each note with an index
        fmt.Printf("%d: %s\n", i+1, note) // Index starts from 1 for user convenience
    }
}


// Function to add a note
func addNote(notes []string) []string {
    reader := bufio.NewReader(os.Stdin) // Buffered reader to handle input
    fmt.Print("Enter your note: ")
    note, err := reader.ReadString('\n') // Read user input for the new note
    if err != nil {
        fmt.Println("Error reading note:", err) // Handle input errors
        return notes
    }

    note = strings.TrimSpace(note) // Remove any leading/trailing spaces
    if note == "" { // If the note is empty, ask for valid input
        fmt.Println("Empty note. Please enter some text.")
        return notes
    }

    notes = append(notes, note) // Add the new note to the list
    fmt.Println("Note added successfully!")
    return notes
}


// Function to delete a note
func deleteNote(notes []string) []string {
    if len(notes) == 0 { // Check if there are any notes to delete
        fmt.Println("No notes to delete.")
        return notes
    }

    showNotes(notes) // Show the current list of notes
    reader := bufio.NewReader(os.Stdin) // Buffered reader for input
    fmt.Print("Enter the number of the note to delete: ")
    input, err := reader.ReadString('\n') // Read user input for the note number
    if err != nil {
        fmt.Println("Error reading input:", err) // Handle input errors
        return notes
    }

    input = strings.TrimSpace(input)      // Remove extra spaces and newlines
    index, err := strconv.Atoi(input)     // Convert the input to an integer
    if err != nil || index < 1 || index > len(notes) { // Validate the note number
        fmt.Println("Invalid input. Please enter a valid note number.")
        return notes
    }

    // Remove the note by its index (adjust for zero-based index)
    notes = append(notes[:index-1], notes[index:]...)
    fmt.Println("Note deleted successfully!")
    return notes
}




func main() {




}
