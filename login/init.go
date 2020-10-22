package login

import (
	"github.com/labstack/echo/v4"
)

var (
	// LoginRepository ...
	LoginRepository IRepository
)

// New ...
func New(app *echo.Echo) {
	LoginRepository = NewRepository()

	app.GET("/login/oauth2/:provider", ProviderHandler)
	app.GET("/login/oauth2/:provider/callback", ProviderCallbackHandler)

	app.GET("/api/v1/login", APILoginHandler)
	app.POST("/api/v1/register", APIRegisterHandler)
}
