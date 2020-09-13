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
	projectsWithOngoingEntry, err := db.GetAllOngoingEntriesForUser(session.LoggedInUser.ID)
	if err != nil {
		return err
	}
	session.Data["Projects"] = user.Projects
	session.Data["ProjectsWithOngoingEntry"] = projectsWithOngoingEntry
	return c.Render(http.StatusOK, "dashboard", session)
}
