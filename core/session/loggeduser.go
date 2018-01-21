package session

import (
	"github.com/gorilla/sessions"
)

type LoggedUser struct {
	UserId int
	TypeId int8
	Name   string
	Email  string
}

func GetLoggedUser(s *sessions.Session) interface{} {
	if s.Values["UserId"] == nil {
		return nil
	}

	return &LoggedUser{s.Values["UserId"].(int), s.Values["UserTypeId"].(int8), s.Values["Name"].(string), s.Values["Email"].(string)}
}
