package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Init() {
	// Load .env file
	godotenv.Load()
	Database()
	app()

	// Set up Viper to automatically use environment variables
	viper.AutomaticEnv()
}

func app() {
	viper.SetDefault("PORT", "8080")

	conf("DEBUG", "false")

	conf("DB_HOST", "localhost")
	conf("DB_PORT", "3306")
	conf("DB_DATABASE", "testdb")
	conf("DB_USER", "root")
	conf("DB_PASSWORD", "secret")
	conf("DB_CHARSET", "utf8mb4")
	conf("DB_PARSETIME", "true")
	conf("DB_LOC", "Local")

	conf("JWT_SECRET", "secret")
	conf("JWT_DURATION", 720)

	conf("EMAIL_HOST", "")
	conf("EMAIL_PORT", "")
	conf("EMAIL_USERNAME", "")
	conf("EMAIL_PASSWORD", "")

	conf("HTTP_JSON_NAMING", "camel_case")
}
