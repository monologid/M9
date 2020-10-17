package login

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func emptyHandler(c echo.Context) error {
	return nil
}

func New(app *echo.Echo) {
	app.GET("/login/:social_media", SocialMediaHandler)
	app.GET("/login/:social_media/callback", SocialMediaCallbackHandler)

	app.GET("/api/v1/login", APILoginHandler)
	app.GET("/api/v1/login/metrics", emptyHandler, echo.WrapMiddleware(func(handler http.Handler) http.Handler {
		return promhttp.Handler()
	}))
	app.POST("/api/v1/register", APIRegisterHandler)
}
