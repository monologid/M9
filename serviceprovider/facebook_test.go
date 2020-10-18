package serviceprovider_test

import (
	"os"
	"testing"

	"github.com/monologid/m9/serviceprovider"
	"github.com/stretchr/testify/assert"
)

func setFacebookEnvironmentVariables() {
	os.Setenv("FACEBOOK_GRAPH_URL", "https://graph.facebook.com/v4.0")
	os.Setenv("FACEBOOK_OAUTH_URL", "https://www.facebook.com/v4.0/dialog/oauth")
	os.Setenv("FACEBOOK_OAUTH_REDIRECT_URI", "https://monolog.id")
	os.Setenv("FACEBOOK_CLIENT_ID", "123")
	os.Setenv("FACEBOOK_CLIENT_SECRET", "123456789")
	os.Setenv("FACEBOOK_OAUTH_SCOPE", "email")
	os.Setenv("FACEBOOK_PROFILE_ATTRIBUTE", "id,first_name,last_name")
}

func TestFacebookShouldReturnServiceProvider(t *testing.T) {
	fb := serviceprovider.NewFacebook()
	assert.Equal(t, fb.Get(), serviceprovider.FACEBOOK)
}

func TestFacebookShouldReturnGenerateOauthURI(t *testing.T) {
	setFacebookEnvironmentVariables()

	fb := serviceprovider.NewFacebook()

	expectedURL := "https://www.facebook.com/v4.0/dialog/oauth?client_id=123&redirect_uri=https://monolog.id&scope=email"
	assert.Equal(t, fb.GenerateOauthURI(), expectedURL)
}

func TestFacebookShouldReturnGenerateGetAccessTokenURI(t *testing.T) {
	setFacebookEnvironmentVariables()

	fb := serviceprovider.NewFacebook()

	expectedURL := "https://graph.facebook.com/v4.0/oauth/access_token?client_id=123&client_secret=123456789&redirect_uri=https://monolog.id&code=12345&grant_type=authorization_code"
	code := "12345"
	assert.Equal(t, fb.GenerateGetAccessTokenURI(code), expectedURL)
}

func TestFacebookShouldReturnGenerateGetProfileURI(t *testing.T) {
	setFacebookEnvironmentVariables()

	fb := serviceprovider.NewFacebook()

	expectedURL := "https://graph.facebook.com/v4.0/me?fields=id,first_name,last_name&access_token=12345"
	accessToken := "12345"
	assert.Equal(t, fb.GenerateGetProfileURI(accessToken), expectedURL)
}
