package models

type Tasklist struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Priority  string `json:"priority"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	DueDate   string `json:"due_date"`
}
