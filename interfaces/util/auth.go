package util

import (
	"net/http"
)

func BasicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if token == "" {
			Error(w, 400, nil, "invalid or empty token")
			return
		}

		authOK, err := checkToken(token)
		if err != nil {
			Error(w, 500, err, "connection error with Security Center")
		}

		if authOK == false {
			Error(w, 401, nil, "Not authorized")
			return
		}

		h.ServeHTTP(w, r)
	}
}

func checkToken(token string) (bool, error) {
	return true, nil
}
