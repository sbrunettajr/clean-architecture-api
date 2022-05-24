package usecase

import (
	"errors"

	"github.com/sbrunettajr/clean-architecture-api/domain/entity"
	"github.com/sbrunettajr/clean-architecture-api/domain/repository"
)

type ItemUseCase struct {
	categoryRepo repository.CategoryRepo
	itemRepo     repository.ItemRepo
}

func NewItemUseCase(categoryRepository repository.CategoryRepo, itemRepository repository.ItemRepo) ItemUseCase {
	return ItemUseCase{
		categoryRepo: categoryRepository,
		itemRepo:     itemRepository,
	}
}

func (u ItemUseCase) AddItem(item entity.Item) error {
	category, err := u.categoryRepo.GetCategoryByUUID(item.Category.UUID)
	if err != nil {
		return err
	}

	if !category.IsActive() {
		return errors.New("this category is inactive")
	}
	item.Category = category

	return u.itemRepo.AddItem(item)
}

func (u ItemUseCase) GetItems() ([]entity.Item, error) {
	return u.itemRepo.GetItems()
}
