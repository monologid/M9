package login

import (
	"github.com/monologid/m9/db"
	"gorm.io/gorm"
)

// IRepository is an interface for login repository
type IRepository interface {
	FindOneByEmail(email string) (*AccountModel, error)
	Insert(account AccountModel) error
}

// Repository is the implementation of login.IRepository
type Repository struct {
	DB *gorm.DB
}

// FindOneByEmail returns account data filtered by email
func (r *Repository) FindOneByEmail(email string) (*AccountModel, error) {
	var account AccountModel

	if err := r.DB.Table("account").Where("email = ?", email).Find(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}

// Insert inserts new account data
func (r *Repository) Insert(account AccountModel) error {
	account.GenerateID()
	account.AccountServiceProvider.GenerateID()

	return r.DB.Create(&account).Error
}

// NewRepository initiates login repository
func NewRepository() IRepository {
	return &Repository{
		DB: db.DB,
	}
}
