package common

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
