package services

import (
	"errors"

	"github.com/Safiramdhn/project-app-inventory-golang-safira/models"
	"github.com/Safiramdhn/project-app-inventory-golang-safira/repositories"
)

type ItemService struct {
	ItemRepo repositories.ItemRepositoryDB
}

func NewItemService(itemRepo repositories.ItemRepositoryDB) *ItemService {
	return &ItemService{ItemRepo: itemRepo}
}

func (s *ItemService) GetItems() ([]models.Item, error) {
	return s.ItemRepo.GetAll()
}

func (s *ItemService) GetItemByID(id int) (*models.Item, error) {
	if id <= 0 {
		return nil, nil
	}
	return s.ItemRepo.GetByID(id)
}

func (s *ItemService) AddItem(item models.Item) error {
	if item.Name == "" || item.Quantity == 0 || item.Price == 0 {
		return nil
	}
	return s.ItemRepo.Add(&item)
}

func (s *ItemService) UpdateItem(item models.Item) error {
	if item.ID <= 0 {
		return errors.New("invalid item ID")
	}
	return s.ItemRepo.Update(&item)
}

func (s *ItemService) DeleteItem(id int) error {
	if id <= 0 {
		return nil
	}
	return s.ItemRepo.Delete(id)
}

func (s *ItemService) SearchItems(item models.Item) ([]models.Item, error) {
	return s.ItemRepo.Search(item)
}
