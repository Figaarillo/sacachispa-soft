package util

import (
	"net/http"
	"strconv"

	"github.com/Figaarillo/sacachispa-soft/internal/exeption"
	"github.com/gorilla/mux"
)

func GetPagination(r *http.Request) (int, int) {
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	return offset, limit
}

func GetURLParam(r *http.Request, key string) (string, error) {
	param := mux.Vars(r)[key]

	if param == "" {
		return "", exeption.ErrorMissingParam
	}

	return param, nil
}
