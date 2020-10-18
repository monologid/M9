package serviceprovider_test

import (
	"testing"

	"github.com/monologid/m9/serviceprovider"
	"github.com/stretchr/testify/assert"
)

func TestGetServiceProviderShouldReturnError(t *testing.T) {
	serviceProvider, err := new(serviceprovider.Provider).Get("test")

	assert.Nil(t, serviceProvider)
	assert.Error(t, err)
}

func TestGetServiceProviderGetShouldReturnFacebook(t *testing.T) {
	serviceProvider, err := new(serviceprovider.Provider).Get("facebook")

	assert.Equal(t, serviceProvider.Get(), serviceprovider.FACEBOOK)
	assert.Nil(t, err)
}

func TestGetServiceProviderGetShouldReturnGoogle(t *testing.T) {
	serviceProvider, err := new(serviceprovider.Provider).Get("google")

	assert.Equal(t, serviceProvider.Get(), serviceprovider.GOOGLE)
	assert.Nil(t, err)
}
