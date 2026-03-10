package testing

import (
	"fmt"
	logic "note-cli/logic"
	"testing"
)

func TestUpdateTitle(t *testing.T) {

	testId := 1
	testTitle := "Updated Test Title"
	
	err := logic.UpdateTitle(testId, testTitle)
	notes := logic.ReadData()
	
	if err != nil {
		fmt.Println(err, testId)
		fmt.Println()
	} else if notes[0].Title != testTitle {
		t.Errorf("Update malfunctioning. Title didn't update.")
	}

}

func TestUpdateDesc(t *testing.T) {

	testId := 1
	testDesc := "Updated Test Description"
	
	err := logic.UpdateDesc(testId, testDesc)
	notes := logic.ReadData()
	
	if err != nil {
		fmt.Println(err, testId)
		fmt.Println()
	} else if notes[0].Description != testDesc {
		t.Errorf("Update malfunctioning. Title didn't update.")
	}

}
