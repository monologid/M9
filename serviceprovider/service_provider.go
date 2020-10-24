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
)

// IProvider is an interface that should be implemented
// when creating a service provider object
type IProvider interface {
	Get() string

	GenerateOauthURI() string
	GenerateGetAccessTokenURI(code string) string
	GenerateGetProfileURI(accessToken string) string

	GenerateAccessToken(uri string) (*AccessTokenSchema, error)
	GetProfile(uri string) (*map[string]interface{}, error)
}

// New initates service provider object based on provided service provider
func New(serviceProvider string) (IProvider, error) {
	var provider IProvider
	var err error

	switch strings.ToUpper(serviceProvider) {
	case FACEBOOK:
		provider = NewFacebook()
	case GOOGLE:
		provider = NewGoogle()
	default:
		err = errors.New("unidentified service provider")
	}

	return provider, err
}
