package helper

import "github.com/go-playground/validator/v10"

type response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	jsonresponse := response{
		Meta: meta,
		Data: data,
	}
	return jsonresponse
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
