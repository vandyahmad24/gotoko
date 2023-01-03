package helper

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type ErrorResponseUpdate struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Error   []ErrorMetaUpdate `json:"error"`
}

type ErrorMetaUpdate struct {
	Message string                `json:"message"`
	Path    []string              `json:"path"`
	Type    string                `json:"type"`
	Context ContextResponseUpdate `json:"context"`
}

type ContextResponseUpdate struct {
	PeersWithLabels []string `json:"peersWithLabels"`
	Peers           []string `json:"peers"`
	Label           string   `json:"label"`
	Value           struct{} `json:"value"`
}

func FormatErrorValidationUpdate(err error) ErrorResponseUpdate {
	var tempError string
	var messageBody string
	var pers []string
	var errorResponse ErrorResponseUpdate
	var errMeta []ErrorMetaUpdate
	var tempMessageBody []string
	for _, v := range err.(validator.ValidationErrors) {
		lower := ToSnakeCase(v.Field())
		switch v.Tag() {
		case "required":
			messageBody = lower
		default:
			messageBody = lower
		}
		persTemp := ToSnakeCase(v.Field())
		pers = append(pers, persTemp)

		context := ContextResponseUpdate{
			Label: "value",
		}
		tempErrorMeta := ErrorMetaUpdate{
			Message: tempError,
			Path:    []string{},
			Type:    "object.missing",
			Context: context,
		}

		errMeta = append(errMeta, tempErrorMeta)
		tempMessageBody = append(tempMessageBody, messageBody)

		errorResponse.Status = false

	}

	justString := strings.Join(tempMessageBody, ", ")

	errorResponse.Message = `body ValidationError:  "value" must contain at least one of [` + justString + `]`
	errorResponse.Error = errMeta
	errorResponse.Error[0].Context.Peers = pers
	errorResponse.Error[0].Context.PeersWithLabels = pers
	errorResponse.Error[0].Message = `"value" must contain at least one of [` + justString + `]`
	return errorResponse
}
