package database

import (
	"fmt"
	"task/app/application/interface/database"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type databaseHandller struct {
	db *gorm.DB
}

func NewDatabaseHandller() database.IDatabaseHandller {
	return &databaseHandller{}
}

func (dh *databaseHandller) OpenDB(url string) error {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("postgres://%s", url)), &gorm.Config{})
	if err != nil {
		return err
	}

	dh.db = db
	return nil
}

func (dh *databaseHandller) CloseDB() error {
	sqlDB, _ := dh.db.DB()
	if err := sqlDB.Close(); err != nil {
		return err
	}

	return nil
}

func (dh *databaseHandller) Transaction(fn func(database.IRepositoryFactory) error) error {
	return dh.db.Transaction(func(tx *gorm.DB) error {
		return fn(NewRepositoryFactory(tx))
	})
}
