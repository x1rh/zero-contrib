package middleware

import (
	"net/http"
)

func ForwardHeader() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			//r.Header.Set("Grpc-Metadata-uid", uid)
			r.Header.Set("name", "lrh")
			next(w, r)
		}
	}
}
