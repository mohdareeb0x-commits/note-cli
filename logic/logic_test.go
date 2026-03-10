package logic

import (
	"os"
	"testing"
	"fmt"
)

func clearDataFile(t *testing.T) {
	homeDir, _ := os.UserHomeDir()

	path := fmt.Sprintf("%s/data.json", homeDir)

	if err := os.WriteFile(path, []byte("[]"), 0644); err != nil {
		t.Fatalf("unable to reset data file: %v", err)
	}
}

func TestAddNote(t *testing.T) {
	clearDataFile(t)

	noteTitle := "Test Title"
	noteDescription := "Test Description"
	err := AddNote(noteTitle, noteDescription)

	if err != nil {
		t.Fatalf("AddNote returned unexpected error: %v", err)
	}

	notes := ReadData()

	if len(notes) == 0 {
		t.Errorf("expected at least one note after AddNote")

	} else if notes[0].Title != noteTitle || notes[0].Description != noteDescription {
		t.Errorf("Expected Title: %s Expected Description: %s Got Title: %s Description: %s", noteTitle, noteDescription, notes[0].Title, notes[0].Description)
	}
}

func TestDeleteByID(t *testing.T) {
	clearDataFile(t)

	noteTitle := "for delete"
	noteDesc := "will be removed"
	if err := AddNote(noteTitle, noteDesc); err != nil {
		t.Fatalf("setup AddNote failed: %v", err)
	}

	id := 1

	err := DeleteByID(id)
	if err != nil {
		t.Fatalf("DeleteByID returned error: %v", err)
	}

	notes := ReadData()
	if len(notes) != 0 {
		t.Errorf("expected file to be empty after DeleteByID, got %d notes", len(notes))
	}
}

func TestDeleteAll(t *testing.T) {

	clearDataFile(t)
	if err := AddNote("foo", "bar"); err != nil {
		t.Fatalf("setup AddNote failed: %v", err)
	}

	DeleteAll()
	notes := ReadData()
	if len(notes) != 0 {
		t.Errorf("expected zero notes after DeleteAll, got %d", len(notes))
	}
}

func TestGetNotes(t *testing.T) {
	clearDataFile(t)
	if err := AddNote("get title", "get desc"); err != nil {
		t.Fatalf("setup AddNote failed: %v", err)
	}

	testID := 1
	err := GetNotes(testID)

	if err != nil {
		t.Fatalf("GetNotes returned unexpected error: %v", err)
	}
}

func TestListNotes(t *testing.T) {
	clearDataFile(t)
	if err := AddNote("list title", "list desc"); err != nil {
		t.Fatalf("setup AddNote failed: %v", err)
	}

	err := ListNotes()
	if err != nil {
		t.Fatalf("ListNotes returned unexpected error: %v", err)
	}
}

func TestUpdate(t *testing.T) {
	clearDataFile(t)
	if err := AddNote("orig title", "orig desc"); err != nil {
		t.Fatalf("setup AddNote failed: %v", err)
	}

	testId := 1
	testTitle := "Updated Test Title"
	testDesc := "Updated Test Description"

	err := Update(testId, testTitle, testDesc)
	if err != nil {
		t.Fatalf("Update returned unexpected error: %v", err)
	}

	notes := ReadData()
	if len(notes) == 0 {
		t.Fatalf("no notes found after update")
	}
	if notes[0].Title != testTitle {
		t.Errorf("Update malfunctioning. Title didn't update.")
	}
	if notes[0].Description != testDesc {
		t.Errorf("Update malfunctioning. Description didn't update.")
	}
}
