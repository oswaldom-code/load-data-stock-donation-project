package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/oswaldom-code/load-data-stock-donation-project/pkg/config"
	"github.com/oswaldom-code/load-data-stock-donation-project/pkg/log"
	"github.com/oswaldom-code/load-data-stock-donation-project/src/services/ports"
)

// store handles the database context
type store struct {
	db *gorm.DB
}

var dbStore *store

// New returns a new instance of a Store
func New(dsn config.DBConfig) ports.Repository {
	var err error
	db := &gorm.DB{}
	//var db *gorm.DB
	log.DebugWithFields("Creating new database connection",
		log.Fields{"dsn": dsn})

	switch dsn.Engine {
	case "postgre":
		dsnStrConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=%s TimeZone=America/Lima",
			dsn.Host, dsn.User, dsn.Password, dsn.Database, dsn.Port, dsn.SSLMode)
		db, err = gorm.Open(postgres.Open(dsnStrConnection), &gorm.Config{SkipDefaultTransaction: true, FullSaveAssociations: false})
	case "mysql":
		dsnStrConnection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", dsn.User, dsn.Password, dsn.Host, dsn.Port, dsn.Database)
		db, err = gorm.Open(mysql.Open(dsnStrConnection), &gorm.Config{SkipDefaultTransaction: true, FullSaveAssociations: false, Logger: logger.Default.LogMode(4)})
	default:
		log.ErrorWithFields("Invalid database engine", log.Fields{"engine": dsn.Engine})
	}
	if err != nil {
		log.ErrorWithFields("error connecting to db ", log.Fields{
			"engine":   dsn.Engine,
			"host":     dsn.Host,
			"port":     dsn.Port,
			"database": dsn.Database,
			"username": dsn.User,
			"err":      err,
		})
		os.Exit(1)
	}
	dbStore = &store{db: db.Set("gorm:auto_preload", true)}
	return dbStore
}

func NewRepository() ports.Repository {
	log.DebugWithFields("Creating new database connection",
		log.Fields{"dsn": config.GetDBConfig()})
	if dbStore == nil {
		New(config.GetDBConfig())
		return dbStore
	}
	return dbStore
}
