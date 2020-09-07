package web

import (
	"net/http"
	"strconv"
	"tempus/db"

	"github.com/labstack/echo/v4"
)

func ProjectPage(c echo.Context) error {
	session := fillDataFromContext(c)
	projectIDParam := c.Param("projectID")
	if !session.IsLoggedIn {
		return c.Redirect(http.StatusFound, "/")
	}
	projectID, err := strconv.Atoi(projectIDParam)
	if err != nil {
		return err
	}
	userIsOwner, err := db.ProjectIsOwnedByUser(uint(projectID), session.LoggedInUser.ID)
	if err != nil {
		return err
	}
	if !userIsOwner {
		return c.Redirect(http.StatusForbidden, "/")
	}
	project, err := db.GetProjectByID(uint(projectID))
	if err != nil {
		return err
	}
	session.Data["Project"] = project
	return c.Render(http.StatusOK, "project", session)
}
