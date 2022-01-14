package repository

import "DarkSide1710/CRUD-go-gin/models"

type Repository interface {
	Contact() Contact
	Task() Task
}

type Contact interface {
	GetAllContacts() ([]models.Contactlist, error)
	GetContact(id string) (models.Contactlist, error)
	CreateContact(cList models.Contactlist) (models.Contactlist, error)
	UpdateContact(cList models.Contactlist) error
	DeleteContact(id string) error
}

type Task interface {
	GetAllTasks() ([]models.Tasklist, error)
	GetTask(id string) (models.Tasklist, error)
	CreateTask(tList models.Tasklist) (models.Tasklist, error)
	UpdateTask(tList models.Tasklist) error
	DeleteTask(id string) error
}
