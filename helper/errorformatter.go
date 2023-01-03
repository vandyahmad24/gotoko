package helper

import (
	"bytes"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   []ErrorMeta `json:"error"`
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

func FormatErrorValidationCreate(err error) ErrorResponse {
	var tempError string
	var messageBody string
	var path []string
	var errorResponse ErrorResponse
	var errMeta []ErrorMeta
	var tempMessageBody []string
	for _, v := range err.(validator.ValidationErrors) {
		tag := Q(v.Field())
		lower := ToSnakeCase(tag)
		switch v.Tag() {
		case "required":
			tempError = lower + `is required`
			messageBody = lower + ` is required`
		case "min":
			tempError = lower + `min ` + v.Param()
			messageBody = lower + ` min ` + v.Param()
		case "max":
			tempError = lower + `max ` + v.Param()
			messageBody = lower + ` max ` + v.Param()
		default:
			tempError = v.Error()
		}
		path = append(path, ToSnakeCase(v.Field()))
		context := ContextResponse{
			Label: ToSnakeCase(v.Field()),
			Key:   ToSnakeCase(v.Field()),
		}
		tempErrorMeta := ErrorMeta{
			Message: tempError,
			Path:    path,
			Type:    "any." + v.Tag(),
			Context: context,
		}

		errMeta = append(errMeta, tempErrorMeta)
		tempMessageBody = append(tempMessageBody, messageBody)

		errorResponse.Status = false

	}

	justString := strings.Join(tempMessageBody, ". ")

	errorResponse.Message = "body ValidationError: " + justString
	errorResponse.Error = errMeta
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
