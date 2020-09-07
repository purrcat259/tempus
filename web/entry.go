package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleNewEntryType(c echo.Context) error {
	sessionFromContext := fillDataFromContext(c)
	if !sessionFromContext.IsLoggedIn {
		return c.Redirect(http.StatusFound, "/")
	}

	projectID := c.Param("projectID")
	entryType := c.FormValue("entryType")
	log.Println(projectID, entryType)

	return c.Redirect(http.StatusFound, fmt.Sprintf("/projects/%s", projectID))
}

func HandleNewEntry(c echo.Context) error {
	sessionFromContext := fillDataFromContext(c)
	if !sessionFromContext.IsLoggedIn {
		return c.Redirect(http.StatusFound, "/")
	}

	projectID := c.Param("projectID")
	entryType := c.QueryParam("entryType")
	log.Println(projectID, entryType)

	return c.Redirect(http.StatusFound, fmt.Sprintf("/projects/%s", projectID))
}
