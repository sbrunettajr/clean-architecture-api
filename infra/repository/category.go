package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sbrunettajr/clean-architecture-api/domain/entity"
)

type categoryRepository struct {
	quantity   int
	categories []entity.Category
}

func NewCategoryRepository() *categoryRepository {
	return &categoryRepository{
		categories: make([]entity.Category, 0),
	}
}

func (r *categoryRepository) AddCategory(category entity.Category) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	// Remove this
	r.quantity++

	category.ID = r.quantity
	category.UUID = uuid.String()
	r.categories = append(r.categories, category)
	return nil
}

func (r categoryRepository) GetCategories() ([]entity.Category, error) {
	return r.categories, nil
}

func (r categoryRepository) GetCategoryByUUID(categoryUUID string) (entity.Category, error) {
	for _, category := range r.categories {
		if category.UUID == categoryUUID {
			return category, nil
		}
	}
	return entity.Category{}, errors.New("category not found")
}
