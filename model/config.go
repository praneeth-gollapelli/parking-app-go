package model

type Config struct {
	DBPort string `validate:"required"`
	DBUser string `validate:"required"`
	DBHost string `validate:"required"`
	DBPass string `validate:"required"`
	DBType string `validate:"required"`
	Port   string `validate:"required"`
}
