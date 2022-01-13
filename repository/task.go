package repository

import (
	"errors"
	models "example/web-service-gin/models"

	_ "example/web-service-gin/models"
)

type TaskDB struct {
	DB map[string]models.Tasklist
}

func NewRepositoryTask() *TaskDB {
	example := models.Tasklist{
		ID:        "1",
		Name:      "Example Tittle",
		Status:    "123",
		Priority:  "324",
		CreatedAt: "sad",
		CreatedBy: "123",
		DueDate:   "324",
	}

	return &TaskDB{
		DB: map[string]models.Tasklist{example.ID: example},
	}
}

func (repo *TaskDB) GetAllTasks() ([]models.Tasklist, error) {
	var tLists []models.Tasklist

	for _, tList := range repo.DB {
		tLists = append(tLists, tList)
	}
	return tLists, nil

}

func (repo *TaskDB) GetTask(id string) (models.Tasklist, error) {
	if tList, exist := repo.DB[id]; exist {
		return tList, nil
	}

	return models.Tasklist{}, errors.New("not found")
}

func (repo *TaskDB) CreateTask(tList models.Tasklist) (models.Tasklist, error) {
	if _, exist := repo.DB[tList.ID]; exist {
		return models.Tasklist{}, errors.New("already exist")
	}

	repo.DB[tList.ID] = tList

	return tList, nil
}

func (repo *TaskDB) UpdateTask(tList models.Tasklist) error {
	if _, exist := repo.DB[tList.ID]; !exist {
		return errors.New("not found")
	}

	repo.DB[tList.ID] = tList

	return nil
}

func (repo *TaskDB) DeleteTask(id string) error {
	if _, exist := repo.DB[id]; !exist {
		return errors.New("not found")
	}

	delete(repo.DB, id)

	return nil
}
