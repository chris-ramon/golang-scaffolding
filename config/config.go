package config

import (
	"encoding/base64"
	"os"
)

type Config struct {
	Port string

	// JWTConfig is the JWT configuration.
	JWTConfig *JWTConfig
}

type DBConfig struct {
	User    string
	PWD     string
	Host    string
	Name    string
	SSLMode string
}

func New() (*Config, error) {
	jwtConfig, err := NewJWTConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:      os.Getenv("PORT"),
		JWTConfig: jwtConfig,
	}, nil
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		User:    os.Getenv("DB_USER"),
		PWD:     os.Getenv("DB_PWD"),
		Host:    os.Getenv("DB_HOST"),
		Name:    os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSL_MODE"),
	}
}

// JWTConfig represents a JWT configuration.
type JWTConfig struct {
	// AppRsa is the RSA private key value.
	AppRsa []byte

	// AppRsaPub is the RSA public key value.
	AppRsaPub []byte
}

func NewJWTConfig() (*JWTConfig, error) {
	// openssl genrsa -out app.rsa 2048
	appRsa, err := base64.StdEncoding.DecodeString(os.Getenv("APP_RSA"))
	if err != nil {
		return nil, err
	}

	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	appRsaPub, err := base64.StdEncoding.DecodeString(os.Getenv("APP_RSA_PUB"))
	if err != nil {
		return nil, err
	}

	return &JWTConfig{
		AppRsa:    []byte(appRsa),
		AppRsaPub: []byte(appRsaPub),
	}, nil
}
