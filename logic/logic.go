package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Note struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Note addition logic
func AddNote(noteTitle string, noteDescription string) error {

	if noteTitle == "" {
		return &TitleError{}
	}

	notes, err := ReadData()

	if err != nil {
		return err
	}

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

	err = WriteData(notes)

	if err != nil {
		return err
	}

	return nil

}

// Note deletion logic
func DeleteAll() error {

	notes, err := ReadData()

	if err != nil {
		return err
	}

	if len(notes) == 0 {
		return &EmptyNotesErr{}
	}

	path, err := GetPath()

	if err != nil {
		return err
	}

	os.WriteFile(path, []byte("[]"), 0644)

	return nil
}

// Note deletion by ID logic
func DeleteByID(id int) error {

	notes, err := ReadData()

	if err != nil {
		return err
	}

	if len(notes) == 0 {
		return &EmptyNotesErr{}
	}

	var idMatched bool

	for i, note := range notes {
		if note.ID == id {
			idMatched = true
			notes = append(notes[:i], notes[i+1:]...)
			break
		} else {
			idMatched = false
		}
	}

	if !idMatched {
		return &IDError{}
	}

	err = WriteData(notes)

	if err != nil {
		return err
	}

	return nil
}

// Get notes logic
func GetNotes(id int) (string, error) {

	notes, err := ReadData()

	if err != nil {
		return "", err
	}

	if len(notes) == 0 {
		return "", &EmptyNotesErr{}
	}

	var idMatched bool
	var requiredNote string

	for _, note := range notes {

		if note.ID == id {
			requiredNote = fmt.Sprintf("ID: %d \nTitle: '%s' \nDescription: '%s'\n", note.ID, note.Title, note.Description)

			idMatched = true
			break

		} else {
			idMatched = false
		}
	}

	if !idMatched {
		return "", &IDError{}

	}

	return requiredNote, nil

}

// Get path logic
func GetPath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", errors.New("\033[31mERROR: Unable to get Home directory\033[0m")
	}

	dir := filepath.Join(homeDir, "note-cli")

	_, err = os.Stat(dir)

	if os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return "", err
		}
	}

	path := fmt.Sprintf("%s/data.json", dir)

	return path, nil
}

// List notes logic
func ListNotes() (string, error) {

	notes, err := ReadData()

	if err != nil {
		return "", err
	}

	if len(notes) == 0 {
		return "", &EmptyNotesErr{}
	}

	var noteList string

	for _, note := range notes {
		noteList += fmt.Sprintf("\nID: %d \nTitle: '%s' \nDescription: '%s'\n", note.ID, note.Title, note.Description)

	}

	return noteList, nil

}

// Data reading logic
func ReadData() ([]Note, error) {

	path, err := GetPath()

	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(path)

	if err != nil {
		_, ierr := os.Create(path)

		if ierr != nil {
			return nil, err
		}
		newErr := fmt.Sprintf("\n\033[31mERROR: Error Fetching Data.\033[0m: %s \nCreating 'data.json'.....\nDONE!\n\033[33mAdd Note again.\n\033[0m", err.Error())

		return nil, errors.New(newErr)
	}

	var notes []Note

	err = json.Unmarshal(file, &notes)

	if err != nil {
		return nil, err
	}

	return notes, nil
}

// Note update logic
func Update(id int, newTitle string, newDesc string) error {

	notes, err := ReadData()

	if err != nil {
		return err
	}

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

	err = WriteData(notes)

	if err != nil {
		return err
	}

	return nil
}

// Data writing logic
func WriteData(notes []Note) error {

	data, err := json.MarshalIndent(notes, "", " ")

	if err != nil {
		return errors.New("\033[31mERROR: Error Parsing Data to JSON.\033[0m")
	}

	path, err := GetPath()

	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)

	if err != nil {
		return err
	}

	return nil
}
