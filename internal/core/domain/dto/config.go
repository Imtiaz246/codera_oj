package dto

import "time"

type TokenConfig struct {
	Key                  []byte
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

type AppConfig struct {
	AppName string
	RunMode string
}

type ServerConfig struct {
	Protocol string
	Url      string
}

type EmailConfig struct {
	SenderEmail string
	SenderPass  string
}

type DatabaseConfig struct {
	DbType   string
	Name     string
	Host     string
	Path     string
	Username string
	Password string
}
