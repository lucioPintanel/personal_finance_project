package configs

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type APIConfig struct {
	Port    string
	ApiPath string
	Version string
}

type DBSQLiteConfig struct {
	DbDir  string
	DbPath string
}

type configApp struct {
	apicfg      APIConfig
	dbsqlitecfg DBSQLiteConfig
}

var (
	db     *gorm.DB
	logger *Logger
	cfgApp *configApp
)

func init() {
	viper.SetDefault("api.port", "9090")
	viper.SetDefault("api.apipath", "/api")
	viper.SetDefault("api.version", "/v1")
	viper.SetDefault("database.dbdir", "./db")
	viper.SetDefault("database.dbfile", "/personal_finance.db")
}

func Init() error {
	var err error
	err = InitConfigApp()
	if err != nil {
		return err
	}
	db, err = InitializeSQLite()
	if err != nil {
		return err
	}
	return nil
}

func InitConfigApp() error {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // path to look for the config file in
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}
	cfgApp = new(configApp)
	cfgApp.apicfg = APIConfig{
		Port:    viper.GetString("api.port"),
		ApiPath: viper.GetString("api.apipath"),
		Version: viper.GetString("api.version"),
	}
	cfgApp.dbsqlitecfg = DBSQLiteConfig{
		DbDir:  viper.GetString("database.dbdir"),
		DbPath: viper.GetString("database.dbfile"),
	}
	return nil
}

func GetServerPort() string {
	return cfgApp.apicfg.Port
}

func GetServerApiPath() string {
	return cfgApp.apicfg.ApiPath
}

func GetServerVersion() string {
	return cfgApp.apicfg.Version
}

func GetDbDir() string {
	return cfgApp.dbsqlitecfg.DbDir
}

func GetDbPath() string {
	return cfgApp.dbsqlitecfg.DbPath
}

func GetSQLite() *gorm.DB {
	return db
}
func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
