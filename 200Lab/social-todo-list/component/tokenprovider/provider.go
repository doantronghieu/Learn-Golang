package tokenprovider

import (
	"errors"

	"social-todo-list/common"
)

// Interface for generating and validating tokens.
type Token interface {
	GetToken() string
}

// Interface for extracting user-specific data from tokens.
type TokenPayload interface {
	UserId() int
	Role() string
}

// Interface that defines methods for token generation, validation, and secret key retrieval.
type Provider interface {
	// creates a new token using the provided data and expiration time.
	Generate(data TokenPayload, expiry int) (Token, error)

	// checks the validity of a given token and returns the associated payload.
	Validate(token string) (TokenPayload, error)

	// returns the secret key used for token generation and validation.
	SecretKey() string
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)

	ErrEncodingToken = common.NewCustomError(
		errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = common.NewCustomError(
		errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
	)
)
