package util

import (
	"net/http"
	"strconv"
)

func GetPagination(r *http.Request) (int, int) {
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	return offset, limit
}
