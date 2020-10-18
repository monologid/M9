package serviceprovider_test

import (
	"os"
	"testing"

	"github.com/monologid/m9/serviceprovider"
	"github.com/stretchr/testify/assert"
)

func setGoogleEnvironmentVariables() {
	os.Setenv("GOOGLE_API_URL", "https://www.googleapis.com/oauth2")
	os.Setenv("GOOGLE_OAUTH_URL", "https://accounts.google.com/o/oauth2")
	os.Setenv("GOOGLE_OAUTH_REDIRECT_URI", "https://monolog.id")
	os.Setenv("GOOGLE_CLIENT_ID", "123")
	os.Setenv("GOOGLE_CLIENT_SECRET", "123456789")
	os.Setenv("GOOGLE_OAUTH_SCOPE", "https://www.googleapis.com/auth/userinfo.email")
}

func TestGoogleShouldReturnServiceProvider(t *testing.T) {
	fb := serviceprovider.NewGoogle()
	assert.Equal(t, fb.Get(), serviceprovider.GOOGLE)
}

func TestGoogleShouldReturnGenerateOauthURI(t *testing.T) {
	setGoogleEnvironmentVariables()

	fb := serviceprovider.NewGoogle()

	expectedURL := "https://accounts.google.com/o/oauth2/auth?client_id=123&redirect_uri=https://monolog.id&scope=https://www.googleapis.com/auth/userinfo.email&access_type=offline&response_type=code"
	assert.Equal(t, fb.GenerateOauthURI(), expectedURL)
}

func TestGoogleShouldReturnGenerateGetAccessTokenURI(t *testing.T) {
	setGoogleEnvironmentVariables()

	fb := serviceprovider.NewGoogle()

	expectedURL := "https://www.googleapis.com/oauth2/v4/token?client_id=123&client_secret=123456789&redirect_uri=https://monolog.id&code=12345&grant_type=authorization_code"
	code := "12345"
	assert.Equal(t, fb.GenerateGetAccessTokenURI(code), expectedURL)
}

func TestGoogleShouldReturnGenerateGetProfileURI(t *testing.T) {
	setGoogleEnvironmentVariables()

	fb := serviceprovider.NewGoogle()

	expectedURL := "https://www.googleapis.com/oauth2/v3/userinfo?alt=json&access_token=12345"
	accessToken := "12345"
	assert.Equal(t, fb.GenerateGetProfileURI(accessToken), expectedURL)
}
