package login

import (
	"github.com/labstack/echo/v4"
)

var (
	LoginRepository IRepository
)

// New initiates login module by registering the API endpoint into the server
func New(app *echo.Echo) {
	LoginRepository = NewRepository()

	app.GET("/login/oauth2/:provider", ProviderHandler)
	app.GET("/login/oauth2/:provider/callback", ProviderCallbackHandler)

	app.POST("/login/api/v1/sign-up", APIRegisterHandler)
	app.POST("/login/api/v1/sign-in", APILoginHandler)
	app.GET("/login/api/v1/access-token/validate", APIValidateAccessTokenHandler)
}
