package config

import "os"

type Config struct {
	Port uint
}

type DBConfig struct {
	User    string
	PWD     string
	Host    string
	Name    string
	SSLMode string
}

func New(port uint) *Config {
	return &Config{
		Port: port,
	}
}

func NewDBConfig(user, pwd, host, name, sslMode string) *DBConfig {
	return &DBConfig{
		User:    os.Getenv("DB_USER"),
		PWD:     os.Getenv("DB_PWD"),
		Host:    os.Getenv("DB_HOST"),
		Name:    os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSL_MODE"),
	}
}
