package models

type Tasklist struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Priority  string `json:"priority"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	DueDate   string `json:"dueDate"`
}
