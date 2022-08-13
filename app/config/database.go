package config

type DatabaseConfig struct {
	URL      string
	Host     string
	Database string
	User     string
	Password string
	Port     string
	SSLMode  string
	Options  string
}
