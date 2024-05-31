package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Figaarillo/sacachispa-soft/internal/config"
	"github.com/Figaarillo/sacachispa-soft/internal/setup"
)

func main() {
	env, err := config.NewEnvConf()
	if err != nil {
		log.Fatalf("could not load config file: %v", err)
	}

	PORT := env.SERVER_PORT

	db := config.InitGorm(env)
	router := config.InitRouter()

	setup.NewBrand(router, db)
	setup.NewProduct(router, db)

	log.Printf("Server is running!🔥 Go to http://localhost:%d/\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), router)
}
