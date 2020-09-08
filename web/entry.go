package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tempus/db"

	"github.com/labstack/echo/v4"
)

func HandleNewEntryType(c echo.Context) error {
	session := fillDataFromContext(c)
	if !session.IsLoggedIn {
		return c.Redirect(http.StatusFound, "/")
	}

	projectIDParam := c.Param("projectID")
	projectURL := fmt.Sprintf("/projects/%s", projectIDParam)
	newEntryType := c.FormValue("entryType")

	projectID, err := strconv.Atoi(projectIDParam)
	if err != nil {
		// session.Success = false
		// session.Error = err.Error()
		// TODO: Add flashes
		return c.Redirect(http.StatusBadRequest, projectURL)
	}
	projectExists := db.ProjectExists(uint(projectID))
	if !projectExists {
		return c.Redirect(http.StatusBadRequest, "/")
	}

	ownedByUser, err := db.ProjectIsOwnedByUser(uint(projectID), session.LoggedInUser.ID)
	if err != nil {
		// session.Success = false
		// session.Error = err.Error()
		// TODO: Add flashes
		return c.Redirect(http.StatusBadRequest, projectURL)
	}

	if !ownedByUser {
		return c.Redirect(http.StatusForbidden, "/")
	}

	entryTypeAlreadyExists := db.EntryTypeExistsInProject(newEntryType, uint(projectID))
	if entryTypeAlreadyExists {
		return c.Redirect(http.StatusFound, projectURL)
	}

	err = db.CreateEntryType(newEntryType, uint(projectID))
	if err != nil {
		log.Println(err.Error())
		// session.Success = false
		// session.Error = err.Error()
		// TODO: Add flashes
		return c.Redirect(http.StatusInternalServerError, projectURL)
	}

	return c.Redirect(http.StatusFound, projectURL)
}

func HandleNewEntry(c echo.Context) error {
	session := fillDataFromContext(c)
	if !session.IsLoggedIn {
		return c.Redirect(http.StatusFound, "/")
	}

	projectIDParam := c.Param("projectID")
	projectURL := fmt.Sprintf("/projects/%s", projectIDParam)
	entryType := c.FormValue("entryType")

	projectID, err := strconv.Atoi(projectIDParam)
	if err != nil {
		// session.Success = false
		// session.Error = err.Error()
		// TODO: Add flashes
		return c.Redirect(http.StatusBadRequest, projectURL)
	}
	projectExists := db.ProjectExists(uint(projectID))
	if !projectExists {
		return c.Redirect(http.StatusBadRequest, "/")
	}

	ownedByUser, err := db.ProjectIsOwnedByUser(uint(projectID), session.LoggedInUser.ID)
	if err != nil {
		// session.Success = false
		// session.Error = err.Error()
		// TODO: Add flashes
		return c.Redirect(http.StatusBadRequest, projectURL)
	}

	if !ownedByUser {
		return c.Redirect(http.StatusForbidden, "/")
	}

	hasOngoingEntry, _, err := db.GetOngoingEntry(uint(projectID))
	// hasOngoingEntry, ongoingEntry, err := db.GetOngoingEntry(uint(projectID))
	if err != nil {
		// session.Success = false
		// session.Error = err.Error()
		// TODO: Add flashes
		return c.Redirect(http.StatusBadRequest, projectURL)
	}

	if hasOngoingEntry {
		// Go to switch page

		log.Println("Switch page")

		// TODO

		return c.Redirect(http.StatusFound, projectURL)

	} else {
		err := db.CreateEntry(uint(projectID), entryType)

		if err != nil {
			// session.Success = false
			// session.Error = err.Error()
			// TODO: Add flashes
			return c.Redirect(http.StatusBadRequest, projectURL)
		}

		return c.Redirect(http.StatusFound, projectURL)
	}

}

func HandleCloseEntry(c echo.Context) error {
	session := fillDataFromContext(c)
	if !session.IsLoggedIn {
		return c.Redirect(http.StatusForbidden, "/")
	}

	projectIDParam := c.Param("projectID")
	projectURL := fmt.Sprintf("/projects/%s", projectIDParam)
	entryIDParam := c.Param("entryID")

	projectID, err := strconv.Atoi(projectIDParam)
	if err != nil {
		// TODO: Add flashes
		return c.Redirect(http.StatusBadRequest, projectURL)
	}
	entryID, err := strconv.Atoi(entryIDParam)
	if err != nil {
		// TODO: Add flashes
		return c.Redirect(http.StatusBadRequest, projectURL)
	}

	projectExists := db.ProjectExists(uint(projectID))
	if !projectExists {
		return c.Redirect(http.StatusBadRequest, "/")
	}

	ownedByUser, err := db.ProjectIsOwnedByUser(uint(projectID), session.LoggedInUser.ID)
	if err != nil {
		// TODO: Add flashes
		return c.Redirect(http.StatusBadRequest, projectURL)
	}

	if !ownedByUser {
		return c.Redirect(http.StatusForbidden, "/")
	}

	entryExists := db.EntryExists(uint(entryID))
	if !entryExists {
		return c.Redirect(http.StatusBadRequest, projectURL)
	}

	err = db.CloseEntry(uint(entryID))
	if err != nil {
		// TODO: Add flashes
		return c.Redirect(http.StatusBadRequest, projectURL)
	}

	return c.Redirect(http.StatusFound, projectURL)

}
