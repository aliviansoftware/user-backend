package config

import (
	"os"
	"user-backend/pkg"
)

func GetConfig() *pkg.Config {
	return &pkg.Config{
		Mongo: &pkg.MongoConfig{
			Ip:     envOrDefaultString("user-backend:mongo:ip", "127.0.0.1:27017"),
			DbName: envOrDefaultString("user-backend:mongo:dbName", "user-backend")},
		Server: &pkg.ServerConfig{Port: envOrDefaultString("user-backend:server:port", ":1377")},
		Auth:   &pkg.AuthConfig{Secret: envOrDefaultString("user-backend:auth:secret", "mysecret")}}
}

func envOrDefaultString(envVar string, defaultValue string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}

	return value
}
