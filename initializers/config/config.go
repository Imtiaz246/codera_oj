package config

import (
	_ "embed"
	"errors"
	"gopkg.in/yaml.v3"
)

var (
	//go:embed app.yaml
	AppConfigFile []byte

	// GlobalCfg contains the configuration data
	// of the whole app.
	GlobalCfg config
)

type config struct {
	App      AppConfig      `yaml:"app"`
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Auth     AuthConfig     `yaml:"auth"`
	Email    EmailConfig    `yaml:"email"`
}

type AppConfig struct {
	AppName string `yaml:"APP_NAME"`
	RunMode string `yaml:"RUN_MODE"`
}

type DatabaseConfig struct {
	DbType   string `yaml:"DB_TYPE"`
	Name     string `yaml:"NAME"`
	Host     string `yaml:"HOST"`
	Path     string `yaml:"PATH"`
	Username string `yaml:"USERNAME"`
	Password string `yaml:"PASSWORD"`
}

type ServerConfig struct {
	Protocol string `yaml:"PROTOCOL"`
	Domain   string `yaml:"DOMAIN"`
	Port     string `yaml:"PORT"`
}

type AuthConfig struct {
	PublicKey  string `yaml:"PUBLIC_KEY"`
	PrivateKey string `yaml:"PRIVATE_KEY"`
}

type EmailConfig struct {
	SenderEmail string `yaml:"SENDER_EMAIL"`
	SenderPass  string `yaml:"SENDER_PASSWORD"`
}

func LoadConfigs() error {
	// Load the app.yaml file content to GlobalCfg variable
	if err := yaml.Unmarshal(AppConfigFile, &GlobalCfg); err != nil {
		return errors.New("load configs failed")
	}
	return nil
}

func GetDBConfig() *DatabaseConfig {
	return &GlobalCfg.Database
}

func GetAuthConfig() *AuthConfig {
	return &GlobalCfg.Auth
}

func GetEmailConfig() *EmailConfig {
	return &GlobalCfg.Email
}

func GetServerConfig() *ServerConfig {
	return &GlobalCfg.Server
}
