package repository

import (
	"DarkSide1710/CRUD-go-gin/models"
	"github.com/google/uuid"
)

type Repository interface {
	Contact() Contact
	Task() Task
}

type Contact interface {
	GetAllContacts() ([]models.Contactlist, error)
	GetContact(id uuid.UUID) (models.Contactlist, error)
	CreateContact(t *models.Contactlist) error
	UpdateContact(t *models.Contactlist) error
	DeleteContact(id uuid.UUID) error
}

type Task interface {
	GetAllTasks() ([]models.Tasklist, error)
	GetTask(id uuid.UUID) (models.Tasklist, error)
	CreateTask(t *models.Tasklist) error
	UpdateTask(t *models.Tasklist) error
	DeleteTask(id uuid.UUID) error
}
