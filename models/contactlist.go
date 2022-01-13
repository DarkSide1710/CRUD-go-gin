package models

type Contactlist struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     int    `json:"phone"`
	Email     string `json:"email"`
	Position  string `json:"position"`
}
