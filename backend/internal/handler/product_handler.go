package handler

import (
	"net/http"

	"github.com/Figaarillo/sacachispa-soft/internal/entity"
	"github.com/Figaarillo/sacachispa-soft/internal/repository"
	"github.com/Figaarillo/sacachispa-soft/internal/usecase"
	"github.com/Figaarillo/sacachispa-soft/internal/util"
)

type ProductHandler struct {
	repository repository.ProductRepository
	usecase    *usecase.ProductUsecase
}

func NewProductHandler(repository repository.ProductRepository) *ProductHandler {
	return &ProductHandler{
		repository: repository,
		usecase:    usecase.NewProductUsecase(repository),
	}
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	offset, limit := util.GetPagination(r)

	products, err := h.usecase.List(offset, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	util.HandleHTTPResponse(w, "Products retrieved successfully", http.StatusOK, products)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := util.GetURLParam(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	product, err := h.usecase.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	util.HandleHTTPResponse(w, "Product retrieved successfully", http.StatusOK, product)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var product entity.Product
	if err := util.DecodeReqBody(r, &product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.usecase.Create(product); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	util.HandleHTTPResponse(w, "Product created successfully", http.StatusCreated)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, err := util.GetURLParam(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var product entity.Product
	if err := util.DecodeReqBody(r, &product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.usecase.Update(id, product); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	util.HandleHTTPResponse(w, "Product updated successfully", http.StatusOK)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := util.GetURLParam(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.usecase.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	util.HandleHTTPResponse(w, "Product deleted successfully", http.StatusOK)
}
