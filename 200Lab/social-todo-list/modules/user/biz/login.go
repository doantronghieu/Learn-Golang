package biz

import (
	"context"

	"social-todo-list/common"
	"social-todo-list/component/tokenprovider"
	"social-todo-list/modules/user/model"
)

type LoginStorage interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*model.User, error)
}

type loginBusiness struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(
	storeUser LoginStorage,
	tokenProvider tokenprovider.Provider,
	hasher Hasher,
	expiry int,
) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

/*
- Find user, email
- Hass input pw, compare with pw in db
- Provider: Issue JWT Access Token and Refresh Token for client
- Return Token(s)
*/
func (business *loginBusiness) Login(
	ctx context.Context, data *model.UserLogin,
) (tokenprovider.Token, error) {
	user, err := business.storeUser.FindUser(
		ctx,
		map[string]interface{}{"email": data.Email},
	)
	if err != nil {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	passHashed := business.hasher.Hash(data.Password + user.Salt)
	if user.Password != passHashed {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	payload := &common.TokenPayload{
		UId:   user.Id,
		URole: user.Role.String(),
	}
	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
