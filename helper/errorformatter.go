package helper

import (
	"bytes"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Error   ErrorMeta `json:"error"`
}

type ErrorMeta struct {
	Message string          `json:"message"`
	Path    interface{}     `json:"path"`
	Type    string          `json:"type"`
	Context ContextResponse `json:"context"`
}

type ContextResponse struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}

func FormatErrorValidation(err error) ErrorResponse {
	var tempError string
	var messageBody string
	var path []string
	var errorResponse ErrorResponse
	for _, v := range err.(validator.ValidationErrors) {
		tag := Q(v.Field())
		lower := ToSnakeCase(tag)
		switch v.Tag() {
		case "required":
			tempError = lower + `is required`
			messageBody = "body ValidationError: " + lower + ` is required`
		case "min":
			tempError = lower + `min ` + v.Param()
			messageBody = "body ValidationError: " + lower + ` min ` + v.Param()
		case "max":
			tempError = lower + `max ` + v.Param()
			messageBody = "body ValidationError: " + lower + ` max ` + v.Param()
		default:
			tempError = v.Error()
		}
		path = append(path, ToSnakeCase(v.Field()))
		context := ContextResponse{
			Label: ToSnakeCase(v.Field()),
			Key:   ToSnakeCase(v.Field()),
		}
		errorMeta := ErrorMeta{
			Message: tempError,
			Path:    path,
			Type:    "any." + v.Tag(),
			Context: context,
		}

		errorResponse.Status = false
		errorResponse.Message = messageBody
		errorResponse.Error = errorMeta
		return errorResponse

	}
	return errorResponse
}

func Q(q string) string {
	result := `"` + strings.ReplaceAll(q, `"`, `\"`) + `"`
	return result
}

func MakeFirstLowerCase(s string) string {

	if len(s) < 2 {
		return strings.ToLower(s)
	}

	bts := []byte(s)

	lc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}
