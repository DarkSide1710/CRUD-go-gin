package services

import "DarkSide1710/CRUD-go-gin/pkg/models"

type Container interface {
	Contact() Contact
	Task() Task
}

type Contact interface {
	GetAllContacts() ([]models.Contactlist, error)
	GetContact(id string) (models.Contactlist, error)
	CreateContact(cList models.Contactlist) (models.Contactlist, int, error)
	UpdateContact(cList models.Contactlist) (models.Contactlist, error)
	DeleteContact(id string) error
}

type Task interface {
	GetAllTasks() ([]models.Tasklist, error)
	GetTask(id string) (models.Tasklist, error)
	CreateTask(tList models.Tasklist) (models.Tasklist, int, error)
	UpdateTask(tList models.Tasklist) (models.Tasklist, error)
	DeleteTask(id string) error
}
