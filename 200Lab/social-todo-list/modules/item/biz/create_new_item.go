package biz

import (
	"context"

	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

// Client
// 		-> Handler/Transport: Process/Parse <JSON> request => Biz desired structure
// 				-> Biz: Logic
// 						[-> Respository]: Aggregate information from storage(s)
// 								-> Storage (MySQL, MongoDB ...)

// Layers communicate via interface. Use interface in layer that need to use

// Layer:
// type TYPE_NAME struct { KEY: (*)VALUE }
// func NewTYPE_NAME(KEY (*)VALUE) *TYPE_NAME
// Layer logic:
// func (t *TYPE_NAME) LOGIC_NAME(ctx context.Context,
// 																LOGIC_DATA *LOGIC_DATA) error

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}

type createItemBiz struct {
	store CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz {
	return &createItemBiz{store: store}
}

func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := data.Validate(); err != nil {
		return model.ErrTitleCannotBeEmpty
	}

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}