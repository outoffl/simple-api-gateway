package auth

import (
	"api-gateway/internal/response"
	"net/http"
)

func MiddleWire(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if len(token) == 0 {
			resp := response.Response{}
			resp.Unauthorized(w)
			return
		} else {
			// todo::鉴权
			if token != "debug-mod" {
				resp := response.Response{}
				resp.Unauthorized(w)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
