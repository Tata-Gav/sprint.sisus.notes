package main

import (
    "bufio"    // For reading input with a buffered reader / Чтение данных через буфер
    "fmt"      // For formatted I/O (e.g., printing to the console) / Форматированный ввод-вывод
    "os"       // For access to OS functions like standard input/output / Доступ к функциям ОС (ввод-вывод)
    "strconv"  // For converting strings to other types (like int) / Для преобразования строк в другие типы (например, в int)
    "strings"  // For string manipulation functions / Для работы со строками
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
	if err != nil {
		return []string{}
	}
	defer file.Close()

	var notes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}

	return notes
}

// Tatiana
func menuLoop(notes []string, collectionName string){
	reader := bufio.NewReader(os.Stdin) // Buffered reader to handle user input / Буферизированный ридер для работы с вводом пользователя

    for {
        // Display the menu / Отображение меню
        fmt.Println("Select operation:")
        fmt.Println("1. Show notes.")  // Option to show all notes / Показать все заметки
        fmt.Println("2. Add a note.")  // Option to add a new note / Добавить новую заметку
        fmt.Println("3. Delete a note.") // Option to delete a note / Удалить заметку
        fmt.Println("4. Exit.") // Option to exit / Выйти

        fmt.Print("Enter your choice (1-4): ")
        input, err := reader.ReadString('\n') // Read user input / Читаем ввод пользователя
        if err != nil {
            fmt.Println("Error reading input:", err) // Handle input error / Обрабатываем ошибку ввода
            continue
        }

        input = strings.TrimSpace(input) // Remove any leading/trailing spaces / Убираем лишние пробелы и символы новой строки
        choice, err := strconv.Atoi(input) // Convert input to integer / Преобразуем строку в число
        if err != nil || choice < 1 || choice > 4 {
            fmt.Println("Invalid input. Please enter a valid number (1-4).") // If input is invalid, show an error message / Показываем сообщение об ошибке при некорректном вводе
            continue
        }

        // Process user's choice / Обрабатываем выбор пользователя
        switch choice {
        case 1:
            showNotes(notes) // Show the list of notes / Показать список заметок
        case 2:
            notes = addNote(notes) // Add the note to the list / Добавляем заметку в список
        case 3:
            notes = deleteNote(notes) // Delete a note / Удаляем заметку
        case 4:
            saveNotes(notes, collectionName) // Save notes to file before exiting / Сохраняем заметки в файл перед выходом
            fmt.Println("Exiting... Goodbye!") // Exit message / Сообщение о выходе
            return // Exit the program / Выход из программы
        }
    }
}

// Tatiana
// Function to save notes to a file / Функция для сохранения заметок в файл
func saveNotes(notes []string, filename string) {
    file, err := os.Create(filename) // Create or overwrite the file / Создаем или перезаписываем файл
    if err != nil {
        fmt.Println("Error saving notes:", err) // Print error if file creation fails / Выводим ошибку, если файл не удалось создать
        return
    }
    defer file.Close() // Ensure the file is closed after writing / Закрываем файл после записи

    for _, note := range notes { // Write each note to the file / Записываем каждую заметку в файл
        _, err := fmt.Fprintln(file, note) // Write the note as a line / Записываем заметку в виде строки
        if err != nil {
            fmt.Println("Error writing notes:", err) // Print error if writing fails / Выводим ошибку записи в файл
            return
        }
    }
}

// Jukka
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

// Function to add notes
func addNote(notes []string) []string{
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







