package testing

import (
	"fmt"
	logic "note-cli/logic"
	"testing"
)

func TestDeleteByID(t *testing.T) {

	id := 1

	err := logic.DeleteByID(id)

	if err != nil {
		fmt.Println(err.Error(), id)
		fmt.Println()
	}

	notes := logic.ReadData()

	if len(notes) != 0 {
		t.Errorf("Deletion malfunctioning. Nothing got deleted.")
	}

}

func TestDeleteAll(t *testing.T) {
	logic.DeleteAll()

	notes := logic.ReadData()

	if len(notes) != 0 {
		t.Errorf("Deletion malfunctioning. Nothing got deleted.")
	}

}

