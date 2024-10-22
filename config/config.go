package config

import (
	"os"
)

type Config struct {
	ServiceName string

	Postgres struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"DBName"`
		SSLMode  string `json:"sslMode"`
		PgDriver string `json:"pgDriver"`
	} `json:"postgres"`

	Server struct {
		Host                        string `json:"host"`
		Port                        string `json:"port"`
		ShowUnknownErrorsInResponse bool   `json:"showUnknownErrorsInResponse"`
	} `json:"server"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		ServiceName: "Thumbnail Service",
		Postgres: struct {
			Host     string `json:"host"`
			Port     string `json:"port"`
			User     string `json:"user"`
			Password string `json:"password"`
			DBName   string `json:"DBName"`
			SSLMode  string `json:"sslMode"`
			PgDriver string `json:"pgDriver"`
		}{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DATABASE"),
			SSLMode:  "disable",
			PgDriver: "pgx",
		},
		Server: struct {
			Host                        string `json:"host"`
			Port                        string `json:"port"`
			ShowUnknownErrorsInResponse bool   `json:"showUnknownErrorsInResponse"`
		}{
			Host:                        os.Getenv("SERVER_HOST"),
			Port:                        os.Getenv("SERVER_PORT"),
			ShowUnknownErrorsInResponse: false,
		},
	}

	return cfg, nil
}
