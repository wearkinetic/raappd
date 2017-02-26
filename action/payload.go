package action

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type PayloadParser func(c *gin.Context) (payloadItf interface{}, err error)
type PayloadDescriptor func(payloadItf interface{}) map[string]*ParameterDescription

type ParameterDescription struct {
	Name        string `json:"name"`
	DataType    string `json:"datatype"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

type GetDefaultPayload func() interface{}

func (a *Action) MakePayloadDescription() map[string]*ParameterDescription {
	defaultPayloadItf := a.GetDefaultPayload()
	payloadDescription := map[string]*ParameterDescription{}

	// Now verify all the fields are there
	defaultPayloadValue := reflect.ValueOf(defaultPayloadItf)
	defaultPayloadType := defaultPayloadValue.Type()
	for i := 0; i < defaultPayloadValue.NumField(); i++ {
		field := defaultPayloadType.Field(i).Tag.Get("json")
		description := defaultPayloadType.Field(i).Tag.Get("description")
		required := defaultPayloadType.Field(i).Tag.Get("binding") == "required"

		if description == "" {
			FatalIncompleteParametersDescription(a.Resource, a.Verb, a.Name, field)
		}

		// Extract the information from struct definition
		payloadDescription[field] = &ParameterDescription{
			Name:        field,
			DataType:    defaultPayloadType.Field(i).Type.String(),
			Description: description,
			Required:    required,
		}

	}

	return payloadDescription
}
