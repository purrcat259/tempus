package web

import (
	"net/http"
	"tempus/db"

	"github.com/labstack/echo/v4"
)

func IndexPage(c echo.Context) error {
	session := fillDataFromContext(c)
	users, err := db.GetAllUsers()
	if err != nil {
		return err
	}
	session.Data["Users"] = users
	return c.Render(http.StatusOK, "index", session)
}
