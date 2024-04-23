package handler

import (
	"net/http"

	"github.com/Figaarillo/sacachispa-soft/internal/entity"
	"github.com/Figaarillo/sacachispa-soft/internal/repository"
	"github.com/Figaarillo/sacachispa-soft/internal/usecase"
	"github.com/Figaarillo/sacachispa-soft/internal/util"
)

type BrandHandler struct {
	repository repository.BrandRepository
	usecase    *usecase.BrandUsecase
}

func NewBrandHandler(repository repository.BrandRepository) *BrandHandler {
	return &BrandHandler{
		repository: repository,
		usecase:    usecase.NewBrandUsecase(repository),
	}
}

func (h *BrandHandler) List(w http.ResponseWriter, r *http.Request) {
	offset, limit := util.GetPagination(r)

	brands, err := h.usecase.List(offset, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	util.HandleHTTPResponse(w, "Brands retrieved successfully", brands)
}

func (h *BrandHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := util.GetURLParam(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	brand, err := h.usecase.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	util.HandleHTTPResponse(w, "Brand retrieved successfully", brand)
}

func (h *BrandHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var brad entity.Brand

	if err := util.DecodeReqBody(r, &brad); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	if err := h.usecase.Create(brad); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	util.HandleHTTPResponse(w, "Brand created successfully")
}
