package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

var configVariables map[string]string

func LoadEnv() error {
	var err error
	configVariables, err = godotenv.Read(".env")
	if err != nil {
		fmt.Print("Error loading .env file", err.Error())
		return err
	}

	return nil
}

func GetVar(key string) string {
	if configVariables == nil {
		log.Printf("Warning: Config variables not loaded")
		return ""
	}

	value, isValid := configVariables[key]
	if !isValid {
		log.Printf("Warning: Config key '%s' not found", key)
		return ""
	}

	return value
}
