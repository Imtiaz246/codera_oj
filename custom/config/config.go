package config

import (
	_ "embed"
	"fmt"
	"gopkg.in/yaml.v3"
)

var (
	//go:embed app.yaml
	ConfigFS []byte
	Cfg      AppSettings
)

func init() {
	if err := yaml.Unmarshal(ConfigFS, &Cfg); err != nil {
		panic(fmt.Errorf("setting init error: `%v`", err))
	}
}

type AppSettings struct {
	App      appConfig      `yaml:"app"`
	Server   serverConfig   `yaml:"server"`
	Database databaseConfig `yaml:"database"`
	Auth     authConfig     `yaml:"auth"`
	Email    emailConfig    `yaml:"email"`
}

type authConfig struct {
	Key                  string `yaml:"PASETO_SYMMETRIC_KEY"`
	AccessTokenDuration  string `yaml:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration string `yaml:"REFRESH_TOKEN_DURATION"`
}

type appConfig struct {
	AppName string `yaml:"APP_NAME"`
	RunMode string `yaml:"RUN_MODE"`
}

type serverConfig struct {
	Protocol string `yaml:"PROTOCOL"`
	Url      string `yaml:"URL"`
}

type emailConfig struct {
	SenderEmail string `yaml:"SENDER_EMAIL"`
	SenderPass  string `yaml:"SENDER_PASSWORD"`
}

type databaseConfig struct {
	DbType   string `yaml:"DB_TYPE"`
	Name     string `yaml:"NAME"`
	Host     string `yaml:"HOST"`
	Path     string `yaml:"PATH"`
	Username string `yaml:"USERNAME"`
	Password string `yaml:"PASSWORD"`
}
