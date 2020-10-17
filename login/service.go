package login

import (
	"encoding/json"
	"os"

	"github.com/go-resty/resty/v2"
)

// IService ...
type IService interface {
	FacebookOAuth() string
	FacebookOAuthGetAccessToken(code string) (*FacebookAccessTokenSchema, error)
	FacebookOAuthGetProfile(accessToken string) (*FacebookProfileSchema, error)
}

// Service ...
type Service struct{}

// FacebookOAuth ...
func (s *Service) FacebookOAuth() string {
	oauthURL := os.Getenv("FACEBOOK_OAUTH_URL")
	oauthURL += "?client_id=" + os.Getenv("FACEBOOK_CLIENT_ID")
	oauthURL += "&redirect_uri=" + os.Getenv("FACEBOOK_OAUTH_REDIRECT_URI")
	oauthURL += "&scope=" + os.Getenv("FACEBOOK_OAUTH_SCOPE")

	return oauthURL
}

// FacebookOAuthGetAccessToken ..
func (s *Service) FacebookOAuthGetAccessToken(code string) (*FacebookAccessTokenSchema, error) {
	getTokenURL := os.Getenv("FACEBOOK_GRAPH_URL") + "/oauth/access_token"
	getTokenURL += "?client_id=" + os.Getenv("FACEBOOK_CLIENT_ID")
	getTokenURL += "&client_secret=" + os.Getenv("FACEBOOK_CLIENT_SECRET")
	getTokenURL += "&redirect_uri=" + os.Getenv("FACEBOOK_OAUTH_REDIRECT_URI")
	getTokenURL += "&code=" + code
	getTokenURL += "&grant_type=authorization_code"

	resp, err := resty.New().R().Get(getTokenURL)
	if err != nil {
		return nil, err
	}

	var fbAccesTokenSchema FacebookAccessTokenSchema
	if err := json.Unmarshal(resp.Body(), &fbAccesTokenSchema); err != nil {
		return nil, err
	}

	return &fbAccesTokenSchema, nil
}

// FacebookOAuthGetProfile ...
func (s *Service) FacebookOAuthGetProfile(accessToken string) (*FacebookProfileSchema, error) {
	profileURL := os.Getenv("FACEBOOK_GRAPH_URL") + "/me"
	profileURL += "?fields=id,first_name,last_name,email,birthday"
	profileURL += "&access_token=" + accessToken

	resp, err := resty.New().R().Get(profileURL)
	if err != nil {
		return nil, err
	}

	var fbProfileSchema FacebookProfileSchema
	if err := json.Unmarshal(resp.Body(), &fbProfileSchema); err != nil {
		return nil, err
	}

	return &fbProfileSchema, nil
}

// NewService ...
func NewService() IService {
	return &Service{}
}
