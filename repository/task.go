package repository

import (
	"fmt"
	"github.com/google/uuid"
	sqlx "github.com/jmoiron/sqlx"

	"DarkSide1710/CRUD-go-gin/models"
)

type taskDB struct {
	DB *sqlx.DB
}

func newTask() *taskDB {
	example := models.Tasklist{
		ID:        "1",
		Name:      "Example Tittle",
		Status:    "123",
		Priority:  "324",
		CreatedAt: "sad",
		CreatedBy: "123",
		DueDate:   "324",
	}

	return &taskDB{
		DB: *sqlx.DB{example.ID: example},
	}
}

func (s *contactDB) GetTask(id uuid.UUID) (models.Contactlist, error) {
	var t models.Contactlist
	if err := s.DB.Get(&t, `SELECT * FROM task WHERE id = $1`, id); err != nil {
		return models.Contactlist{}, fmt.Errorf("error gerring task: %w", err)
	}
	return t, nil
}

func (s *contactDB) GetAllTasks() ([]models.Tasklist, error) {
	var tt []models.Contactlist
	if err := s.DB.Select(&tt, `SELECT * FROM task`); err != nil {
		return []models.Tasklist{}, fmt.Errorf("error gerring task: %w", err)
	}
	return tt, nil
}

func (s *contactDB) CreateTask(t *models.Tasklist) error {
	if err := s.DB.Get(t, `INSERT INTO task VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
		t.ID,
		t.Name,
		t.Status,
		t.Priority,
		t.CreatedBy,
		t.CreatedBy,
		t.DueDate); err != nil {
		return fmt.Errorf("error creating task: %w ", err)
	}
	return nil
}

func (s *contactDB) UpdateTask(t *models.Tasklist) error {
	if err := s.DB.Get(t, `UPDATE contacts SET name =$1, status=$2, priority=$3, createat=$4, createby=$5, duedate=$6 WHERE id = $7 RETURNING`,
		t.ID,
		t.Name,
		t.Status,
		t.Priority,
		t.CreatedAt,
		t.CreatedBy,
		t.DueDate); err != nil {
		return fmt.Errorf("error updating task: %w ", err)
	}
	return nil
}
func (s *contactDB) DeleteTask(id uuid.UUID) error {
	if _, err := s.DB.Exec(`DELETE FROM task WHERE id = %1`, id); err != nil {
		return fmt.Errorf("error deleting task: %w", err)
	}
	return nil
}
