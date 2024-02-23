package storage

import "gorm.io/gorm"

// Layer:
// type TYPE_NAME struct { KEY: (*)VALUE }
// func NewTYPE_NAME(KEY (*)VALUE) *TYPE_NAME
// Layer logic:
// func (t *TYPE_NAME) LOGIC_NAME(ctx context.Context, 
// 																LOGIC_DATA *LOGIC_DATA) error

// SQL data store struct that holds a reference to a gorm.DB instance
type sqlStore struct {
	db *gorm.DB
}

// Constructor function for creating a new sqlStore instance
func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}