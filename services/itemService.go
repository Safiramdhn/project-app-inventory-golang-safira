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

func (s *ItemService) GetItems(pagination models.Pagination) ([]models.Item, error) {
	limit := pagination.PerPage
	offset := (pagination.Page - 1) * pagination.PerPage
	return s.ItemRepo.GetAll(limit, offset)
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

func (s *ItemService) GetAllItemsWithFilter(item models.Item) ([]models.Item, error) {
	limit := item.Pagination.PerPage
	offset := (item.Pagination.Page - 1) * item.Pagination.PerPage
	return s.ItemRepo.GetAllWithFilter(item, limit, offset)
}
