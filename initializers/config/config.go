package config

import (
	_ "embed"
	"fmt"
	"gopkg.in/yaml.v3"
)

// ErrLoadConfigFailed indicates the error for the failure
// of loading configuration files
const ErrLoadConfigFailed = "load config failed: %v"

var (
	//go:embed app.yaml
	AppConfigFile []byte

	// GlobalCfg contains the configuration data
	// of the whole app.
	GlobalCfg config
)

// LoadConfigs loads the app config
func LoadConfigs() error {
	// Load the app.yaml file content to GlobalCfg variable
	if err := yaml.Unmarshal(AppConfigFile, &GlobalCfg); err != nil {
		return fmt.Errorf(ErrLoadConfigFailed, err)
	}
	return nil
}

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
	Url      string `yaml:"URL"`
}

type AuthConfig struct {
	Key                  string `yaml:"PASETO_SYMMETRIC_KEY"`
	AccessTokenDuration  string `yaml:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration string `yaml:"REFRESH_TOKEN_DURATION"`
}

type EmailConfig struct {
	SenderEmail string `yaml:"SENDER_EMAIL"`
	SenderPass  string `yaml:"SENDER_PASSWORD"`
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

func GetAppConfig() *AppConfig {
	return &GlobalCfg.App
}
