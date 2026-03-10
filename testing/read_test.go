package testing

import (
	"errors"
	"fmt"
	customErr "note-cli/custom"
	logic "note-cli/logic"
	"testing"
)

func TestReadAll(t *testing.T) {

	err := logic.ReadNotes()

	var emptyErr *customErr.EmptyNotesErr

	if err != nil {

		if errors.As(err, &emptyErr) {
			fmt.Println(err.Error())
			fmt.Println()

		} else {
			t.Errorf("Unexpected Error occured.")
		}
	}

}

func TestReadByID(t *testing.T) {

	testID := 1

	err := logic.ReadNotesID(testID)

	var emptyErr *customErr.EmptyNotesErr

	if err != nil {

		if errors.As(err, &emptyErr) {
			fmt.Println(err.Error())
			fmt.Println()

		} else {
			t.Errorf("Unexpected Error occured.")
		}
	}

}
