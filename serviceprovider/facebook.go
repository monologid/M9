package serviceprovider

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/monologid/m9/config"
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
	Fields           string
}

// Get returns service provider
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
		"?fields=" + fb.Fields +
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
func (fb *Facebook) GetProfile(uri string) (*map[string]interface{}, error) {
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

// NewFacebook ...
func NewFacebook() IProvider {
	var graphqlURL string = config.C.Facebook.GraphURL
	if len(graphqlURL) == 0 {
		graphqlURL = "https://graph.facebook.com/v4.0"
	}

	var oauthURL string = config.C.Facebook.Oauth.URL
	if len(oauthURL) == 0 {
		oauthURL = "https://www.facebook.com/v4.0/dialog/oauth"
	}

	return &Facebook{
		ServiceProvider:  FACEBOOK,
		GraphURL:         graphqlURL,
		OauthURL:         oauthURL,
		ClientID:         config.C.Facebook.ClientID,
		ClientSecret:     config.C.Facebook.ClientSecret,
		OauthRedirectURI: config.C.Facebook.Oauth.RedirectURI,
		OauthScope:       config.C.Facebook.Oauth.Scope,
		Fields:           config.C.Facebook.Oauth.Fields,
	}
}
