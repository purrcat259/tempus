package web

import (
	"net/http"
	"tempus/db"

	"github.com/labstack/echo/v4"
)

func DashboardPage(c echo.Context) error {
	session := fillDataFromContext(c)
	if !session.IsLoggedIn {
		return c.Redirect(http.StatusForbidden, "/")
	}
	user, err := db.GetUserByID(session.LoggedInUser.ID)
	if err != nil {
		return err
	}
	session.Data["Projects"] = user.Projects
	return c.Render(http.StatusOK, "dashboard", session)
}
