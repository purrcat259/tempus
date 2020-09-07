package web

import (
	"tempus/db"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func CreateTempusContextMW(domain string, isProd bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Compute wheter user is logged in
			isLoggedIn := false
			var user db.User
			sess, err := session.Get("session", c)
			if err == nil && sess != nil {
				userSess := sess.Values["user"]
				isLoggedIn = userSess != nil
				if isLoggedIn {
					user = userSess.(db.User)
				}
			}
			// The rest are just params
			tc := &TempusContext{
				Context:    c,
				Domain:     domain,
				IsProd:     isProd,
				IsLoggedIn: isLoggedIn,
				User:       user,
			}
			return next(tc)
		}
	}
}
