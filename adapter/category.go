package adapter

import (
	"github.com/sbrunettajr/clean-architecture-api/domain/entity"
	"github.com/sbrunettajr/clean-architecture-api/viewmodel"
)

func NewCategoryEntityFromVM(categoryVM viewmodel.CategoryRequestVM) entity.Category {
	return entity.Category{
		Name:   categoryVM.Name,
		Active: categoryVM.Active,
	}
}

func NewCategoryVMFromEntity(category entity.Category) viewmodel.CategoryResponseVM {
	return viewmodel.CategoryResponseVM{
		UUID:   category.UUID,
		Name:   category.Name,
		Active: category.Active,
	}
}

func NewCategoryVMListFromEntityList(categoryList []entity.Category) []viewmodel.CategoryResponseVM {
	var categoryResponseVMList []viewmodel.CategoryResponseVM
	for _, category := range categoryList {
		categoryVM := NewCategoryVMFromEntity(category)
		categoryResponseVMList = append(categoryResponseVMList, categoryVM)
	}
	return categoryResponseVMList
}
