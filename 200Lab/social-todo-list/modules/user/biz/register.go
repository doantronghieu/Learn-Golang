package biz

import (
	"context"

	"social-todo-list/common"
	"social-todo-list/modules/user/model"
)

type RegisterStorage interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*model.User, error)

	CreateUser(ctx context.Context, data *model.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		registerStorage: registerStorage,
		hasher:          hasher,
	}
}

func (business *registerBusiness) Register(
	ctx context.Context,
	data *model.UserCreate,
) error {
	user, _ := business.registerStorage.FindUser(
		ctx, map[string]interface{}{"email": data.Email},
	)

	if user != nil {
		return model.ErrEmailExisted
	}

	salt, _ := common.GenerateSalt(10)

	data.Password = business.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = model.RoleUser

	if err := business.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}
