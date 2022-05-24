package repository

import "github.com/sbrunettajr/clean-architecture-api/domain/entity"

type CategoryRepo interface {
	AddCategory(category entity.Category) error
	GetCategories() ([]entity.Category, error)
	GetCategoryByUUID(categoryUUID string) (entity.Category, error)
}
