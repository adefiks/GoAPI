package middleware

import (
	"net/http"
	"strings"
)

// StripSlashes is a middleware function that removes trailing slashes from the URL.
func StripSlashes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
