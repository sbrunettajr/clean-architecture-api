package usecase

import (
	"github.com/sbrunettajr/clean-architecture-api/domain/entity"
	"github.com/sbrunettajr/clean-architecture-api/domain/repository"
)

type CategoryUseCase struct {
	categoryRepo repository.CategoryRepo
}

func NewCategoryUseCase(categoryRepo repository.CategoryRepo) CategoryUseCase {
	return CategoryUseCase{
		categoryRepo: categoryRepo,
	}
}

func (u CategoryUseCase) AddCategory(category entity.Category) error {
	return u.categoryRepo.AddCategory(category)
}

func (u CategoryUseCase) GetCategories() ([]entity.Category, error) {
	return u.categoryRepo.GetCategories()
}
