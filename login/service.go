package login

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/monologid/m9/config"
	"golang.org/x/crypto/bcrypt"
)

// IService is an interface for login service
type IService interface {
	SignInUsingServiceProvider(serviceProvider string, profile map[string]interface{}) error
	RegisterNewAccountUsingServiceProvider(serviceProvider string, profile map[string]interface{}) error

	RegisterNewAccount(newAccount ReqRegisterModel) error
	SignIn(loginReq ReqLoginModel) error

	GenerateAccessToken() (string, error)
}

// Service is the implementation of login.IService
type Service struct {
	LoginRepository IRepository

	Account AccountModel
}

// SignInUsingServiceProvider returns nil once the login or registration using service provider is succeeded
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

// RegisterNewAccountUsingServiceProvider returns nil when registration new account using service provider is succeded.
// This method will parse data from service provider and construct it using the AccountModel and AccountServiceModel.
// All data other than basic fields will be stored in AccountServiceModel.Metadata.
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
	} else {
		return errors.New("invalid service provider")
	}

	s.Account = account

	return s.LoginRepository.Insert(account)
}

// RegisterNewAccount returns nil if registration using default method is success.
func (s *Service) RegisterNewAccount(newAccount ReqRegisterModel) error {
	errValidate := newAccount.Validate()
	if errValidate != nil {
		return errValidate
	}

	tempAccount, err := s.LoginRepository.FindOneByEmail(newAccount.Email)
	if err != nil {
		return err
	}

	if tempAccount.Email == newAccount.Email && tempAccount.ID.String() != "00000000-0000-0000-0000-000000000000" {
		return errors.New("email has been registered")
	}

	account := AccountModel{
		FirstName: newAccount.FirstName,
		LastName:  newAccount.LastName,
		Email:     newAccount.Email,
	}

	accountServiceProvider := AccountServiceProviderModel{
		ServiceProvider: "default",
	}

	account.AccountServiceProvider = accountServiceProvider

	passwd, err := bcrypt.GenerateFromPassword([]byte(newAccount.Password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
		return errors.New("failed to encrypt password")
	}

	account.Password = string(passwd)

	s.Account = account

	return s.LoginRepository.Insert(account)
}

// SignIn returns nil if email and password are verified
func (s *Service) SignIn(loginReq ReqLoginModel) error {
	errValidate := loginReq.Validate()
	if errValidate != nil {
		return errValidate
	}

	tempAccount, err := s.LoginRepository.FindOneByEmail(loginReq.Email)
	if err != nil {
		return err
	}

	if tempAccount.Email != loginReq.Email && tempAccount.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return errors.New("invalid email address")
	}

	if errCheckPassword := bcrypt.CompareHashAndPassword([]byte(tempAccount.Password), []byte(loginReq.Password)); errCheckPassword != nil {
		return errCheckPassword
	}

	s.Account = *tempAccount

	return nil
}

// GenerateAccessToken returns access token in JWT format.
// The claim is based on data when being set in registration or login methods using the AccountModel.
func (s *Service) GenerateAccessToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":        s.Account.ID,
		"firstName":  s.Account.FirstName,
		"lastName":   s.Account.LastName,
		"email":      s.Account.Email,
		"profilePic": s.Account.PicURL,
	})

	return token.SignedString([]byte(config.C.Security.JWT.Secret))
}

// NewService initiates new login service
func NewService(loginRepository IRepository) IService {
	return &Service{
		LoginRepository: loginRepository,
	}
}
