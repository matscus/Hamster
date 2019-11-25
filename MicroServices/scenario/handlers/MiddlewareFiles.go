package handlers

import "net/http"

//MiddlewareFiles - func
func MiddlewareFiles(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Keep-Alive", "300")
		w.Header().Add("Content-Disposition", "attachment")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.Header().Add("Access-Control-Max-Age", "600")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Strict-Transport-Security", "max-age=31536000")
		w.Header().Add("X-Frame-Options", "SAMEORIGIN")
		h.ServeHTTP(w, r)
	}
}
