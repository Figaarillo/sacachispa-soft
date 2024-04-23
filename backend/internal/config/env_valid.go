package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type EnvVars struct {
	HOST                       string        `validate:"required,hostname"`
	POSTGRES_HOST              string        `validate:"required,hostname"`
	POSTGRES_USER              string        `validate:"required"`
	POSTGRES_PASSWORD          string        `validate:"required"`
	POSTGRES_DATABASE          string        `validate:"required"`
	PORT                       uint          `validate:"required,gte=1,lte=65535"`
	POSTGRES_PORT              uint          `validate:"required,gte=1,lte=65535"`
	POSTGRES_MAX_CONNS         int           `validate:"required,gte=1"`
	POSTGRES_MAX_IDLE_CONNS    int           `validate:"required,gte=1"`
	POSTGRES_MAX_CONN_LIFETIME time.Duration `validate:"required"`
}

func NewEnvConf(envFile string) (*EnvVars, error) {
	err := godotenv.Load(envFile)
	if err != nil {
		return &EnvVars{}, fmt.Errorf("error cargando el archivo .env: %w", err)
	}

	env := &EnvVars{
		HOST:                       os.Getenv("HOST"),
		PORT:                       uint(atoi(os.Getenv("PORT"))),
		POSTGRES_HOST:              os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT:              uint(atoi(os.Getenv("POSTGRES_PORT"))),
		POSTGRES_USER:              os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD:          os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_DATABASE:          os.Getenv("POSTGRES_DATABASE"),
		POSTGRES_MAX_CONNS:         atoi(os.Getenv("POSTGRES_MAX_CONNS")),
		POSTGRES_MAX_IDLE_CONNS:    atoi(os.Getenv("POSTGRES_MAX_IDLE_CONNS")),
		POSTGRES_MAX_CONN_LIFETIME: time.Duration(atoi(os.Getenv("POSTGRES_MAX_CONN_LIFETIME"))) * time.Second,
	}

	err = validateEnvVars(env)
	if err != nil {
		return &EnvVars{}, err
	}

	return env, nil
}

func validateEnvVars(env *EnvVars) error {
	validate := validator.New()
	if err := validate.Struct(env); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(*env).FieldByName(err.StructField()) // Cambio aquí
			return fmt.Errorf("error en el campo: %s, condición: %s", field.Name, err.Tag())
		}
	}
	return nil
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
