package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sbrunettajr/clean-architecture-api/adapter"
	"github.com/sbrunettajr/clean-architecture-api/domain/usecase"
	"github.com/sbrunettajr/clean-architecture-api/viewmodel"
)

type categoryController struct {
	categoryUseCase usecase.CategoryUseCase
}

func NewCategoryController(categoryUseCase usecase.CategoryUseCase) categoryController {
	return categoryController{
		categoryUseCase: categoryUseCase,
	}
}

func (c categoryController) AddCategory(ctx echo.Context) error {
	var categoryVM viewmodel.CategoryRequestVM

	err := ctx.Bind(&categoryVM)
	if err != nil {
		return err
	}

	return c.categoryUseCase.AddCategory(
		adapter.NewCategoryEntityFromVM(categoryVM),
	)
}

func (c categoryController) GetCategories(ctx echo.Context) error {
	categories, err := c.categoryUseCase.GetCategories()
	if err != nil {
		return err
	}

	categoryResponseVMList := adapter.NewCategoryVMListFromEntityList(categories)
	return ctx.JSON(http.StatusOK, categoryResponseVMList)
}
