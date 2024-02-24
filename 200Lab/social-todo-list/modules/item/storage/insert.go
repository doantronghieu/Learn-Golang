package storage

import (
	"context"

	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

// Creates a new TodoItem in the SQL data store using the provided data.
// Executes the gorm.DB's Create method with the specified TodoItemCreation data.
// Returns any encountered error during the creation process.
func (s *sqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
