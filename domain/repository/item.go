package repository

import "github.com/sbrunettajr/clean-architecture-api/domain/entity"

type ItemRepo interface {
	AddItem(entity.Item) error
	GetItems() ([]entity.Item, error)
}
