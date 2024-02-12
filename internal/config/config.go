package config

import "os"

type Config struct {
	Secret string
}

func NewConfig() Config {
	secret := os.Getenv("SECRET")

	return Config{
		Secret: secret,
	}
}
