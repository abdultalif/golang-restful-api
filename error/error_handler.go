package error

import (
	"fmt"
	"net/http"
	"strings"

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

        errors := make(map[string]interface{})
        for _, e := range exception {
            field := e.Field()

            fieldName := strings.ToLower(field[strings.LastIndex(field, ".")+1:])
            
            if _, ok := errors[fieldName]; !ok {
                errors[fieldName] = make([]string, 0)
            }
            
            var message string
            switch e.Tag() {
            case "required":
                message = "This field is required"
            case "min":
                message = fmt.Sprintf("Minimum length is %s", e.Param())
            case "max":
                message = fmt.Sprintf("Maximum length is %s", e.Param())
            default:
                message = fmt.Sprintf("Field validation failed on '%s' tag", e.Tag())
            }
            
            errors[fieldName] = append(errors[fieldName].([]string), message)
        }

        errorResponse := web.WebResponseFailed{
            Success: false,
            Code:    http.StatusBadRequest,
            Status:  "BAD REQUEST",
            Error:   errors,
        }

		helper.WriteToResponseBody(writer, errorResponse)
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
