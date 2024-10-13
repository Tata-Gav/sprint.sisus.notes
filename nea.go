package main

import (
	"fmt"
	"os"
)

func main() {

	// check if the number of arguments are correct, or if the user asks for help
	if len(os.Args) != 2 || os.Args[1] == "help" {
		fmt.Println("Usage: ./notestool [COLLECTION_NAME]")
		fmt.Println()
		fmt.Println("Description:\nThis command-line tool allows you to look, modify and delete your notes efficiently.")
		fmt.Println()
		fmt.Println("Arguments:\nCOLLECTION_NAME: The name of the collection of notes you want to create or access.")
		fmt.Println()
		fmt.Println("Example:\nTo create a new collection named my_notes or open an existing one named my_notes:\nType: ./notestool my_notes")
		fmt.Println()
		fmt.Println("Note: Only one argument is allowed. If no arguments or two or more arguments are provided, this help message is displayd.")
		return // exit if the arguments are invalid
	}

    fmt.Println("Welcome to the notes tool!")
	filename := os.Args[1] + ".txt" //Name of the file based on the argument
	var file *os.File //declare file variable

	//check, if the file already exists
	_, err := os.Stat(filename) //os.Stat returns nil if the file already exists
	if err != nil{
		// check if the error indicates that the file does not exist
		if os.IsNotExist(err){ //os.IsNotExist returns true if the file does not exist
			var errCreate error
			// create a new file
			file, errCreate = os.Create(filename)
			// if there's an error creating the file
			if errCreate != nil{ 
				fmt.Println("Error creating file: ", errCreate)
				return //exit if error in creating
			} 

		}else{ // if the error is not related to the file not existing
			fmt.Println("Error occurred while checking the file: ", err)
			return // Exit is there's another error
		}

	}else{ // file exists, open it for reading and writing
		var errOpen error
		file, errOpen = os.Open(filename)
		if errOpen != nil{ 
			fmt.Println("Error opening file: ", errOpen)
			return // exit if error
		}
	}
	// close the file when done
	defer file.Close()
}
