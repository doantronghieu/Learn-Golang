package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"

	"social-todo-list/common"
	"social-todo-list/component/tokenprovider"
)

// struct representing a provider for generating and validating JWT tokens.
type jwtProvider struct {
	prefix string
	secret string
}

// creates a new instance of jwtProvider with the specified prefix.
func NewTokenJWTProvider(prefix string) *jwtProvider {
	return &jwtProvider{prefix: prefix}
}

// represents custom claims structure for JWT.
type myClaims struct {
	Payload common.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

// represents the structure of the JWT token response.
type token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

// returns the token string from the token structure.
func (t *token) GetToken() string {
	return t.Token
}

// returns the secret key used for JWT signing.
func (j *jwtProvider) SecretKey() string {
	return j.secret
}

// generates a new JWT token with the provided payload and expiry time.
func (j *jwtProvider) Generate(
	data tokenprovider.TokenPayload,
	expiry int,
) (tokenprovider.Token, error) {
	now := time.Now()

	// Create a new JWT token with custom claims and standard claims.
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		common.TokenPayload{
			UId:   data.UserId(),
			URole: data.Role(),
		},
		jwt.StandardClaims{
			ExpiresAt: now.Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  now.Local().Unix(),
			Id:        fmt.Sprintf("%d", now.UnixNano()),
		},
	})

	// Sign the token with the secret key.
	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	return &token{Token: myToken, Expiry: expiry, Created: now}, nil
}

// validates the given JWT token and returns the decoded payload.
func (j *jwtProvider) Validate(myToken string) (tokenprovider.TokenPayload, error) {
	// Parse the JWT token with custom claims.
	res, err := jwt.ParseWithClaims(
		myToken,
		&myClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.secret), nil
		},
	)

	// Extract the claims from the result.
	claims, ok := res.Claims.(*myClaims)

	// Check for validation errors.
	if err != nil || !res.Valid || !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	return claims.Payload, nil
}