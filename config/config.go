package config

type Config struct {
	Port uint
}

func New(port uint) *Config {
	return &Config{
		Port: port,
	}
}
