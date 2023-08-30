package config

import (
	"os"

	"github.com/glebarez/sqlite"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/personal_finance.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database not found...creating")
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Opening erro: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(
		&schemas.Tipo{},
		&schemas.TipoItem{},
		&schemas.User{},
		&schemas.Card{},
	)
	if err != nil {
		logger.Errorf("AutoMigrate erro: %v", err)
		return nil, err
	}
	return db, err
}
