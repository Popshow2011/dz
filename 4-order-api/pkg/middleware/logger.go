package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logrus.WithFields(logrus.Fields{
			"Path":   r.URL.Path,
			"Method": r.Method,
		}).Info("A path and method")

		next.ServeHTTP(w, r)
	})

}
