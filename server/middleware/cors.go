package middleware

import "net/http"

func Cors(w http.ResponseWriter) {
	w.Header().Set("access-control-allow-origin", "*")
	w.Header().Set("access-control-allow-methods", "GET, POST, PATCH, DELETE")
	w.Header().Set("access-control-allow-headers", "accept, content-type")
	w.Header().Set("content-type", "application/json; charset=UTF-8")
}
