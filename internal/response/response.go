package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Data interface{} `json:"data"`
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
}

func (r Response) Unauthorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	return
}
