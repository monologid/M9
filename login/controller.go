package login

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/monologid/m9/serviceprovider"
)

// ProviderHandler ...
func ProviderHandler(c echo.Context) error {
	provider := c.Param("provider")

	serviceProvider, err := serviceprovider.New(provider)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	redirectURL := serviceProvider.GenerateOauthURI()

	MetricInitiateLogin(provider)

	return c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

// ProviderCallbackHandler ...
func ProviderCallbackHandler(c echo.Context) error {
	provider := c.Param("provider")
	code := c.QueryParam("code")

	serviceProvider, err := serviceprovider.New(provider)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	tokenURL := serviceProvider.GenerateGetAccessTokenURI(code)
	token, err := serviceProvider.GenerateAccessToken(tokenURL)
	if err != nil {
		MetricLoginFailed(provider)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	profileURL := serviceProvider.GenerateGetProfileURI(token.AccessToken)
	profile, err := serviceProvider.GetProfile(profileURL)
	if err != nil {
		MetricLoginFailed(provider)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	loginService := NewService(LoginRepository)

	// - Store data into database
	errInsert := loginService.SignInUsingServiceProvider(provider, *profile)
	if errInsert != nil {
		MetricLoginFailed(provider)
		return c.String(http.StatusInternalServerError, "failed to sign in using service provider "+provider)
	}

	// - Generate and redirect to LOGIN_REDIRECT_URI with ?accessToken=...
	accessToken, errGenerateAccessToken := loginService.GenerateAccessToken()
	if errGenerateAccessToken != nil {
		MetricLoginFailed(provider)
		return c.String(http.StatusInternalServerError, "failed to generate access token using service provider "+provider)
	}

	MetricLoginSuccess(provider)

	return c.JSON(http.StatusOK, map[string]interface{}{"accessToken": accessToken})
}

// APILoginHandler ...
func APILoginHandler(c echo.Context) error {
	return nil
}

// APIRegisterHandler ...
func APIRegisterHandler(c echo.Context) error {
	return nil
}
