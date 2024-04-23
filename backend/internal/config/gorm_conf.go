package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm(env *EnvVars) *gorm.DB {
	DB_HOST := env.POSTGRES_HOST
	DB_PORT := env.POSTGRES_PORT
	DB_USER := env.POSTGRES_USER
	DB_PASSWORD := env.POSTGRES_PASSWORD
	DB_NAME := env.POSTGRES_DATABASE

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to database 🤟")

	return db
}
