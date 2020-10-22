package login

import (
	"time"

	"github.com/google/uuid"
	"github.com/monologid/m9/db"
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

// AccountModel ...
type AccountModel struct {
	db.Model

	FirstName string `gorm:"type:varchar(255);not null;"`
	LastName  string `gorm:"type:varchar(255)"`
	Email     string `gorm:"type:varchar(255);not null;index"`
	Password  string `gorm:"type:varchar(50);"`
	PicURL    string
	Birthday  *time.Time

	AccountServiceProvider AccountServiceProviderModel `gorm:"foreignKey:AccountID"`
}

// TableName ...
func (m *AccountModel) TableName() string {
	return "account"
}

// AccountServiceProviderModel ...
type AccountServiceProviderModel struct {
	db.Model

	AccountID         uuid.UUID `gorm:"not null;"`
	ServiceProvider   string    `gorm:"type:varchar(20);not null;"`
	ServiceProviderID string    `gorm:"type:varchar(50);not null;"`
	PicURL            string
	Metadata          datatypes.JSON
}

// TableName ...
func (m *AccountServiceProviderModel) TableName() string {
	return "account_service_provider"
}
