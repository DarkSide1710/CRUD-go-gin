package repository

import (
	"DarkSide1710/CRUD-go-gin/pkg/models"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type contactDB struct {
	DB *sqlx.DB
}

func newContact(db *sqlx.DB) *contactDB {
	return &contactDB{
		DB: db,
	}
}

func (s *contactDB) GetContact(id uuid.UUID) (models.Contactlist, error) {
	var c models.Contactlist
	if err := s.DB.Get(&c, `SELECT * FROM contacts WHERE id = $1`, id); err != nil {
		return models.Contactlist{}, fmt.Errorf("error gerring contact: %w", err)
	}
	return c, nil
}
func (s *contactDB) GetAllContacts() ([]models.Contactlist, error) {
	var cc []models.Contactlist
	if err := s.DB.Select(&cc, `SELECT * FROM contacts`); err != nil {
		return []models.Contactlist{}, fmt.Errorf("error gerring contact: %w", err)
	}
	return cc, nil
}

func (s *contactDB) CreateContact(t *models.Contactlist) error {
	if err := s.DB.Get(t, `INSERT INTO contacts VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`,
		t.ID,
		t.FirstName,
		t.LastName,
		t.Email,
		t.Phone,
		t.Position); err != nil {
		return fmt.Errorf("error creating contact: %w ", err)
	}
	return nil
}

func (s *contactDB) UpdateContact(t *models.Contactlist) error {
	if err := s.DB.Get(t, `UPDATE contacts SET firstname =$1, lastname=$2, email=$3, phone=$4, position=$5 WHERE id = $6 RETURNING`,
		t.ID,
		t.FirstName,
		t.LastName,
		t.Email,
		t.Phone,
		t.Position); err != nil {
		return fmt.Errorf("error updating contact: %w ", err)
	}
	return nil
}

func (s *contactDB) DeleteContact(id uuid.UUID) error {
	if _, err := s.DB.Exec(`DELETE FROM contacts WHERE id=$1`, id); err != nil {
		return fmt.Errorf("error deleting contact: %w", err)
	}
	return nil
}
