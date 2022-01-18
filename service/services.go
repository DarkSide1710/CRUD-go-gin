package services

import "DarkSide1710/CRUD-go-gin/repository"

type Services struct {
	repo repository.Repository

	contact *contactService
	task    *taskService
}

func New(repo *repository.Repositories) *Services {
	return &Services{
		repo: repo,
	}
}

func (s *Services) Contact() Contact {
	if s.contact == nil {
		s.contact = newContactService(s.repo)
	}
	return s.contact
}

func (s *Services) Task() Task {
	if s.task == nil {
		s.task = newTaskService(s.repo)
	}

	return s.task
}
