package custom

import (
	"errors"
)

var IndexError = errors.New("\033[31mERROR: Index out of range\033[0m")

var IDError = errors.New("\033[31mERROR: No ID matched\033[0m")

var TitleError = errors.New("\033[31mERROR: Invalid Title\033[0m")

var EmptyNotesErr = errors.New("\033[33mNo Notes Created.\033[0m")