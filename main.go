package main

import (
	"encoding/gob"
	"errors"
	"io"
	"log"
	"math"
	"os"
	"tempus/db"
	"tempus/web"
	"text/template"
	"time"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

// Define the template registry struct
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}

	return tmpl.ExecuteTemplate(w, "base.html", data)
}

var isProd bool
var domain string

func main() {
	// Env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	sessionSecret := os.Getenv("TEMPUS_SESSION_SECRET")
	if sessionSecret == "" {
		panic("Session secret env not set")
	}
	isProd = os.Getenv("TEMPUS_ENV") == "production"
	domain = "localhost"
	if isProd {
		domain = "tempus.simonam.dev"
	}

	// DB
	gob.Register(db.User{})

	// db.Clear()
	db.Open()
	db.Seed()
	defer db.DB.Close()

	// Web
	funcs := make(map[string]interface{})
	funcs["not"] = func(value interface{}) bool { return !value.(bool) }
	funcs["hasItems"] = func(value []interface{}) bool { return len(value) != 0 }
	funcs["isNil"] = func(value interface{}) bool { return value == nil }
	funcs["formatDate"] = func(t time.Time) string {
		return t.Format("Jan 2 15:04:05")
	}
	funcs["round"] = func(f float64) float64 {
		return math.Round(f*100) / 100
	}

	funcMap := template.FuncMap(funcs)

	// Echo instance
	e := echo.New()
	e.Debug = true
	e.Static("/static", "public/static")
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(sessionSecret))))
	e.Use(web.CreateTempusContextMW(domain, isProd))

	parseFuncs := func() *template.Template {
		return template.New("").Funcs(funcMap)
	}

	// Ref: https://gist.github.com/rand99/808e6e9702c00ce64803d94abff65678
	templates := make(map[string]*template.Template)
	templates["index"] = template.Must(parseFuncs().ParseFiles("public/views/index.html", "public/views/base.html"))
	templates["login"] = template.Must(parseFuncs().ParseFiles("public/views/login.html", "public/views/base.html"))
	templates["dashboard"] = template.Must(parseFuncs().ParseFiles("public/views/dashboard.html", "public/views/base.html"))
	templates["project"] = template.Must(parseFuncs().ParseFiles("public/views/project.html", "public/views/base.html"))
	templates["entryswitch"] = template.Must(parseFuncs().ParseFiles("public/views/entryswitch.html", "public/views/base.html"))

	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	e.GET("/", web.IndexPage)
	e.GET("/login", web.LoginPage)
	e.POST("/login", web.HandleLogin)
	e.GET("/logout", web.HandleLogout)
	e.GET("/dashboard", web.DashboardPage)
	e.POST("/projects/new", web.HandleCreateProject)
	e.GET("/projects/:projectID", web.ProjectPage)
	e.POST("/projects/:projectID/entry", web.HandleNewEntry)
	e.POST("/projects/:projectID/entry/types", web.HandleNewEntryType)
	e.GET("/projects/:projectID/entry/switch", web.EntrySwitchPage)
	e.POST("/projects/:projectID/entry/switch", web.HandleSwitchEntry)
	e.POST("/projects/:projectID/entry/:entryID", web.HandleCloseEntry)

	log.Printf("We in Prod? %t", isProd)
	e.Logger.Fatal(e.Start(":1323"))
}
