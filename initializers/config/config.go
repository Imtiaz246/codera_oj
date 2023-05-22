package config

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

type Config struct {
	server   serverConfig
	database databaseConfig
	auth     authConfig
	email    emailConfig
}

var (
	GlobalCfg Config
)

type databaseConfig struct {
	DB_TYPE  string
	NAME     string
	HOST     string
	PATH     string
	USERNAME string
	PASSWORD string
}

type serverConfig struct {
	PROTOCOL string
	DOMAIN   string
	PORT     string
}

type authConfig struct {
	JWTSecret []byte
}

type emailConfig struct {
	SenderAddr string
	SenderPass string
}

func LoadConfigs() error {
	GlobalCfg = Config{
		server: serverConfig{
			PROTOCOL: "http",
			DOMAIN:   "localhost",
			PORT:     "3000",
		},
		database: databaseConfig{
			DB_TYPE: "sqlite3",
			PATH:    "gorm.sqlite",
		},
		auth: authConfig{
			JWTSecret: []byte("MY_TEST_JWT_SECRET"),
		},
		email: emailConfig{
			SenderAddr: "imtiazuddincho001@gmail.com",
			SenderPass: "jlwtvuagzwztpdez",
		},
	}

	return nil
	// todo: resolve panic, use viper instead
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}
	configDir := workingDir + "/config/app.ini"
	cfg, err := ini.Load(configDir)
	if err != nil {
		return err
	}

	err = cfg.MapTo(&GlobalCfg)
	if err != nil {
		return err
	}

	log.Println(GlobalCfg)

	return nil
}

func GetDBConfig() *databaseConfig {
	return &GlobalCfg.database
}

func GetAuthConfig() *authConfig {
	return &GlobalCfg.auth
}

func GetEmailConfig() *emailConfig {
	return &GlobalCfg.email
}

func GetServerConfig() *serverConfig {
	return &GlobalCfg.server
}
