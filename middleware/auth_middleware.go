package middleware

import (
	"net/http"

	"github.com/abdultalif/restful-api/helper"
	"github.com/abdultalif/restful-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware  {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-KEY") {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		WebResponseFailed := web.WebResponseFailed{
			Success: false,
			Code:    http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Error:   "Email or password incorrect",
		}

		helper.WriteToResponseBody(writer, WebResponseFailed)
	}
}
