package util

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Body interface{} `json:"body"`
	Msg  string      `json:"message"`
}

func HandleHTTPResponse(w http.ResponseWriter, msg string, code int, body ...interface{}) {
	var res Response

	res.Body = body
	res.Msg = msg

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}
