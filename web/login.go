package web

import (
	"net/http"
	"tempus/db"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func LoginPage(c echo.Context) error {
	session := fillDataFromContext(c)
	if session.IsLoggedIn {
		return c.Redirect(http.StatusFound, "/")
	}
	return c.Render(http.StatusOK, "login", session)
}

func HandleLogin(c echo.Context) error {
	tc := c.(*TempusContext)
	sessionFromContext := fillDataFromContext(c)
	if sessionFromContext.IsLoggedIn {
		return c.Redirect(http.StatusFound, "/")
	}
	sess, _ := session.Get("session", c)

	// If not logged in, check email and password
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Confirm email exists
	user, err := db.GetUserByEmail(email)
	if err != nil {
		sessionFromContext.Success = false
		sessionFromContext.Error = "Your email/password combination are incorrect"
		return c.Render(http.StatusOK, "login", sessionFromContext)
	}
	// Confirm password matches
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		sessionFromContext.Success = false
		sessionFromContext.Error = "Your email/password combination are incorrect."
		return c.Render(http.StatusOK, "login", sessionFromContext)
	}
	sess.Options = &sessions.Options{
		Domain:   tc.Domain,
		Path:     "/",
		MaxAge:   365 * 24 * 60 * 60 * 1000,
		HttpOnly: true,
		Secure:   tc.IsProd,
	}
	sess.Values["user"] = user
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		sessionFromContext.Success = false
		sessionFromContext.Error = err.Error()
		return c.Render(http.StatusOK, "login", sessionFromContext)
	}
	return c.Redirect(http.StatusFound, "/")
}

func HandleLogout(c echo.Context) error {
	tc := c.(*TempusContext)
	if !tc.IsLoggedIn {
		return c.Redirect(http.StatusFound, "/")
	}
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Domain:   tc.Domain,
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   tc.IsProd,
	}
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusFound, "/")
}
