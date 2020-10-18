package serviceprovider

import (
	"encoding/json"
	"os"

	"github.com/go-resty/resty/v2"
)

// Facebook ...
type Facebook struct {
	ServiceProvider  string
	GraphURL         string
	OauthURL         string
	ClientID         string
	ClientSecret     string
	OauthRedirectURI string
	OauthScope       string
	ProfileAttribute string
}

// Get returns ...
func (fb *Facebook) Get() string {
	return fb.ServiceProvider
}

// GenerateOauthURI ...
func (fb *Facebook) GenerateOauthURI() string {
	return fb.OauthURL +
		"?client_id=" + fb.ClientID +
		"&redirect_uri=" + fb.OauthRedirectURI +
		"&scope=" + fb.OauthScope
}

// GenerateGetAccessTokenURI ...
func (fb *Facebook) GenerateGetAccessTokenURI(code string) string {
	return fb.GraphURL + "/oauth/access_token" +
		"?client_id=" + fb.ClientID +
		"&client_secret=" + fb.ClientSecret +
		"&redirect_uri=" + fb.OauthRedirectURI +
		"&code=" + code +
		"&grant_type=authorization_code"
}

// GenerateGetProfileURI ...
func (fb *Facebook) GenerateGetProfileURI(accessToken string) string {
	return fb.GraphURL + "/me" +
		"?fields=" + fb.ProfileAttribute +
		"&access_token=" + accessToken
}

// GenerateAccessToken ...
func (fb *Facebook) GenerateAccessToken(uri string) (*AccessTokenSchema, error) {
	resp, err := resty.New().R().Get(uri)
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
func (fb *Facebook) GetProfile(uri string) (*ProfileSchema, error) {
	resp, err := resty.New().R().Get(uri)
	if err != nil {
		return nil, err
	}

	var profileSchema ProfileSchema
	if err := json.Unmarshal(resp.Body(), &profileSchema); err != nil {
		return nil, err
	}

	return &profileSchema, nil
}

// NewFacebook ...
func NewFacebook() IProvider {
	return &Facebook{
		ServiceProvider:  FACEBOOK,
		GraphURL:         os.Getenv("FACEBOOK_GRAPH_URL"),
		OauthURL:         os.Getenv("FACEBOOK_OAUTH_URL"),
		ClientID:         os.Getenv("FACEBOOK_CLIENT_ID"),
		ClientSecret:     os.Getenv("FACEBOOK_CLIENT_SECRET"),
		OauthRedirectURI: os.Getenv("FACEBOOK_OAUTH_REDIRECT_URI"),
		OauthScope:       os.Getenv("FACEBOOK_OAUTH_SCOPE"),
		ProfileAttribute: os.Getenv("FACEBOOK_PROFILE_ATTRIBUTE"),
	}
}
