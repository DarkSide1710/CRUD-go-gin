package services

import "DarkSide1710/CRUD-go-gin/repository"

type Services struct {
	repo repository.Contact

	contact *contactService
	task    *taskService
}

func New(repo *interface{}) *Services {
	return &Services{
		repo: repo,
	}
}

func (s *Services) Contact() Contact {
	if s.contact == nil {
		s.contact = newContactService()
	}
	return s.contact
}

func (s *Services) Task() Task {
	if s.task == nil {
		s.task = newTaskService()
	}

	return s.task
}
