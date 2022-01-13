package services

import (
	"example/web-service-gin/models"
	"example/web-service-gin/repository"
	"github.com/google/uuid"
)

type ServicesT struct {
	Repository *repository.Task
}

func NewservicesT() *ServicesT {
	return &ServicesT{
		Repository: repository.NewRepositoryTask(),
	}
}

func (s *ServicesT) GetAllTasks() ([]models.Tasklist, error) {
	return s.Repository.GetAllTasks()
}

func (s *ServicesT) GetTask(id string) (models.Tasklist, error) {
	return s.Repository.GetTask(id)
}

func (s *ServicesT) CreateTask(tList models.Tasklist) (models.Tasklist, int, error) {
	tList.ID = uuid.New().String()

	tList, err := s.Repository.CreateTask(tList)
	if err != nil {
		return models.Tasklist{}, 400, err
	}

	return tList, 201, nil
}

func (s *ServicesT) UpdateTask(tList models.Tasklist) (models.Tasklist, error) {
	if err := s.Repository.UpdateTask(tList); err != nil {
		return models.Tasklist{}, err
	}

	return tList, nil
}

func (s *ServicesT) DeleteTask(id string) error {
	if err := s.Repository.DeleteTask(id); err != nil {
		return err
	}

	return nil
}
