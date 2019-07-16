package groups

import (
	"asira_lender/handlers"
	"asira_lender/middlewares"

	"github.com/labstack/echo"
)

func AdminGroup(e *echo.Echo) {
	g := e.Group("/admin")
	middlewares.SetClientJWTmiddlewares(g, "admin")

	// OTP
	g.GET("/info", handlers.AsiraAppInfo)
}
