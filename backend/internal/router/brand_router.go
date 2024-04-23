package router

import (
	"net/http"

	"github.com/Figaarillo/sacachispa-soft/internal/handler"
	"github.com/gorilla/mux"
)

type BrandRouter struct {
	router  *mux.Router
	handler handler.BrandHandler
}

func NewBrandRouter(router *mux.Router, handler handler.BrandHandler) *BrandRouter {
	subrouter := router.PathPrefix("/api/brands").Subrouter()

	return &BrandRouter{
		router:  subrouter,
		handler: handler,
	}
}

func (r *BrandRouter) RegisterRoutes() {
	r.router.HandleFunc("", r.handler.List).Methods(http.MethodGet)
	r.router.HandleFunc("/{id}", r.handler.GetByID).Methods(http.MethodGet)
	r.router.HandleFunc("", r.handler.Create).Methods(http.MethodPost)
	r.router.HandleFunc("/{id}", r.handler.Update).Methods(http.MethodPut)
	r.router.HandleFunc("/{id}", r.handler.Delete).Methods(http.MethodDelete)
}
