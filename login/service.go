package login

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/monologid/m9/config"
)

// IService ...
type IService interface {
	SignInUsingServiceProvider(serviceProvider string, profile map[string]interface{}) error
	RegisterNewAccountUsingServiceProvider(serviceProvider string, profile map[string]interface{}) error
	GenerateAccessToken() (string, error)
}

// Service ...
type Service struct {
	LoginRepository IRepository

	Account AccountModel
}

// SignInUsingServiceProvider ...
func (s *Service) SignInUsingServiceProvider(serviceProvider string, profile map[string]interface{}) error {
	email, ok := profile["email"]
	if !ok {
		return errors.New("unable to retrieve profile from" + serviceProvider)
	}

	account, err := s.LoginRepository.FindOneByEmail(fmt.Sprintf("%v", email))
	if err != nil {
		return err
	}

	if account.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return s.RegisterNewAccountUsingServiceProvider(serviceProvider, profile)
	}

	s.Account = *account

	return nil
}

// RegisterNewAccountUsingServiceProvider ...
func (s *Service) RegisterNewAccountUsingServiceProvider(serviceProvider string, profile map[string]interface{}) error {
	MetricInitiateAccountRegistration(serviceProvider)

	var account AccountModel

	metadata, err := json.Marshal(profile)
	if err != nil {
		MetricInitiateAccountRegistrationFailed(serviceProvider)
		return errors.New("failed to convert profile metadata")
	}

	account.AccountServiceProvider.ServiceProvider = serviceProvider
	account.AccountServiceProvider.Metadata = metadata

	if serviceProvider == "facebook" {
		account.Email = fmt.Sprintf("%v", profile["email"])
		account.FirstName = fmt.Sprintf("%v", profile["first_name"])
		account.LastName = fmt.Sprintf("%v", profile["last_name"])

		account.AccountServiceProvider.ServiceProviderID = fmt.Sprintf("%v", profile["id"])
		if picURL, ok := profile["picture"].(map[string]interface{})["data"].(map[string]interface{})["url"]; ok {
			account.PicURL = fmt.Sprintf("%v", picURL)
			account.AccountServiceProvider.PicURL = account.PicURL
		}
	} else if serviceProvider == "google" {
		account.Email = fmt.Sprintf("%v", profile["email"])
		account.FirstName = fmt.Sprintf("%v", profile["given_name"])
		account.LastName = fmt.Sprintf("%v", profile["family_name"])
		account.PicURL = fmt.Sprintf("%v", profile["picture"])

		account.AccountServiceProvider.ServiceProviderID = fmt.Sprintf("%v", profile["sub"])
		account.AccountServiceProvider.PicURL = account.PicURL
	}

	s.Account = account

	return s.LoginRepository.Insert(account)
}

// Register ...
func (s *Service) Register(serviceProvider, profile map[string]interface{}) error {
	return nil
}

// GenerateAccessToken ...
func (s *Service) GenerateAccessToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"firstName":  s.Account.FirstName,
		"lastName":   s.Account.LastName,
		"email":      s.Account.Email,
		"profilePic": s.Account.PicURL,
	})

	return token.SignedString([]byte(config.C.Security.JWT.Secret))
}

// NewService ...
func NewService(loginRepository IRepository) IService {
	return &Service{
		LoginRepository: loginRepository,
	}
}
