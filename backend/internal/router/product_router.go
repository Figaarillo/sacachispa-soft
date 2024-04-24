package router

import (
	"net/http"

	"github.com/Figaarillo/sacachispa-soft/internal/handler"
	"github.com/gorilla/mux"
)

type ProductRouter struct {
	router  *mux.Router
	handler handler.ProductHandler
}

func NewProductRouter(router *mux.Router, handler handler.ProductHandler) *ProductRouter {
	subrouter := router.PathPrefix("/api/products").Subrouter()

	return &ProductRouter{
		router:  subrouter,
		handler: handler,
	}
}

func (r *ProductRouter) RegisterRoutes() {
	r.router.HandleFunc("", r.handler.List).Methods(http.MethodGet)
	r.router.HandleFunc("/{id}", r.handler.GetByID).Methods(http.MethodGet)
	r.router.HandleFunc("", r.handler.Create).Methods(http.MethodPost)
	r.router.HandleFunc("/{id}", r.handler.Update).Methods(http.MethodPut)
	r.router.HandleFunc("/{id}", r.handler.Delete).Methods(http.MethodDelete)
}
