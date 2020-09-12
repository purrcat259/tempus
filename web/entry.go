package web

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
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

	hasOngoingEntry, ongoingEntry, err := db.GetOngoingEntry(uint(projectID))
	if err != nil {
		// TODO: Add flashes
		return c.Redirect(http.StatusBadRequest, projectURL)
	}

	if hasOngoingEntry {

		// BUSINESS LOGIC DECISION: You cannot start a new task of the same type for now
		if entryType == ongoingEntry.EntryType {
			// TODO: Add flashes
			return c.Redirect(http.StatusFound, projectURL)
		}

		entryTypeSupported, err := db.ProjectSupportsEntryType(uint(projectID), entryType)

		if err != nil {
			// TODO: Add flashes
			return c.Redirect(http.StatusInternalServerError, projectURL)
		}

		if !entryTypeSupported {
			// TODO: Add flashes
			return c.Redirect(http.StatusFound, projectURL)
		}

		// Go to switch page
		entrySwitchURL := fmt.Sprintf("/projects/%d/entry/switch?newType=%s", projectID, url.QueryEscape(entryType))

		return c.Redirect(http.StatusFound, entrySwitchURL)

	} else {
		err := db.CreateEntry(db.DB, uint(projectID), entryType)

		if err != nil {
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

	err = db.CloseEntry(db.DB, uint(entryID))
	if err != nil {
		// TODO: Add flashes
		return c.Redirect(http.StatusBadRequest, projectURL)
	}

	return c.Redirect(http.StatusFound, projectURL)
}

func EntrySwitchPage(c echo.Context) error {
	// Route params: projectID
	// Query Param: newType
	session := fillDataFromContext(c)

	if !session.IsLoggedIn {
		return c.Redirect(http.StatusFound, "/")
	}

	projectIDParam := c.Param("projectID")
	projcetURL := fmt.Sprintf("/projects/%s", projectIDParam)
	encodedTargetEntryType := c.QueryParam("newType")

	targetEntryType, err := url.QueryUnescape(encodedTargetEntryType)

	if err != nil {
		return c.Redirect(http.StatusBadRequest, projcetURL)
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
	if !hasOngoingEntry {
		// No ongoing entry to switch from, go back to project page
		return c.Redirect(http.StatusFound, projcetURL)
	}
	session.Data["Project"] = project
	session.Data["OngoingEntry"] = ongoingEntry
	session.Data["TargetEntryType"] = targetEntryType
	return c.Render(http.StatusOK, "entryswitch", session)
}

func HandleSwitchEntry(c echo.Context) error {
	session := fillDataFromContext(c)
	if !session.IsLoggedIn {
		return c.Redirect(http.StatusForbidden, "/")
	}

	projectIDParam := c.Param("projectID")
	projectURL := fmt.Sprintf("/projects/%s", projectIDParam)
	targetEntryType := c.FormValue("TargetEntryType")

	projectID, err := strconv.Atoi(projectIDParam)
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

	hasOngoingEntry, _, err := db.GetOngoingEntry(uint(projectID))
	if err != nil {
		// TODO: Add flashes
		return c.Redirect(http.StatusFound, projectURL)
	}

	if !hasOngoingEntry {
		// TODO: Add flashes
		return c.Redirect(http.StatusFound, projectURL)
	}

	err = db.SwitchEntry(uint(projectID), targetEntryType)
	if err != nil {
		// TODO: Add flashes
		return c.Redirect(http.StatusBadRequest, projectURL)
	}

	return c.Redirect(http.StatusFound, projectURL)
}
