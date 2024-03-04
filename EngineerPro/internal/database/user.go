package database

import "context"

type User struct {
	ID   uint64
	Name string
}

type UserDataAccessor interface {
	GetUser(ctx context.Context, id uint64) (User, error)
	CreateUser(ctx context.Context, user User) error
	UpdateUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, id uint64) error
}

type mockUserDataAccessor struct {
}

// CreateUser implements UserDataAccessor.
func (m *mockUserDataAccessor) CreateUser(ctx context.Context, user User) error {
	panic("unimplemented")
}

// DeleteUser implements UserDataAccessor.
func (m *mockUserDataAccessor) DeleteUser(ctx context.Context, id uint64) error {
	panic("unimplemented")
}

// GetUser implements UserDataAccessor.
func (m *mockUserDataAccessor) GetUser(ctx context.Context, id uint64) (User, error) {
	panic("unimplemented")
}

// UpdateUser implements UserDataAccessor.
func (m *mockUserDataAccessor) UpdateUser(ctx context.Context, user User) error {
	panic("unimplemented")
}

func NewMockUserDataAccessor() UserDataAccessor {
	return &mockUserDataAccessor{}
}
