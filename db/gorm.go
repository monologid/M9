package db

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DB is a global variable for database
var DB *gorm.DB

// NewMysql returns database object using mysql as database engine
func NewMysql(dsn string, dbconf *gorm.Config) (*gorm.DB, error) {
	dbconn, err := gorm.Open(mysql.Open(dsn), dbconf)
	if err != nil {
		return nil, err
	}

	return dbconn, nil
}

// NewPostgresql returns database object using postgres as database engine
func NewPostgresql(dsn string, dbconf *gorm.Config) (*gorm.DB, error) {
	dbconn, err := gorm.Open(postgres.Open(dsn), dbconf)
	if err != nil {
		return nil, err
	}

	return dbconn, nil
}

// New initiates new database connection based on submitted database engine
func New(dbEngine string, dsn string) {
	var dbconn *gorm.DB
	var err error

	dbconf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	switch dbEngine {
	case "mysql":
		dbconn, err = NewMysql(dsn, dbconf)
	case "postgres":
		dbconn, err = NewPostgresql(dsn, dbconf)
	default:
		err = errors.New("invalid database engine")
	}

	if err != nil {
		panic(err)
	}

	DB = dbconn
}

// SetVerbose sets verbose mode on database query
func SetVerbose(verbose bool) {
	if verbose {
		DB = DB.Debug()
	}
}
