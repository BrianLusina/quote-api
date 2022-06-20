package config

type CorsConfig struct {
	AllowedOrigins   string
	AllowedMethods   string
	AllowedHeaders   string
	AllowCredentials string
	MaxAge           string
	ExposedHeaders   string
}
