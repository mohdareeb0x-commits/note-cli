package logic

import (
	"encoding/json"
	noteApp "note-cli/model"
	"fmt"
	"os"
)

func readData() []noteApp.Note {

	file, err := os.ReadFile("data.json")

	if err != nil {
		fmt.Println("ERROR: Error Fetching Data.\n", err, "\nCreating 'data.json' ...\nDONE!")
	}

	var notes []noteApp.Note

	json.Unmarshal(file, &notes)

	return notes
}

func writeData(notes []noteApp.Note) {
	data, err := json.Marshal(notes)

	if err != nil {
		fmt.Println("ERROR: Error Parsing Data to JSON.")
	}

	os.WriteFile("data.json", data, 0644)
}

func AddNote(noteTitle string, noteDescription string) {

	notes := readData()

	note := []noteApp.Note{
		{ID: uint(len(notes)) + 1,
			Title:       noteTitle,
			Description: noteDescription},
	}

	notes = append(notes, note...)

	writeData(notes)

}

func DeleteAll() {
	os.WriteFile("data.json", []byte("[]"), 0644)
}