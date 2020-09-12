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
	hasOngoingEntry, ongoingEntry, err := db.GetOngoingEntry(uint(projectID))
	if err != nil {
		return err
	}
	session.Data["Project"] = project
	session.Data["HasOngoingEntry"] = hasOngoingEntry
	session.Data["OngoingEntry"] = ongoingEntry
	return c.Render(http.StatusOK, "project", session)
}

func HandleCreateProject(c echo.Context) error {
	session := fillDataFromContext(c)
	title := c.FormValue("title")
	if !session.IsLoggedIn {
		return c.Redirect(http.StatusFound, "/")
	}

	alreadyExists := db.ProjectAlreadyExistsByTitleForUser(title, session.LoggedInUser.ID)
	if alreadyExists {
		return c.Redirect(http.StatusFound, "/dashboard")
	}

	err := db.CreateProject(title, session.LoggedInUser.ID)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/dashboard")
}
