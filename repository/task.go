package repository

import (
	"errors"
	models "example/web-service-gin/models"

	_ "example/web-service-gin/models"
)

type Task struct {
	DB map[string]models.Tasklist
}

func NewRepositoryTask() *Task {
	example := models.Tasklist{
		ID:        "1",
		Name:      "Example Tittle",
		Status:    "123",
		Priority:  "324",
		CreatedAt: "sad",
		CreatedBy: "123",
		DueDate:   "324",
	}

	return &Task{
		DB: map[string]models.Tasklist{example.ID: example},
	}
}

func (repo *Task) GetAllTasks() ([]models.Tasklist, error) {
	var tLists []models.Tasklist

	for _, tList := range repo.DB {
		tLists = append(tLists, tList)
	}
	return tLists, nil

}

func (repo *Task) GetTask(id string) (models.Tasklist, error) {
	if tList, exist := repo.DB[id]; exist {
		return tList, nil
	}

	return models.Tasklist{}, errors.New("not found")
}

func (repo *Task) CreateTask(tList models.Tasklist) (models.Tasklist, error) {
	if _, exist := repo.DB[tList.ID]; exist {
		return models.Tasklist{}, errors.New("already exist")
	}

	repo.DB[tList.ID] = tList

	return tList, nil
}

func (repo *Task) UpdateTask(tList models.Tasklist) error {
	if _, exist := repo.DB[tList.ID]; !exist {
		return errors.New("not found")
	}

	repo.DB[tList.ID] = tList

	return nil
}

func (repo *Task) DeleteTask(id string) error {
	if _, exist := repo.DB[id]; !exist {
		return errors.New("not found")
	}

	delete(repo.DB, id)

	return nil
}
