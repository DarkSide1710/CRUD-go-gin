package repository

import (
	"errors"
	"example/web-service-gin/models"

	_ "example/web-service-gin/models"
)

type Contact struct {
	DB map[string]models.Contactlist
}

func NewRepository() *Contact {
	example := models.Contactlist{
		ID:        "1",
		FirstName: "Example Name",
		LastName:  "Example lastname",
		Phone:     999931999,
		Email:     "admin@mail.ru",
		Position:  "director",
	}

	return &Contact{
		DB: map[string]models.Contactlist{example.ID: example},
	}
}

func (repo *Contact) GetAllContacts() ([]models.Contactlist, error) {
	var cLists []models.Contactlist

	for _, cList := range repo.DB {
		cLists = append(cLists, cList)
	}

	return cLists, nil
}

func (repo *Contact) GetContact(id string) (models.Contactlist, error) {
	if cList, exist := repo.DB[id]; exist {
		return cList, nil
	}

	return models.Contactlist{}, errors.New("not found")
}

func (repo *Contact) CreateContact(cList models.Contactlist) (models.Contactlist, error) {
	if _, exist := repo.DB[cList.ID]; exist {
		return models.Contactlist{}, errors.New("already exist")
	}

	repo.DB[cList.ID] = cList

	return cList, nil
}

func (repo *Contact) UpdateContact(cList models.Contactlist) error {
	if _, exist := repo.DB[cList.ID]; !exist {
		return errors.New("not found")
	}

	repo.DB[cList.ID] = cList

	return nil
}

func (repo *Contact) DeleteContact(id string) error {
	if _, exist := repo.DB[id]; !exist {
		return errors.New("not found")
	}

	delete(repo.DB, id)

	return nil
}
