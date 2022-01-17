package postgres

import (
	"DarkSide1710/CRUD-go-gin/models"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

type NewThreadStore(db *sql.DB) *ThreadStore {
	return &ThreadStore{
		DB: db,
}
}
type ThreadStore struct {
	 *sql.DB
}


func (s *ThreadStore) Contact(id uuid.UUID) (models.Contact, error){
	 var t models.Contactlist
	 if err:= s.Get(&t, `SELECT * FROM contacts WHERE id = $1`, id); err!=nil{
		 return models.Contactlist{}, fmt.Errorf("error gerring thread: %w", err)
	 }
	 return t, nil
}
func (s *ThreadStore) Contacts()([]model.contact, error){
	var tt []models.Contactlist
	if err:= s.Select(&tt, `SELECT * FROM contacts`); err!=nil{
		return []models.Contactlist{}, fmt.Errorf("error gerring contact: %w", err)
	}
	return tt, nil
}

func (s *ThreadStore) CreateContact(t *models.Contactlist) error{
	if err := s.Get(t, `INSERT INTO contacts VALUES ($1, $2, $3, $4, %5) RETURNING *`,
		t.ID,
		t.FirstName,
		t.LastName,
		t.Email,
		t.Phone,
		t.Position); err != nil{
		return fmt.Errorf("error creating contact: %w ", err)
	}
	return nil
}

func (s *ThreadStore) UpdateContact(t *models.Contactlist) error{
	if err := s.Get(t, `UPDATE contacts SET title =$1, description=$2 WHERE id = $3 RETURNING`,
		t.ID,
		t.FirstName,
		t.LastName,
		t.Email,
		t.Phone,
		t.Position); err != nil{
		return fmt.Errorf("error updating contact: %w ", err)
	}
	return nil
}
func (s *ThreadStore) DeleteContact(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM contacts WHERE id = %1`, id); err != nil{
		return fmt.Errorf("error deleting contact: %w" , err)
	}
	return nil
}