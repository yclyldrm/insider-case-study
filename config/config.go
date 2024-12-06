package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

var configVariables map[string]string

func LoadEnv(key string) error {
	var err error
	configVariables, err = godotenv.Read(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
		return err
	}

	return nil
}

func GetVar(key string) string {
	value, isValid := configVariables[key]
	if !isValid {
		return ""
	}

	return value
}
