package structs

type Todo struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Start       string `json:"start"`
	Finish      string `json:"finish"`
	IsDone      bool   `json:"isDone"`
}
