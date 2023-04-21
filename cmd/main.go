package main

import (
	"api-gateway/internal/auth"
	"api-gateway/internal/cache"
	_ "api-gateway/internal/cfg"
	"api-gateway/internal/serv"
	"net/http"
)

func main() {
	cache.InitPoolRedis()
	servHandler := http.HandlerFunc(serv.Handler)
	http.Handle("/", auth.MiddleWire(servHandler))
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err.Error())
	}
}
