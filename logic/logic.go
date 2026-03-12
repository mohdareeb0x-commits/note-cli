package logic

import (
	"encoding/json"
	"fmt"
	"os"
)

type Note struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func AddNote(noteTitle string, noteDescription string) error {

	if noteTitle == "" {
		return &TitleError{}
	}

	notes := ReadData()

	var id int

	if len(notes) == 0 {
		id = 1
	} else {
		id = notes[len(notes)-1].ID + 1
	}

	note := []Note{
		{ID: id,
			Title:       noteTitle,
			Description: noteDescription},
	}

	notes = append(notes, note...)

	WriteData(&notes)

	return nil

}

func DeleteAll() {
	homeDir, _ := os.UserHomeDir()

	path := fmt.Sprintf("%s/data.json", homeDir)

	os.WriteFile(path, []byte("[]"), 0644)
}

func DeleteByID(id int) error {

	notes := ReadData()

	if len(notes) == 0 {
		return &EmptyNotesErr{}
	}

	var idMatched bool

	for _, note := range notes {

		if note.ID == id {
			idMatched = true
			break
		} else {
			idMatched = false
		}
	}

	if !idMatched {
		return &IDError{}
	}

	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			break
		}
	}

	WriteData(&notes)

	return nil
}

func GetNotes(id int) error {

	notes := ReadData()

	if len(notes) == 0 {
		return &EmptyNotesErr{}
	}

	var idMatched bool

	for _, note := range notes {

		if note.ID == id {
			requiredNote := fmt.Sprintf("ID: %d \nTitle: '%s' \nDescription: '%s'\n", note.ID, note.Title, note.Description)
			fmt.Println(requiredNote)
			idMatched = true
			break

		} else {
			idMatched = false
		}
	}

	if !idMatched {
		return &IDError{}

	}

	return nil

}

func ListNotes() error {

	notes := ReadData()

	if len(notes) == 0 {
		return &EmptyNotesErr{}
	}

	for _, note := range notes {
		requiredNote := fmt.Sprintf("ID: %d \nTitle: '%s' \nDescription: '%s'\n", note.ID, note.Title, note.Description)
		fmt.Println(requiredNote)
	}

	return nil

}

func ReadData() []Note {

	homeDir, _ := os.UserHomeDir()

	path := fmt.Sprintf("%s/data.json", homeDir)

	file, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("\033[31mERROR: Error Fetching Data.\033[0m\n", err, "\nCreating 'data.json' ...\nDONE!")
	}

	var notes []Note

	json.Unmarshal(file, &notes)

	return notes
}

func Update(id int, newTitle string, newDesc string) error {

	notes := ReadData()

	if len(notes) == 0 {
		return &EmptyNotesErr{}
	}

	var idMatched bool

	for i, note := range notes {

		if note.ID == id {
			if newDesc == "" && newTitle != "" {
				notes[i].Title = newTitle

			} else if newTitle == "" && newDesc != "" {
				notes[i].Description = newDesc

			} else {
				notes[i].Title = newTitle
				notes[i].Description = newDesc
			}
			idMatched = true
			break

		} else {
			idMatched = false
		}
	}

	if !idMatched {
		return &IDError{}
	}

	WriteData(&notes)

	return nil
}

func WriteData(notes *[]Note) {

	homeDir, _ := os.UserHomeDir()

	path := fmt.Sprintf("%s/data.json", homeDir)

	data, err := json.Marshal(notes)

	if err != nil {
		fmt.Println("\033[31mERROR: Error Parsing Data to JSON.\033[0m")
	}

	os.WriteFile(path, data, 0644)
}
