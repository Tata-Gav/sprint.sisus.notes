package main
import (
    "bufio"    // For reading input with a buffered reader
    "fmt"      // For formatted I/O (e.g., printing to the console)
    "os"       // For access to OS functions like standard input/output
    "strconv"  // For converting strings to other types (like int)
    "strings"  // For string manipulation functions
)
func main() {
    // Create a buffered reader that reads from standard input (keyboard)
    reader := bufio.NewReader(os.Stdin)
    // Infinite loop to keep displaying the menu until the user chooses to exit
    for {
        // Display the menu options to the user
        fmt.Println("Welcome to the Menu!")
        fmt.Println("1.Show notes.")
        fmt.Println("2. Option Two")
        fmt.Println("3. Exit")
        // Prompt the user for input
        fmt.Print("Enter your choice (1-4): ")
        // Read user input as a string until the newline character '\n' is encountered
        input, _ := reader.ReadString('\n')
        // Trim any trailing spaces or newline characters from the input string
        input = strings.TrimSpace(input)
        // Convert the input string to an integer using strconv.Atoi
        // If the input is not a valid number, it will return an error
        choice, err := strconv.Atoi(input)
        if err != nil {
            // If there's an error in conversion (e.g., non-numeric input), show an error message
            fmt.Println("Invalid input. Please enter a valid number (1-4).")
            continue // Skip the rest of the loop and ask the user for input again
        }
        // Handle user input using a switch statement based on their choice
        switch choice {
        case 1:
            fmt.Println("You chose Option One.")
        case 2:
            // If the user chooses option 2, print a goodbye message and exit the loop
            fmt.Println("Exiting... Goodbye!")
            return // Return from the main function, which will end the program
        default:
            // If the user enters a number outside the valid range (1-4), show an error message
            fmt.Println("Invalid choice. Please select a valid option (1-4).")
        }
    }
}