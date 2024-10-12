package main

import (
    "bufio"    // For reading input with a buffered reader / Чтение данных через буфер
    "fmt"      // For formatted I/O (e.g., printing to the console) / Форматированный ввод-вывод
    "os"       // For access to OS functions like standard input/output / Доступ к функциям ОС (ввод-вывод)
    "strconv"  // For converting strings to other types (like int) / Для преобразования строк в другие типы (например, в int)
    "strings"  // For string manipulation functions / Для работы со строками
)

func main() {
    // Check if the user has provided the required argument (TAG) / Проверка, что пользователь указал аргумент (TAG)
    if len(os.Args) != 2 {
        fmt.Println("Usage: ./notestool [TAG]") // Inform the user about correct usage / Информируем пользователя о правильном использовании
        return
    }

    collectionName := os.Args[1] // Get the collection name from the first argument / Получаем название коллекции из аргумента
    notesFile := collectionName + ".txt" // Create the file name for storing notes / Создаем имя файла для хранения заметок

    // Load existing notes from the file / Загружаем существующие заметки из файла
    notes := loadNotes(notesFile)

    fmt.Println("Welcome to the notes tool!\n") // Welcome message / Приветственное сообщение

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
            fmt.Println("Enter the note text:") // Prompt user to enter note text / Запрашиваем текст заметки
            note := strings.TrimSpace(readLine()) // Read the note text / Читаем текст заметки
            notes = append(notes, note) // Add the note to the list / Добавляем заметку в список
        case 3:
            notes = deleteNote(notes) // Delete a note / Удаляем заметку
        case 4:
            saveNotes(notes, notesFile) // Save notes to file before exiting / Сохраняем заметки в файл перед выходом
            fmt.Println("Exiting... Goodbye!") // Exit message / Сообщение о выходе
            return // Exit the program / Выход из программы
        }
    }
}

// Function to load notes from a file / Функция для загрузки заметок из файла
func loadNotes(filename string) []string {
    file, err := os.Open(filename) // Open the notes file / Открываем файл заметок
    if err != nil {
        if os.IsNotExist(err) { // If the file doesn't exist, return an empty list / Если файл не существует, возвращаем пустой список
            return []string{}
        }
        fmt.Println("Error opening notes file:", err) // If there's another error, print it and exit / Выводим ошибку при открытии файла и выходим
        os.Exit(1)
    }
    defer file.Close() // Ensure the file is closed after reading / Закрываем файл после чтения

    var notes []string // Slice to store notes / Срез для хранения заметок
    scanner := bufio.NewScanner(file) // Create a scanner to read the file line by line / Создаем сканер для построчного чтения файла
    for scanner.Scan() {
        notes = append(notes, scanner.Text()) // Add each line as a note / Добавляем каждую строку как заметку
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading notes:", err) // If there's an error while reading, print it / Выводим ошибку чтения файла
        os.Exit(1)
    }

    return notes // Return the list of notes / Возвращаем список заметок
}

// Function to save notes to a file / Функция для сохранения заметок в файл
func saveNotes(notes []string, filename string) {
    file, err := os.Create(filename) // Create or overwrite the file / Создаем или перезаписываем файл
    if err != nil {
        fmt.Println("Error creating notes file:", err) // Print error if file creation fails / Выводим ошибку, если файл не удалось создать
        os.Exit(1)
    }
    defer file.Close() // Ensure the file is closed after writing / Закрываем файл после записи

    for _, note := range notes { // Write each note to the file / Записываем каждую заметку в файл
        _, err := fmt.Fprintln(file, note) // Write the note as a line / Записываем заметку в виде строки
        if err != nil {
            fmt.Println("Error writing notes:", err) // Print error if writing fails / Выводим ошибку записи в файл
            os.Exit(1)
        }
    }
}

// Function to display notes / Функция для отображения заметок
func showNotes(notes []string) {
    if len(notes) == 0 { // If no notes are found / Если заметок нет
        fmt.Println("No notes found.") // Print message / Выводим сообщение
        return
    }
    fmt.Println("Notes:") // Title for the notes / Заголовок для списка заметок
    for i, note := range notes { // Iterate over the list of notes / Проходим по списку заметок
        fmt.Printf("%03d - %s\n", i+1, note) // Display the notes with numbering / Отображаем заметки с номерами
    }
}

// Function to read a line from user input / Функция для чтения строки из ввода пользователя
func readLine() string {
    reader := bufio.NewReader(os.Stdin) // Create a buffered reader / Создаем буферизированный ридер
    text, err := reader.ReadString('\n') // Read input until newline / Читаем строку до символа новой строки
    if err != nil {
        fmt.Println("Error reading input:", err) // Print error if reading fails / Выводим ошибку при чтении
        os.Exit(1)
    }
    return strings.TrimSpace(text) // Return the trimmed input / Возвращаем введенную строку без лишних пробелов
}

// Function to delete a note / Функция для удаления заметки
func deleteNote(notes []string) []string {
    fmt.Println("Enter the number of the note to remove or 0 to cancel:") // Ask for the note number to delete / Запрашиваем номер заметки для удаления
    index, err := strconv.Atoi(readLine()) // Convert input to integer / Преобразуем ввод в число
    if err != nil || index < 0 || index > len(notes) { // If invalid index / Если индекс некорректный
        fmt.Println("Invalid index.") // Print error message / Выводим сообщение об ошибке
        return notes
    }

    if index == 0 { // If user cancels / Если пользователь отменяет удаление
        return notes // Return without changes / Возвращаем исходный список заметок
    }

    return append(notes[:index-1], notes[index:]...) // Remove the selected note / Удаляем выбранную заметку
}
