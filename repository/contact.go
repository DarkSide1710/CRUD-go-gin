package repository

import (
	"errors"

	"DarkSide1710/CRUD-go-gin/models"
)

type contactDB struct {
	DB map[string]models.Contactlist
}

func NewRepository() *contactDB {
	example := models.Contactlist{
		ID:        "1",
		FirstName: "Example Name",
		LastName:  "Example lastname",
		Phone:     "998999990101",
		Email:     "admin@mail.ru",
		Position:  "director",
	}

	return &contactDB{
		DB: map[string]models.Contactlist{example.ID: example},
	}
}

func (repo contactDB) GetAllContacts() ([]models.Contactlist, error) {
	var cLists []models.Contactlist

	for _, cList := range repo.DB {
		cLists = append(cLists, cList)
	}

	return cLists, nil
}

func (repo contactDB) GetContact(id string) (models.Contactlist, error) {
	if cList, exist := repo.DB[id]; exist {
		return cList, nil
	}

	return models.Contactlist{}, errors.New("not found")
}

func (repo contactDB) CreateContact(cList models.Contactlist) (models.Contactlist, error) {
	if _, exist := repo.DB[cList.ID]; exist {
		return models.Contactlist{}, errors.New("already exist")
	}

	repo.DB[cList.ID] = cList

	return cList, nil
}

func (repo contactDB) UpdateContact(cList models.Contactlist) error {
	if _, exist := repo.DB[cList.ID]; !exist {
		return errors.New("not found")
	}

	repo.DB[cList.ID] = cList

	return nil
}

func (repo contactDB) DeleteContact(id string) error {
	if _, exist := repo.DB[id]; !exist {
		return errors.New("not found")
	}

	delete(repo.DB, id)

	return nil
}
