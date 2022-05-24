package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sbrunettajr/clean-architecture-api/adapter"
	"github.com/sbrunettajr/clean-architecture-api/domain/usecase"
	"github.com/sbrunettajr/clean-architecture-api/viewmodel"
)

type ItemController struct {
	itemUseCase usecase.ItemUseCase
}

func NewItemController(itemUseCase usecase.ItemUseCase) ItemController {
	return ItemController{
		itemUseCase: itemUseCase,
	}
}

func (c ItemController) AddItem(ctx echo.Context) error {
	var itemRequestVM viewmodel.ItemRequest

	err := ctx.Bind(&itemRequestVM)
	if err != nil {
		return err
	}

	return c.itemUseCase.AddItem(
		adapter.NewItemEntityFromVM(itemRequestVM),
	)
}

func (c ItemController) GetItems(ctx echo.Context) error {
	items, err := c.itemUseCase.GetItems()
	if err != nil {
		return err
	}

	itemResponseVMList := adapter.NewItemVMListFromEntityList(items)
	return ctx.JSON(http.StatusOK, itemResponseVMList)
}
