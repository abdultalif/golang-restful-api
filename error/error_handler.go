package error

import (
	"net/http"

	"github.com/abdultalif/restful-api/helper"
	"github.com/abdultalif/restful-api/model/web"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	if badRequestError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func badRequestError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		WebResponseFailed := web.WebResponseFailed{
			Success: false,
			Code:    http.StatusBadRequest,
			Status: "BAD REQUEST",
			Error:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, WebResponseFailed)
		return true
	} else {
		return false
	}
}
func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		WebResponseFailed := web.WebResponseFailed{
			Success: false,
			Code:    http.StatusNotFound,
			Status: "NOT FOUND",
			Error:   exception.Error,
		}

	helper.WriteToResponseBody(writer, WebResponseFailed)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	WebResponseFailed := web.WebResponseFailed{
		Success: false,
		Code:    http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Error:   err,
	}

	helper.WriteToResponseBody(writer, WebResponseFailed)
}
