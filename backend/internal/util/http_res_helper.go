package util

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Body interface{} `json:"body"`
	Msg  string      `json:"message"`
}

func HandleHTTPResponse(w http.ResponseWriter, msg string, body ...interface{}) {
	var res Response

	res.Body = body
	res.Msg = msg

	json.NewEncoder(w).Encode(res)
}
