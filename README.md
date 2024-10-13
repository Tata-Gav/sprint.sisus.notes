# Notes Tool

## Description
The Notes Tool is a command-line application, which allows user to create, read, modify and delete short one-line notes. It provides an easy-to-use interface for quickly save and manage important information.

## Usage
To use Notes Tool, execute the following command in your terminal:

Example here


- COLLECTION_NAME: This argument specifies the name of the collection of notes you want to create or access. If a collection with the given name already exists, it will be opened. Otherwise, a new collection will be created.

Example:

- To create or access a notes collection namen "my_notes", use:

Example here

Once the tool is running, user will be presented with a menu taht includes options to:
1. Show all notes in the specified collection
2. Add a new note.
3. Delete an existing note.
4. Exit the application. This option also saves any changes made to the notes.

## Data storage
The Notes Tool stores notes in plain text files with the .txt extension. Each collection of notes corresponds to a separate text file named after the collection (e.g., my_notes.txt). Notes are stored line by line within these files, allowing for easy reading and editing. When notes are added or removed, the text file is updated to reflect the current state of the notes.

