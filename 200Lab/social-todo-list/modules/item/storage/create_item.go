package storage

import (
	"context"

	"social-todo-list/modules/item/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
