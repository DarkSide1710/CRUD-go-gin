package repository

import (
	"fmt"
	"github.com/google/uuid"
	sqlx "github.com/jmoiron/sqlx"

	"DarkSide1710/CRUD-go-gin/pkg/models"
)

type taskDB struct {
	DB *sqlx.DB
}

func newTask(db *sqlx.DB) *taskDB {
	return &taskDB{
		DB: db,
	}
}

func (s *taskDB) GetTask(id uuid.UUID) (models.Tasklist, error) {
	var t models.Tasklist
	if err := s.DB.Get(&t, `SELECT * FROM task WHERE id = $1`, id); err != nil {
		return models.Tasklist{}, fmt.Errorf("error gerring task: %w", err)
	}
	return t, nil
}

func (s *taskDB) GetAllTasks() ([]models.Tasklist, error) {
	var tt []models.Tasklist
	if err := s.DB.Select(&tt, `SELECT * FROM task`); err != nil {
		return []models.Tasklist{}, fmt.Errorf("error gerring task: %w", err)
	}
	return tt, nil
}

func (s *taskDB) CreateTask(t *models.Tasklist) error {
	fmt.Println(t)

	if _, err := s.DB.Exec(`INSERT INTO task (name, status, priority, created_at, created_by, due_date) VALUES ($1, $2, $3, $4, $5, $6)`,
		t.Name,
		t.Status,
		t.Priority,
		t.CreatedAt,
		uuid.MustParse(t.CreatedBy).String(),
		t.DueDate); err != nil {
		return fmt.Errorf("error creating task: %w ", err)
	}
	return nil
}

func (s *taskDB) UpdateTask(t *models.Tasklist) error {
	if err := s.DB.Get(t, `UPDATE task SET name =$1, status=$2, priority=$3, created_at=$4, created_by=$5, due_date=$6 WHERE id = $7 RETURNING`,
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
func (s *taskDB) DeleteTask(id uuid.UUID) error {
	if _, err := s.DB.Exec(`DELETE FROM task WHERE id = %1`, id); err != nil {
		return fmt.Errorf("error deleting task: %w", err)
	}
	return nil
}
