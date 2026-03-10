package testing

import (
	"fmt"
	logic "note-cli/logic"
	"testing"
)

func TestAddNote(t *testing.T) {
	noteTitle := "Test Title"
	noteDescription := "Test Description"
	err := logic.AddNote(noteTitle, noteDescription)

	if err != nil {
		fmt.Println(err, noteTitle)
	}

	notes := logic.ReadData()

	if len(notes) == 0 {
		t.Errorf("Function malfunctioning.")

	} else if notes[0].Title != noteTitle && notes[0].Description != noteDescription {
		t.Errorf("Expected Title: %s Expected Description: %s Got Title: %s Description: %s", noteTitle, noteDescription, notes[0].Title, notes[0].Description)
	}

}
