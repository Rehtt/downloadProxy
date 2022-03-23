/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/3/23 20:08
 */

package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func NewProxy(scheme, path string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(scheme + "://" + path)
	if err != nil {
		return nil, err
	}
	requestUri := strings.Split(path, "/")

	proxy := httputil.NewSingleHostReverseProxy(url)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Header.Set("X-Proxy", "Simple-Reverse-Proxy")
		req.Header.Set("Host", req.URL.Host)
		req.URL = url
		req.Host = url.Host

		if len(requestUri) > 1 {
			req.RequestURI = "/" + strings.Join(requestUri[1:], "/")
		}
		fmt.Println(req.Host)
		fmt.Printf("%+v", req)

	}

	//proxy.ModifyResponse = modifyResponse()
	proxy.ErrorHandler = errorHandler()
	return proxy, nil
}

func errorHandler() func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, req *http.Request, err error) {
		fmt.Printf("Got error while modifying response: %v \n", err)
		return
	}
}

func modifyResponse() func(*http.Response) error {
	return func(resp *http.Response) error {
		// 两个Access-Control-Allow-Origin会被浏览器报错
		resp.Header.Del("Access-Control-Allow-Origin")
		return nil
	}
}

// ProxyRequestHandler handles the http request using proxy
func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}
