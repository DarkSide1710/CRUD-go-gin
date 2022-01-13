package services

import (
	"example/web-service-gin/models"
	"example/web-service-gin/repository"
	"github.com/google/uuid"
)

type Services struct {
	Repository *repository.ContactDB
}

func NewServices() *Services {
	return &Services{
		Repository: repository.NewRepository(),
	}
}

func (s *Services) GetAllContacts() ([]models.Contactlist, error) {
	return s.Repository.GetAllContacts()
}

func (s *Services) GetContact(id string) (models.Contactlist, error) {
	return s.Repository.GetContact(id)
}

func (s *Services) CreateContact(cList models.Contactlist) (models.Contactlist, int, error) {
	cList.ID = uuid.New().String()

	cList, err := s.Repository.CreateContact(cList)
	if err != nil {
		return models.Contactlist{}, 400, err
	}

	return cList, 201, nil
}

func (s *Services) UpdateContact(cList models.Contactlist) (models.Contactlist, error) {
	if err := s.Repository.UpdateContact(cList); err != nil {
		return models.Contactlist{}, err
	}

	return cList, nil
}

func (s *Services) DeleteContact(id string) error {
	if err := s.Repository.DeleteContact(id); err != nil {
		return err
	}

	return nil
}
