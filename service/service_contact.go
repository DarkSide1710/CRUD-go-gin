package services

import (
	"DarkSide1710/CRUD-go-gin/models"
	"DarkSide1710/CRUD-go-gin/repository"
	"github.com/google/uuid"
)

type contactService struct {
	Repository repository.Repository
}

func newContactService(repo repository.Repository) *contactService {
	return &contactService{
		Repository: repo,
	}
}

func (s contactService) GetAllContacts() ([]models.Contactlist, error) {
	return s.Repository.Contact().GetAllContacts()
}

func (s contactService) GetContact(id string) (models.Contactlist, error) {
	return s.Repository.Contact().GetContact(id)
}

func (s contactService) CreateContact(cList models.Contactlist) (models.Contactlist, int, error) {
	cList.ID = uuid.New().String()

	cList, err := s.Repository.Contact().CreateContact(cList)
	if err != nil {
		return models.Contactlist{}, 400, err
	}

	return cList, 201, nil
}

func (s contactService) UpdateContact(cList models.Contactlist) (models.Contactlist, error) {
	if err := s.Repository.Contact().UpdateContact(cList); err != nil {
		return models.Contactlist{}, err
	}

	return cList, nil
}

func (s contactService) DeleteContact(id string) error {
	if err := s.Repository.Contact().DeleteContact(id); err != nil {
		return err
	}

	return nil
}
