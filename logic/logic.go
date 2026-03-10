package logic

import (
	"encoding/json"
	"fmt"
	customErr "note-cli/custom"
	noteModel "note-cli/model"
	"os"
)

func ReadData() []noteModel.Note {

	file, err := os.ReadFile("data.json")

	if err != nil {
		fmt.Println("\033[31mERROR: Error Fetching Data.\033[0m\n", err, "\nCreating 'data.json' ...\nDONE!")
	}

	var notes []noteModel.Note

	json.Unmarshal(file, &notes)

	return notes
}

func WriteData(notes *[]noteModel.Note) {
	data, err := json.Marshal(notes)

	if err != nil {
		fmt.Println("\033[31mERROR: Error Parsing Data to JSON.\033[0m")
	}

	os.WriteFile("data.json", data, 0644)
}

func AddNote(noteTitle string, noteDescription string) error {

	if noteTitle == "" {
		return &customErr.TitleError{}
	}

	notes := ReadData()

	var id int

	if len(notes) == 0 {
		id = 1
	} else {
		id = notes[len(notes)-1].ID + 1
	}

	note := []noteModel.Note{
		{ID: id,
			Title:       noteTitle,
			Description: noteDescription},
	}

	notes = append(notes, note...)

	WriteData(&notes)

	return nil

}

func DeleteAll() {
	os.WriteFile("data.json", []byte("[]"), 0644)
}

func DeleteByID(id int) error {

	notes := ReadData()

	if len(notes) == 0 {
		return &customErr.EmptyNotesErr{}
	}

	for _, note := range notes {
		fmt.Println(note)

		if note.ID == id {
			break
		} else {
			return &customErr.IDError{}
		}
	}

	idx := id - 1

	notes = append(notes[:idx], notes[idx+1:]...)

	WriteData(&notes)

	return nil
}

func ReadNotesID(id int) error {

	notes := ReadData()

	for _, note := range notes {

		if note.ID == id {
			requiredNote := fmt.Sprintf("ID: %d \nTitle: '%s' \nDescription: '%s'\n", note.ID, note.Title, note.Description)
			fmt.Println(requiredNote)
			break

		} else {
			return &customErr.IDError{}
		}
	}

	return nil

}

func ReadNotes() error {

	notes := ReadData()


	if len(notes) == 0 {
		return &customErr.EmptyNotesErr{}
	}

	for _, note := range notes {
		requiredNote := fmt.Sprintf("ID: %d \nTitle: '%s' \nDescription: '%s'\n", note.ID, note.Title, note.Description)
		fmt.Println(requiredNote)
	}

	return nil

}

func UpdateDesc(id int, newDesc string) error {

	notes := ReadData()

	if len(notes) == 0 {
		return &customErr.EmptyNotesErr{}
	} 

	for i, note := range notes {

		if note.ID == id {
			notes[i].Description = newDesc
			break

		} else {
			return &customErr.IDError{}
		}
	}

	WriteData(&notes)

	return nil

}

func UpdateTitle(id int, newTitle string) error {

	
	if newTitle == "" {
		return &customErr.TitleError{}
	}
	
	notes := ReadData()
	
	if len(notes) == 0 {
		return &customErr.EmptyNotesErr{}
	} 

	for i, note := range notes {
		if note.ID == id {
			notes[i].Title = newTitle
			break
		} else {
			return &customErr.IDError{}
		}
	}

	WriteData(&notes)

	return nil

}
