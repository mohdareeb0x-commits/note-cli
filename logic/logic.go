package logic

import (
	"os"
	"fmt"
	"encoding/json"
	noteModel "note-cli/model"
)

func readData() []noteModel.Note {

	file, err := os.ReadFile("data.json")

	if err != nil {
		fmt.Println("ERROR: Error Fetching Data.\n", err, "\nCreating 'data.json' ...\nDONE!")
	}

	var notes []noteModel.Note

	json.Unmarshal(file, &notes)

	return notes
}

func writeData(notes *[]noteModel.Note) {
	data, err := json.Marshal(notes)

	if err != nil {
		fmt.Println("ERROR: Error Parsing Data to JSON.")
	}

	os.WriteFile("data.json", data, 0644)
}

func AddNote(noteTitle string, noteDescription string) {
	notes := readData()

	note := []noteModel.Note{
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

func UpdateDesc(id int, newDesc string) {

	notes := readData()
	
	for i, note := range notes {
		if note.ID == id {
			notes[i].Description = newDesc
			break
		}
	}
	
	writeData(&notes)

}

func UpdateTitle(id int, newTitle string) {

	notes := readData()
	
	for i, note := range notes {
		if note.ID == id {
			notes[i].Title = newTitle
			break
		}
	}
	
	writeData(&notes)

}