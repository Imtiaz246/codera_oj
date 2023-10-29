package config

import (
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
	"github.com/imtiaz246/codera_oj/internal/core/ports"
)

type ConfigAdapter struct {
	// viper
}

var _ ports.ConfigAdapter = (*ConfigAdapter)(nil)

func NewConfigAdapter() ports.ConfigAdapter {
	return &ConfigAdapter{}
}

func (c *ConfigAdapter) GetAuthConfig() *dto.AuthConfig {
	return nil
}

func (c *ConfigAdapter) GetAppConfig() *dto.AppConfig {
	return nil
}

func (c *ConfigAdapter) GetServerConfig() *dto.ServerConfig {
	return nil
}

func (c *ConfigAdapter) GetDatabaseConfig() *dto.DatabaseConfig {
	return nil
}

func (c *ConfigAdapter) GetMailerConfig() *dto.EmailConfig {
	return nil
}
