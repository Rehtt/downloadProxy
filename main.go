/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/3/23 20:04
 */

package main

import (
	"net/http"
	"strings"
)

func main() {
	//	http://127.0.0.1:8080/https/baidu.com
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		scheme, url, get := strings.Cut(request.RequestURI[1:], "/")
		if !get {
			writer.WriteHeader(400)
			return
		}
		proxy, _ := NewProxy(scheme, url)
		proxy.ServeHTTP(writer, request)
	})
	http.ListenAndServe(":8080", nil)
}
