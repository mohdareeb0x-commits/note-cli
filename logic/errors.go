package logic

type IDError struct{}

func (e *IDError) Error() string {
	return "\033[31mERROR: No ID matched\033[0m"
}


type TitleError struct{}

func (e *TitleError) Error() string {
	return "\033[31mERROR: Invalid Title\033[0m"
}


type EmptyNotesErr struct{}

func (e *EmptyNotesErr) Error() string {
	return "\033[33mNo Notes Created.\033[0m"
}