package usecase

import (
	"testing"

	"github.com/sbrunettajr/clean-architecture-api/domain/entity"
	"github.com/sbrunettajr/clean-architecture-api/mock"
	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"
)

func TestItem_AddItem(t *testing.T) {
	testCases := []struct {
		desc     string
		category entity.Category
	}{
		{
			desc: "WithActiveCategory",
			category: entity.Category{
				Active: true,
			},
		},
		{
			desc: "WithInactiveCategory",
			category: entity.Category{
				Active: false,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var (
				categoryRepository = new(mock.CategoryRepo)
				itemRepository     = new(mock.ItemRepo)
			)
			categoryRepository.On("GetCategory", testifymock.Anything).Return(tC.category, nil)
			itemRepository.On("AddItem", testifymock.Anything).Return(nil)

			itemUseCase := NewItemUseCase(categoryRepository, itemRepository)

			err := itemUseCase.AddItem(entity.Item{})
			if tC.category.Active {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
