package login

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// SocialMediaHandler ...
func SocialMediaHandler(c echo.Context) error {
	socialMedia := c.Param("social_media")

	service := NewService()

	var err error

	switch socialMedia {
	case "facebook":
		PrometheusLoginFacebookTotal.Inc()
		return c.Redirect(http.StatusTemporaryRedirect, service.FacebookOAuth())
	case "google":
	case "twitter":
	default:
		err = errors.New("invalid social media")
	}

	if err != nil {
		return err
	}

	return nil
}

// SocialMediaCallbackHandler ...
func SocialMediaCallbackHandler(c echo.Context) error {
	socialMedia := c.Param("social_media")

	service := NewService()

	var err error

	switch socialMedia {
	case "facebook":
		PrometheusLoginFacebookSuccessTotal.Inc()
		fbAccessTokenSchema, _ := service.FacebookOAuthGetAccessToken(c.QueryParam("code"))
		resp, _ := service.FacebookOAuthGetProfile(fbAccessTokenSchema.AccessToken)
		return c.JSON(http.StatusOK, resp)
	case "google":
	case "twitter":
	default:
		err = errors.New("invalid social media")
	}

	if err != nil {
		return err
	}

	return nil
}

// APILoginHandler ...
func APILoginHandler(c echo.Context) error {
	return nil
}

// APIRegisterHandler ...
func APIRegisterHandler(c echo.Context) error {
	return nil
}
