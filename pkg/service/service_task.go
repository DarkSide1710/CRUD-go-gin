package services

import (
	"DarkSide1710/CRUD-go-gin/pkg/models"
	"DarkSide1710/CRUD-go-gin/pkg/repository"
	"github.com/google/uuid"
)

type taskService struct {
	Repository repository.Repository
}

func newTaskService(repo repository.Repository) *taskService {
	return &taskService{
		Repository: repo,
	}
}

func (s taskService) GetAllTasks() ([]models.Tasklist, error) {
	return s.Repository.Task().GetAllTasks()
}

func (s taskService) GetTask(id string) (models.Tasklist, error) {
	return s.Repository.Task().GetTask(uuid.MustParse(id))
}

func (s taskService) CreateTask(tList models.Tasklist) (models.Tasklist, int, error) {
	err := s.Repository.Task().CreateTask(&tList)
	if err != nil {
		return models.Tasklist{}, 400, err
	}

	return tList, 201, nil
}

func (s taskService) UpdateTask(tList models.Tasklist) (models.Tasklist, error) {
	if err := s.Repository.Task().UpdateTask(&tList); err != nil {
		return models.Tasklist{}, err
	}

	return tList, nil
}

func (s taskService) DeleteTask(id string) error {
	if err := s.Repository.Task().DeleteTask(uuid.MustParse(id)); err != nil {
		return err
	}

	return nil
}
