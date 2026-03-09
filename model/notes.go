package app

type Note struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
