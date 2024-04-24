package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Figaarillo/sacachispa-soft/internal/config"
	"github.com/Figaarillo/sacachispa-soft/internal/setup"
)

func main() {
	env, err := config.NewEnvConf(".env")
	if err != nil {
		panic(err)
	}

	PORT := env.PORT
	HOST := env.HOST

	db := config.InitGorm(env)
	router := config.InitRouter()

	setup.NewBrand(router, db)
	setup.NewProduct(router, db)

	log.Printf("Server is running!🔥 Go to http://%s:%d/\n", HOST, PORT)
	http.ListenAndServe(fmt.Sprintf("%s:%d", HOST, PORT), router)
}
