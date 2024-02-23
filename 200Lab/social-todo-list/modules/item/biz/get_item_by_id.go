package biz

import (
	"context"

	"social-todo-list/modules/item/model"
)

// GetItemStorage defines the interface for fetching items from storage.
type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}

// getItemBiz is a business logic layer for fetching items.
type getItemBiz struct {
	store GetItemStorage
}

// NewGetItemBiz creates a new instance of getItemBiz with the provided storage.
func NewGetItemBiz(store GetItemStorage) *getItemBiz {
	return &getItemBiz{store: store}
}

// GetItemById fetches an item by ID using the stored business logic.
func (biz *getItemBiz) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {
	// Retrieve item data from the storage using the provided ID
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		// Handle error during item retrieval
		return nil, err
	}
	// Return the retrieved item data
	return data, nil
}