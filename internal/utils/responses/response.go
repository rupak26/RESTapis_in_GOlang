package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	StatusOk = "OK" 
	StatusError = "ERROR"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func WriteJson(w http.ResponseWriter , status int , data interface{}) error {
    w.Header().Set("Content-Type" , "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error: err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMgs []string 

	for _,err := range errs {
		if err.ActualTag() == "required" {
			errMgs = append(errMgs, fmt.Sprintf("fields %s is a required filed", err.Field()))
		} else {
			errMgs = append(errMgs, fmt.Sprintf("fields %s is a Invalid filed", err.Field()))
		}
	}

	return Response{
		Status: StatusError, 
		Error: strings.Join(errMgs,", "),
	}
}