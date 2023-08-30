package handler

import (
	config "github.com/lucioPintanel/personal_finance_project/configs"
	"gorm.io/gorm"
)

var (
	Logger *config.Logger
	Db     *gorm.DB
)

func InitializeHandler() {
	Logger = config.GetLogger("handlres")
	Db = config.GetSQLite()
}
