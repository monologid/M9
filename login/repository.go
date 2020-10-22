package login

import (
	"github.com/monologid/m9/db"
	"gorm.io/gorm"
)

// IRepository ...
type IRepository interface {
	FindOneByEmail(email string) (*AccountModel, error)
	Insert(account AccountModel) error
}

// Repository ...
type Repository struct {
	DB                          *gorm.DB
	TableAccount                *gorm.DB
	TableAccountServiceProvider *gorm.DB
}

// FindOneByEmail returns account data filtered by email
func (r *Repository) FindOneByEmail(email string) (*AccountModel, error) {
	var account AccountModel

	if err := r.TableAccount.Where("email = ?", email).Find(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}

// Insert ...
func (r *Repository) Insert(account AccountModel) error {
	account.GenerateID()
	account.AccountServiceProvider.GenerateID()

	return r.DB.Create(&account).Error
}

// NewRepository initiates login repository
func NewRepository() IRepository {
	return &Repository{
		DB:                          db.DB,
		TableAccount:                db.DB.Table("account"),
		TableAccountServiceProvider: db.DB.Table("account_service_provider"),
	}
}
