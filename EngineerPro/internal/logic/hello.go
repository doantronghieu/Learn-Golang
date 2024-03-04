package logic

import "EngineerPro/internal/database"

type Hello struct {
	userDataAccessor database.UserDataAccessor
}

func SayHello(thing string) string {
	if thing == "" {
		return ""
	}

	return "Hello " + thing
}
