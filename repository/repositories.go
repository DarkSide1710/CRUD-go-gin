package repository

import models "DarkSide1710/CRUD-go-gin/models"

type Repositories struct {
	model models.Contactlist

	contactDB *contactDB
	taskDB    *taskDB
}

func New() *Repositories {
	return &Repositories{
		model: model,
	}
}

func (s *Repositories) Contact() Contact {
	if s.contactDB == nil {
		s.contactDB = NewRepository()
	}
	return s.contactDB
}

func (s *Repositories) Task() Task {
	if s.taskDB == nil {
		s.taskDB = NewRepositoryTask()
	}
	return s.taskDB
}
