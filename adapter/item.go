package adapter

import (
	"github.com/sbrunettajr/clean-architecture-api/domain/entity"
	"github.com/sbrunettajr/clean-architecture-api/viewmodel"
)

func NewItemEntityFromVM(itemVM viewmodel.ItemRequest) entity.Item {
	item := entity.Item{
		Active: itemVM.Active,
		Category: entity.Category{
			UUID: itemVM.CategoryUUID,
		},
		Description: itemVM.Description,
		Name:        itemVM.Name,
		Price:       itemVM.Price,
	}
	return item
}

func NewItemVMFromEntity(item entity.Item) viewmodel.ItemResponse {
	return viewmodel.ItemResponse{
		Active:      item.Active,
		Category:    NewCategoryVMFromEntity(item.Category),
		Description: item.Description,
		Name:        item.Name,
		Price:       item.Price,
	}
}

func NewItemVMListFromEntityList(items []entity.Item) []viewmodel.ItemResponse {
	var itemResponseVMList []viewmodel.ItemResponse
	for _, item := range items {
		itemVM := NewItemVMFromEntity(item)
		itemResponseVMList = append(itemResponseVMList, itemVM)
	}
	return itemResponseVMList
}
