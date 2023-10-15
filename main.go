package main

import (
	"net/http"

	handler "overusevery/echo-psql/delivery"
	repository "overusevery/echo-psql/repository/psql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//setup client
	db := repository.SetupDB()
	defer db.Close()
	r := repository.NewPSQLTodoRepository(*db)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health", health)
	handler.TodoRouter(e, r)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func health(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
