package login

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/monologid/m9/serviceprovider"
)

// ProviderHandler ...
func ProviderHandler(c echo.Context) error {
	provider := c.Param("provider")

	serviceProvider, err := new(serviceprovider.Provider).Get(provider)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	redirectURL := serviceProvider.GenerateOauthURI()

	return c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

// ProviderCallbackHandler ...
func ProviderCallbackHandler(c echo.Context) error {
	provider := c.Param("provider")
	code := c.QueryParam("code")

	serviceProvider, err := new(serviceprovider.Provider).Get(provider)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	tokenURL := serviceProvider.GenerateGetAccessTokenURI(code)
	token, err := serviceProvider.GenerateAccessToken(tokenURL)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	profileURL := serviceProvider.GenerateGetProfileURI(token.AccessToken)
	profile, err := serviceProvider.GetProfile(profileURL)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, profile)
}

// APILoginHandler ...
func APILoginHandler(c echo.Context) error {
	return nil
}

// APIRegisterHandler ...
func APIRegisterHandler(c echo.Context) error {
	return nil
}
