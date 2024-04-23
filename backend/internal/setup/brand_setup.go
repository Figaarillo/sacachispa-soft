package setup

import (
	"github.com/Figaarillo/sacachispa-soft/internal/entity"
	"github.com/Figaarillo/sacachispa-soft/internal/handler"
	"github.com/Figaarillo/sacachispa-soft/internal/repository"
	"github.com/Figaarillo/sacachispa-soft/internal/router"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewBrand(initRouter *mux.Router, db *gorm.DB) {
	db.AutoMigrate(&entity.Brand{})

	repository := repository.NewBrandGorm(db)

	handler := handler.NewBrandHandler(repository)

	router := router.NewBrandRouter(initRouter, *handler)
	router.RegisterRoutes()
}
