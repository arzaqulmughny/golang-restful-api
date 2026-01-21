package helpers

import (
	"encoding/json"
	"net/http"
)

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}
