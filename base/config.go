package base

import (
	"fmt"
	"log"
	"os"
	"parking-app-go/model"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
)

//LoadConfig ... will load config from .env file
func LoadConfig(config model.Config) model.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.DBPort = os.Getenv("db_port")
	config.DBHost = os.Getenv("db_host")
	config.DBPass = os.Getenv("db_password")
	config.DBType = os.Getenv("db_type")
	config.DBUser = os.Getenv("db_user")
	config.Port = os.Getenv("port")
	fmt.Println("test commit")

	validate := validator.New()
	err = validate.Struct(config)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Fatalf("Invalid .env file errors: %s", err)
		}
	}
	return config
}
