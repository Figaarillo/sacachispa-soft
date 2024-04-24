package setup

import (
	"github.com/Figaarillo/sacachispa-soft/internal/entity"
	"github.com/Figaarillo/sacachispa-soft/internal/handler"
	"github.com/Figaarillo/sacachispa-soft/internal/repository"
	"github.com/Figaarillo/sacachispa-soft/internal/router"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewProduct(initRouter *mux.Router, db *gorm.DB) {
	db.AutoMigrate(&entity.Product{})

	repository := repository.NewProductGorm(db)

	handler := handler.NewProductHandler(repository)

	router := router.NewProductRouter(initRouter, *handler)
	router.RegisterRoutes()
}
