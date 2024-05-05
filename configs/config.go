package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    DatabaseURL string
    Port        string
}

type URL struct{
    URL string
}

func Load() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        return nil, err
    }

    config := &Config{
        DatabaseURL: os.Getenv("DATABASE_URL"),
        Port:        os.Getenv("PORT"),
    }

    return config, nil
}

func ReturnURL() (string, error) {
    err := godotenv.Load()
    if err != nil {
        return "", err
    }

    url := os.Getenv("URL")
    return url, nil
}
