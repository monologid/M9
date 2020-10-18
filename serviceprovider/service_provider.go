package serviceprovider

import (
	"errors"
	"strings"
)

var (
	// FACEBOOK ...
	FACEBOOK = "FACEBOOK"

	// GOOGLE ...
	GOOGLE = "GOOGLE"

	// TWITTER ...
	TWITTER = "TWITTER"
)

// IProvider ...
type IProvider interface {
	Get() string

	GenerateOauthURI() string
	GenerateGetAccessTokenURI(code string) string
	GenerateGetProfileURI(accessToken string) string

	GenerateAccessToken(uri string) (*AccessTokenSchema, error)
	GetProfile(uri string) (*ProfileSchema, error)
}

// Provider ...
type Provider struct {
}

// Get returns ...
func (sp *Provider) Get(providerType string) (IProvider, error) {
	var provider IProvider
	var err error

	switch strings.ToUpper(providerType) {
	case FACEBOOK:
		provider = NewFacebook()
	case GOOGLE:
		provider = NewGoogle()
	case TWITTER:
		provider = NewGoogle()
	default:
		err = errors.New("unidentified service provider")
	}

	return provider, err
}
