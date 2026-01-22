package exceptions

import (
	"golang-restful-api/helpers"
	"golang-restful-api/models/resources"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func PanicHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	var response resources.WebResponse

	if notFoundException, ok := err.(NotFoundException); ok {
		writer.WriteHeader(http.StatusNotFound)
		response = resources.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   notFoundException.Error,
		}

		helpers.WriteToResponseBody(writer, response)
		return
	}

	if validationException, ok := err.(validator.ValidationErrors); ok {
		writer.WriteHeader(http.StatusBadRequest)
		response = resources.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   validationException.Error(),
		}

		helpers.WriteToResponseBody(writer, response)
		return
	}

	if unauthorizedException, ok := err.(UnauthorizedException); ok {
		writer.WriteHeader(http.StatusUnauthorized)
		response := resources.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   unauthorizedException.Error,
		}

		helpers.WriteToResponseBody(writer, response)
		return
	}

	writer.WriteHeader(http.StatusInternalServerError)
	response = resources.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}
	helpers.WriteToResponseBody(writer, response)
}
