package model

import (
	"errors"
	"strings"

	"social-todo-list/common"
)

const (
	EntityName = "Item"
)

var (
	ErrTitleCannotBeEmpty = errors.New("title can not be blank")
	ErrItemDeleted        = errors.New("item is deleted")
)

type TodoItem struct {
	common.SQLModel
	UserId      int           `json:"user_id" gorm:"column:user_id"`
	Title       string        `json:"title" gorm:"column:title"`
	Description string        `json:"description" gorm:"column:description"`
	Status      *ItemStatus   `json:"status" gorm:"column:status"`
	Image       *common.Image `json:"image" gorm:"column:image"`
}

func (TodoItem) TableName() string { return "todo_items" }

type TodoItemCreation struct {
	Id          int           `json:"-" gorm:"column:id"`
	UserId      int           `json:"-" gorm:"column:user_id"`
	Title       string        `json:"title" gorm:"column:title"`
	Description string        `json:"description" gorm:"column:description"`
	Status      *ItemStatus   `json:"status" gorm:"column:status"`
	Image       *common.Image `json:"image" gorm:"column:image"`
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

func (i *TodoItemCreation) Validate() error {
	i.Title = strings.TrimSpace(i.Title)

	if i.Title == "" {
		return ErrTitleCannotBeEmpty
	}

	return nil
}

type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *string `json:"status" gorm:"column:status"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }
