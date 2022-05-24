package repository

import "github.com/sbrunettajr/clean-architecture-api/domain/entity"

type itemRepository struct {
	items []entity.Item
}

func NewItemRepository() *itemRepository {
	return &itemRepository{
		items: make([]entity.Item, 0),
	}
}

func (r *itemRepository) AddItem(item entity.Item) error {
	r.items = append(r.items, item)
	return nil
}

func (r itemRepository) GetItems() ([]entity.Item, error) {
	return r.items, nil
}
