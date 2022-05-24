package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sbrunettajr/clean-architecture-api/controller"
	"github.com/sbrunettajr/clean-architecture-api/domain/usecase"
	"github.com/sbrunettajr/clean-architecture-api/infra/repository"
)

func main() {
	categoryRepository := repository.NewCategoryRepository()
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
	categoryController := controller.NewCategoryController(categoryUseCase)

	itemRepository := repository.NewItemRepository()
	itemUseCase := usecase.NewItemUseCase(categoryRepository, itemRepository)
	itemController := controller.NewItemController(itemUseCase)

	server := echo.New()

	server.GET("/items", itemController.GetItems)
	server.POST("/items", itemController.AddItem)

	server.GET("/categories", categoryController.GetCategories)
	server.POST("/categories", categoryController.AddCategory)

	server.Start(":5000")
}
