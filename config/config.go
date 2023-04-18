package config

type Config struct {
	Port             uint
	JWTSigningSecret string
}

func New(port uint, JWTSigningSecret string) *Config {
	return &Config{
		Port: port,
	}
}
