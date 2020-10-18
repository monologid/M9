package serviceprovider

import (
	"errors"
	"strings"
)

var (
	// FACEBOOK service provider
	FACEBOOK = "FACEBOOK"

	// GOOGLE service provider
	GOOGLE = "GOOGLE"

	// TWITTER service provider
	TWITTER = "TWITTER"
)

// IProvider ...
type IProvider interface {
	Get() string

	GenerateOauthURI() string
	GenerateGetAccessTokenURI(code string) string
	GenerateGetProfileURI(accessToken string) string

	GenerateAccessToken(uri string) (*AccessTokenSchema, error)
	GetProfile(uri string) (*map[string]interface{}, error)
}

// Provider ...
type Provider struct {
}

// Get returns the service provider object
func (sp *Provider) Get(serviceProvider string) (IProvider, error) {
	var provider IProvider
	var err error

	switch strings.ToUpper(serviceProvider) {
	case FACEBOOK:
		provider = NewFacebook()
	case GOOGLE:
		provider = NewGoogle()
	// case TWITTER:
	// 	provider = NewTwitter()
	default:
		err = errors.New("unidentified service provider")
	}

	return provider, err
}
