package middlewares

import (
	"golang-restful-api/helpers"
	"golang-restful-api/models/resources"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middlware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := "SECRET"
	authHeader := request.Header.Get("Authorization")

	if authHeader == "" {
		writer.WriteHeader(http.StatusUnauthorized)
		response := resources.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   "Missing authorization header",
		}

		helpers.WriteToResponseBody(writer, response)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		writer.WriteHeader(http.StatusUnauthorized)
		response := resources.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   "Invalid authorization format",
		}

		helpers.WriteToResponseBody(writer, response)
		return
	}

	token := parts[1]
	if token != key {
		writer.WriteHeader(http.StatusUnauthorized)
		response := resources.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   "Invalid token",
		}

		helpers.WriteToResponseBody(writer, response)
		return
	}

	middlware.Handler.ServeHTTP(writer, request)
}
