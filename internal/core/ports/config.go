package ports

import "github.com/imtiaz246/codera_oj/internal/core/domain/dto"

// ConfigAdapter is an interface that provides access to various configurations required by the application.
type ConfigAdapter interface {
	// GetAuthConfig retrieves the authentication and authorization-related configuration.
	GetAuthConfig() *dto.TokenConfig

	// GetAppConfig retrieves the application setup configuration.
	GetAppConfig() *dto.AppConfig

	// GetServerConfig retrieves the server setup configuration.
	GetServerConfig() *dto.ServerConfig

	// GetMailerConfig retrieves the configuration for email services.
	GetMailerConfig() *dto.EmailConfig

	// GetDatabaseConfig retrieves the configuration for database setup.
	GetDatabaseConfig() *dto.DatabaseConfig
}
