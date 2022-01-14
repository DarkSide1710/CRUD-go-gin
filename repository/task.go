package repository

import (
	"errors"

	"DarkSide1710/CRUD-go-gin/models"
)

type taskDB struct {
	DB map[string]models.Tasklist
}

func NewRepositoryTask() *taskDB {
	example := models.Tasklist{
		ID:        "1",
		Name:      "Example Tittle",
		Status:    "123",
		Priority:  "324",
		CreatedAt: "sad",
		CreatedBy: "123",
		DueDate:   "324",
	}

	return &taskDB{
		DB: map[string]models.Tasklist{example.ID: example},
	}
}

func (repo taskDB) GetAllTasks() ([]models.Tasklist, error) {
	var tLists []models.Tasklist

	for _, tList := range repo.DB {
		tLists = append(tLists, tList)
	}
	return tLists, nil

}

func (repo taskDB) GetTask(id string) (models.Tasklist, error) {
	if tList, exist := repo.DB[id]; exist {
		return tList, nil
	}

	return models.Tasklist{}, errors.New("not found")
}

func (repo taskDB) CreateTask(tList models.Tasklist) (models.Tasklist, error) {
	if _, exist := repo.DB[tList.ID]; exist {
		return models.Tasklist{}, errors.New("already exist")
	}

	repo.DB[tList.ID] = tList

	return tList, nil
}

func (repo taskDB) UpdateTask(tList models.Tasklist) error {
	if _, exist := repo.DB[tList.ID]; !exist {
		return errors.New("not found")
	}

	repo.DB[tList.ID] = tList

	return nil
}

func (repo taskDB) DeleteTask(id string) error {
	if _, exist := repo.DB[id]; !exist {
		return errors.New("not found")
	}

	delete(repo.DB, id)

	return nil
}
