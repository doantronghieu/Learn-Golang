package common

import "log"

func Recovery() {
	if r := recover(); r != nil {
		log.Println("Recovered:", r)
	}
}

// represents the structure of the JWT token payload.
type TokenPayload struct {
	UId   int    `json:"user_id"`
	URole string `json:"role"`
}

// returns the user ID from the token payload.
func (p TokenPayload) UserId() int {
	return p.UId
}

// returns the role from the token payload.
func (p TokenPayload) Role() string {
	return p.URole
}