package login

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/monologid/m9/db"
	"github.com/monologid/m9/util"
	"gorm.io/datatypes"
)

// Migrate executes database auto-migrate schema
// Please note that the auto migrate will create tables,
// missing foreign keys, constraints, columns and indexes,
// and will change existing column’s type if it’s size, precision,
// nullable changed, it WON’T delete unused columns to protect your data
func Migrate() {
	db.DB.AutoMigrate(
		&AccountModel{},
		&AccountServiceProviderModel{},
	)
}

// AccountModel is a schema for account
type AccountModel struct {
	db.Model

	FirstName  string `gorm:"type:varchar(255);not null;"`
	LastName   string `gorm:"type:varchar(255)"`
	Email      string `gorm:"type:varchar(255);not null;index"`
	Password   string `gorm:"type:varchar(255);"`
	PicURL     string
	Birthday   *time.Time
	IsVerified bool `gorm:"default:FALSE"`

	AccountServiceProvider AccountServiceProviderModel `gorm:"foreignKey:AccountID"`
}

// TableName returns account table name
func (m *AccountModel) TableName() string {
	return "account"
}

// AccountServiceProviderModel is a schema for account service provider
type AccountServiceProviderModel struct {
	db.Model

	AccountID         uuid.UUID `gorm:"not null;"`
	ServiceProvider   string    `gorm:"type:varchar(20);not null;"`
	ServiceProviderID string    `gorm:"type:varchar(50);not null;"`
	PicURL            string
	Metadata          datatypes.JSON
}

// TableName returns account service provider table name
func (m *AccountServiceProviderModel) TableName() string {
	return "account_service_provider"
}

// ReqRegisterModel is a schema for submitting new account registration
type ReqRegisterModel struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Validate validates registration submission
func (r ReqRegisterModel) Validate() error {
	if len(r.FirstName) == 0 {
		return errors.New("invalid first name")
	}

	var email util.Email = util.Email(r.Email)
	if !email.Validate() {
		return errors.New("invalid email address")
	}

	if len(r.Password) == 0 {
		return errors.New("invalid password")
	}

	return nil
}

// ReqLoginModel is a schema for submitting login
type ReqLoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate validates login request
func (r ReqLoginModel) Validate() error {
	var email util.Email = util.Email(r.Email)
	if !email.Validate() {
		return errors.New("invalid email address")
	}

	if len(r.Password) == 0 {
		return errors.New("invalid password")
	}

	return nil
}
