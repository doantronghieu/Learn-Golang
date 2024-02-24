package storage

import (
	"context"

	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

func (s *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {

	// Soft delete the TodoItem by updating its status to "Deleted" in the database
	if err := s.db.Table(model.TodoItem{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{"status": "Deleted"}).
		Error; err != nil {

		return common.ErrDB(err)
	}

	return nil
}
