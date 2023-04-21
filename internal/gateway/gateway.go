package gateway

import (
	"api-gateway/internal/cfg"
	"api-gateway/internal/gw_errors"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Forward interface {
	// GetPrefix @description uri转换网关名称
	GetPrefix(uri string) (prefix string)
	// GetRouters @description 根据网关短名称转换路由表
	GetRouters(prefix string) (routers []string, err error)
	// GetProxy @description 计算网关代理地址
	GetProxy(prefix string, routers []string) (proxy string)
	// GetRealUrl @description 获取真实地址
	GetRealUrl(url string, prefix string, proxy string) (realUrl string)
	// ReverseProxy @description 转发请求
	ReverseProxy(addr string)
}

type SimpleGW struct {
	W http.ResponseWriter
	R *http.Request
}

func (SimpleGW) GetPrefix(uri string) (prefix string) {
	pathSplit := strings.Split(uri, "/")
	prefix = pathSplit[1]
	return
}

func (SimpleGW) GetRouters(prefix string) (routers []string, err error) {
	proxyCfg := cfg.LoadCfg()
	var exists bool
	if routers, exists = proxyCfg.Proxy[prefix]; !exists {
		err = gw_errors.ErrEmptyRouters{}.Error()
		return
	}
	return
}

func (SimpleGW) GetProxy(prefix string, routers []string) (proxy string) {
	return routers[0]
}

func (SimpleGW) GetRealUrl(uri string, prefix string, proxy string) (realUrl string) {
	return proxy + uri
}

func (sgw SimpleGW) ReverseProxy(addr string) {
	remote, err := url.Parse(addr)
	if err != nil {
		panic(err)
	}
	log.Printf("ReverseProxy.Log.Addr: %s", addr)

	log.Printf("ReverseProxy.Log.remote.Scheme: %s", remote.Scheme)
	log.Printf("ReverseProxy.Log.remote.Host: %s", remote.Host)
	sgw.R.URL.Host = remote.Host
	sgw.R.URL.Scheme = remote.Scheme
	sgw.R.Header.Set("X-Forwarded-Host", sgw.R.Header.Get("Host"))
	sgw.R.Host = remote.Host
	httputil.NewSingleHostReverseProxy(remote).ServeHTTP(sgw.W, sgw.R)
}
