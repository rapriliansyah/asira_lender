package router

import (
	"asira_lender/groups"
	"asira_lender/handlers"

	"github.com/labstack/echo"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	// e.GET("/test", handlers.Test)
	e.GET("/clientauth", handlers.ClientLogin)

	groups.AdminGroup(e)
	groups.ClientGroup(e)
	groups.LenderGroup(e)

	return e
}
