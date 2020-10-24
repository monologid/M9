package login

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/monologid/m9/config"
	"github.com/monologid/m9/serviceprovider"
)

// ProviderHandler is a handler for login using service provider.
// This handler will generate an oauth URI and
// redirect you to target service provider
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

// ProviderCallbackHandler is a callback handler to handle
// oauth redirect from service provider, usually service provider will include
// ?code=... that can be used to generate access token.
// And this handler will return an access token.
// If the login module redirect is TRUE, this handler will redirect
// to targer url and includes the access token, e.g.
// https://yourdomain.com?accessToken=...
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

	errInsert := loginService.SignInUsingServiceProvider(provider, *profile)
	if errInsert != nil {
		MetricLoginFailed(provider)
		return c.String(http.StatusInternalServerError, "failed to sign in using service provider "+provider)
	}

	accessToken, errGenerateAccessToken := loginService.GenerateAccessToken()
	if errGenerateAccessToken != nil {
		MetricLoginFailed(provider)
		return c.String(http.StatusInternalServerError, "failed to generate access token using service provider "+provider)
	}

	MetricLoginSuccess(provider)

	if config.C.Module.Login.IsRedirect {
		redirectURL := config.C.Module.Login.RedirectURI + "?accessToken=" + accessToken
		return c.Redirect(http.StatusTemporaryRedirect, redirectURL)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"accessToken": accessToken})
}

// APIRegisterHandler is a controller to handle new registration request
// and returns access token in the response in JSON format
func APIRegisterHandler(c echo.Context) error {
	provider := "default"
	MetricInitiateAccountRegistration(provider)

	var account ReqRegisterModel
	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"errmsg": "failed to create new account", "err": err.Error()})
	}

	loginService := NewService(LoginRepository)
	if err := loginService.RegisterNewAccount(account); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"errmsg": "failed to create new account", "err": err.Error()})
	}

	accessToken, errGenerateAccessToken := loginService.GenerateAccessToken()
	if errGenerateAccessToken != nil {
		MetricInitiateAccountRegistrationFailed(provider)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"errmsg": "failed to generate access token using service provider " + provider, "err": errGenerateAccessToken.Error()})
	}

	MetricInitiateAccountRegistrationSuccess(provider)
	return c.JSON(http.StatusOK, map[string]interface{}{"accessToken": accessToken})
}

// APILoginHandler is a controller to handle new login request
// and returns access token in the response in JSON format
func APILoginHandler(c echo.Context) error {
	provider := "default"
	MetricInitiateLogin(provider)

	var account ReqLoginModel
	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"errmsg": "failed to login", "err": err.Error()})
	}

	loginService := NewService(LoginRepository)
	if err := loginService.SignIn(account); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"errmsg": "failed to login", "err": err.Error()})
	}

	accessToken, errGenerateAccessToken := loginService.GenerateAccessToken()
	if errGenerateAccessToken != nil {
		MetricInitiateAccountRegistrationFailed(provider)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"errmsg": "failed to generate access token using service provider " + provider, "err": errGenerateAccessToken.Error()})
	}

	MetricLoginSuccess(provider)
	return c.JSON(http.StatusOK, map[string]interface{}{"accessToken": accessToken})
}

// APIValidateAccessTokenHandler validates access token
// return true if parsed and validated, otherwise false
func APIValidateAccessTokenHandler(c echo.Context) error {
	reqToken := c.QueryParam("token")

	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(config.C.Security.JWT.Secret), nil
	})

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"status": false})
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return c.JSON(http.StatusOK, map[string]interface{}{"status": true})
	}

	return c.JSON(http.StatusUnauthorized, map[string]interface{}{"status": false})
}
