package config

type Config struct {
	Env        string
	Port       string
	Logging    LoggingConfig
	Cors       CorsConfig
	Version    string
	Database   DatabaseConfig
	Monitoring Monitoring
	Auth       AuthConfig
}
