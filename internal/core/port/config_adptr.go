package port

import "github.com/imtiaz246/codera_oj/internal/core/domain/dto"

// ConfigAdapter is an adapter for getting the necessary configuration
type ConfigAdapter interface {
	// GetAuthConfig returns the authN and authZ related configurations
	GetAuthConfig() *dto.AuthConfig
	// GetAppConfig returns the application setup related configuration
	GetAppConfig() *dto.AppConfig
	// GetServerConfig returns server setup related configuration
	GetServerConfig() *dto.ServerConfig
	// GetMailerConfig returns the setup configuration for mailer
	GetMailerConfig() *dto.EmailConfig
	// GetDatabaseConfig returns database setup related configuration
	GetDatabaseConfig() *dto.DatabaseConfig
}
