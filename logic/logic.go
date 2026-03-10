package logic

import (
	"os"
	"fmt"
	"encoding/json"
	noteApp "note-cli/model"
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

func writeData(notes *[]noteApp.Note) {
	data, err := json.Marshal(notes)

	if err != nil {
		fmt.Println("ERROR: Error Parsing Data to JSON.")
	}

	os.WriteFile("data.json", data, 0644)
}

func AddNote(noteTitle string, noteDescription string) {
	notes := readData()

	note := []noteApp.Note{
		{ID: notes[len(notes) - 1].ID + 1,
			Title:       noteTitle,
			Description: noteDescription},
	}

	notes = append(notes, note...)

	writeData(&notes)

}

func DeleteAll() {
	os.WriteFile("data.json", []byte("[]"), 0644)
}

func DeleteByID(id int) {
	notes := readData()
	idx := id - 1

	notes = append(notes[:idx], notes[idx+1:]...)

	writeData(&notes)
}

func ReadNotesID(id int) {

	notes := readData()
	
	for _, note := range notes {
		if note.ID == id {
			requiredNote := fmt.Sprintf("ID: %d \nTitle: '%s' \nDescription: '%s'\n", note.ID, note.Title, note.Description)
			fmt.Println(requiredNote)
			break
		}
	}

}

func ReadNotes() {

	notes := readData()

	for _, note := range notes {
		requiredNote := fmt.Sprintf("ID: %d \nTitle: '%s' \nDescription: '%s'\n", note.ID, note.Title, note.Description)
		fmt.Println(requiredNote)
	}
}