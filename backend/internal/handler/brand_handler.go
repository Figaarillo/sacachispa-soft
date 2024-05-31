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
		util.HandleHTTPError(w, err, http.StatusNotFound)
		return
	}

	util.HandleHTTPResponse(w, "Brands retrieved successfully", http.StatusOK, brands)
}

func (h *BrandHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := util.GetURLParam(r, "id")
	if err != nil {
		util.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	brand, err := h.usecase.GetByID(id)
	if err != nil {
		util.HandleHTTPError(w, err, http.StatusNotFound)
		return
	}

	util.HandleHTTPResponse(w, "Brand retrieved successfully", http.StatusOK, brand)
}

func (h *BrandHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var brad entity.Brand
	if err := util.DecodeReqBody(r, &brad); err != nil {
		util.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	if err := h.usecase.Create(brad); err != nil {
		util.HandleHTTPError(w, err, http.StatusConflict)
		return
	}

	util.HandleHTTPResponse(w, "Brand created successfully", http.StatusCreated, nil)
}

func (h *BrandHandler) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, err := util.GetURLParam(r, "id")
	if err != nil {
		util.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	var brand entity.Brand
	if err := util.DecodeReqBody(r, &brand); err != nil {
		util.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	if err := h.usecase.Update(id, brand); err != nil {
		util.HandleHTTPError(w, err, http.StatusConflict)
		return
	}

	util.HandleHTTPResponse(w, "Brand updated successfully", http.StatusOK, nil)
}

func (h *BrandHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := util.GetURLParam(r, "id")
	if err != nil {
		util.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	if err := h.usecase.Delete(id); err != nil {
		util.HandleHTTPError(w, err, http.StatusNotFound)
		return
	}

	util.HandleHTTPResponse(w, "Brand deleted successfully", http.StatusOK, nil)
}
