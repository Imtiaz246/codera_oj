package config

import (
	_ "embed"
	"fmt"
	"gopkg.in/yaml.v3"
)

var (
	//go:embed app.yaml
	ConfigFS []byte
	Settings AppSettings
)

func init() {
	if err := yaml.Unmarshal(ConfigFS, &Settings); err != nil {
		panic(fmt.Errorf("setting init error: `%v`", err))
	}
}

type AppSettings struct {
	App      AppConfig      `yaml:"app"`
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Auth     AuthConfig     `yaml:"auth"`
	Email    EmailConfig    `yaml:"email"`
}

type AuthConfig struct {
	Key                  string `yaml:"PASETO_SYMMETRIC_KEY"`
	AccessTokenDuration  string `yaml:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration string `yaml:"REFRESH_TOKEN_DURATION"`
}

type AppConfig struct {
	AppName string `yaml:"APP_NAME"`
	RunMode string `yaml:"RUN_MODE"`
}

type ServerConfig struct {
	Protocol string `yaml:"PROTOCOL"`
	Url      string `yaml:"URL"`
}

type EmailConfig struct {
	SenderEmail string `yaml:"SENDER_EMAIL"`
	SenderPass  string `yaml:"SENDER_PASSWORD"`
}

type DatabaseConfig struct {
	DbType   string `yaml:"DB_TYPE"`
	Name     string `yaml:"NAME"`
	Host     string `yaml:"HOST"`
	Path     string `yaml:"PATH"`
	Username string `yaml:"USERNAME"`
	Password string `yaml:"PASSWORD"`
}
