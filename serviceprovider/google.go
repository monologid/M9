package serviceprovider

import (
	"encoding/json"
	"os"

	"github.com/go-resty/resty/v2"
)

// Google ...
type Google struct {
	ServiceProvider  string
	APIURL           string
	OauthURL         string
	OauthTokenURL    string
	ClientID         string
	ClientSecret     string
	OauthRedirectURI string
	OauthScope       string
}

// Get returns ...
func (g *Google) Get() string {
	return g.ServiceProvider
}

// GenerateOauthURI ...
func (g *Google) GenerateOauthURI() string {
	return g.OauthURL + "/auth" +
		"?client_id=" + g.ClientID +
		"&redirect_uri=" + g.OauthRedirectURI +
		"&scope=" + g.OauthScope +
		"&access_type=offline&response_type=code"
}

// GenerateGetAccessTokenURI ...
func (g *Google) GenerateGetAccessTokenURI(code string) string {
	return g.OauthTokenURL +
		"?client_id=" + g.ClientID +
		"&client_secret=" + g.ClientSecret +
		"&redirect_uri=" + g.OauthRedirectURI +
		"&code=" + code +
		"&grant_type=authorization_code"
}

// GenerateGetProfileURI ...
func (g *Google) GenerateGetProfileURI(accessToken string) string {
	return g.APIURL + "/v3/userinfo?alt=json&access_token=" + accessToken
}

// GenerateAccessToken ...
func (g *Google) GenerateAccessToken(uri string) (*AccessTokenSchema, error) {
	resp, err := resty.New().R().Post(uri)
	if err != nil {
		return nil, err
	}

	var accesTokenSchema AccessTokenSchema
	if err := json.Unmarshal(resp.Body(), &accesTokenSchema); err != nil {
		return nil, err
	}

	return &accesTokenSchema, nil
}

// GetProfile ...
func (g *Google) GetProfile(uri string) (*map[string]interface{}, error) {
	resp, err := resty.New().R().Get(uri)
	if err != nil {
		return nil, err
	}

	var profileSchema map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &profileSchema); err != nil {
		return nil, err
	}

	return &profileSchema, nil
}

// NewGoogle ...
func NewGoogle() IProvider {
	apiURL := os.Getenv("GOOGLE_API_URL")
	if len(apiURL) == 0 {
		apiURL = "https://www.googleapis.com/oauth2"
	}

	oauthURL := os.Getenv("GOOGLE_OAUTH_URL")
	if len(oauthURL) == 0 {
		oauthURL = "https://accounts.google.com/o/oauth2"
	}

	oauthTokenURL := os.Getenv("GOOGLE_OAUTH_TOKEN_URL")
	if len(oauthTokenURL) == 0 {
		oauthTokenURL = "https://oauth2.googleapis.com/token"
	}

	return &Google{
		ServiceProvider:  GOOGLE,
		APIURL:           apiURL,
		OauthURL:         oauthURL,
		OauthTokenURL:    oauthTokenURL,
		ClientID:         os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret:     os.Getenv("GOOGLE_CLIENT_SECRET"),
		OauthRedirectURI: os.Getenv("GOOGLE_OAUTH_REDIRECT_URI"),
		OauthScope:       os.Getenv("GOOGLE_OAUTH_SCOPE"),
	}
}
