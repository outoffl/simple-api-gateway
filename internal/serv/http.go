package serv

import (
	"api-gateway/internal/gateway"
	"api-gateway/internal/gw_errors"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	sgw := gateway.SimpleGW{
		W: w,
		R: r,
	}
	if err := forwardHandler(w, r, sgw); err != nil {
		w.WriteHeader(404)
		return
	}
}

func forwardHandler(w http.ResponseWriter, r *http.Request, gf gateway.Forward) error {
	prefixUri := gf.GetPrefix(r.RequestURI)
	if prefixUri == "" {
		return gw_errors.ErrUriNull{}.Error()
	}
	routers, err := gf.GetRouters(prefixUri)
	if err != nil {
		return err
	}
	proxy := gf.GetProxy(prefixUri, routers)
	realUrl := gf.GetRealUrl(r.RequestURI, prefixUri, proxy)
	gf.ReverseProxy(realUrl)
	return nil
}
