package web

import (
	"tempus/db"

	"github.com/labstack/echo/v4"
)

type TempusContext struct {
	echo.Context
	IsLoggedIn bool
	User       db.User
	Domain     string
	IsProd     bool
}

type SessionMeta struct {
	Success      bool
	Error        string
	IsLoggedIn   bool
	LoggedInUser db.User
}

type SessionData struct {
	SessionMeta
	Data map[string]interface{}
}

func fillDataFromContext(context interface{}) SessionData {
	tc := context.(*TempusContext)
	meta := SessionMeta{
		Success:      true,
		Error:        "",
		IsLoggedIn:   tc.IsLoggedIn,
		LoggedInUser: tc.User,
	}
	sessiondata := SessionData{
		SessionMeta: meta,
		Data:        make(map[string]interface{}),
	}
	return sessiondata
}
