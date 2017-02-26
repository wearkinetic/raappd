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
	Mandatory   bool   `json:"mandatory"`
}

type GetDefaultPayload func() interface{}

func (a *Action) MakePayloadDescription() map[string]*ParameterDescription {
	defaultPayloadItf := a.GetDefaultPayload()
	payloadDescription := a.PayloadDescriptor(defaultPayloadItf)

	// Now verify all the fields are there
	defaultPayloadValue := reflect.ValueOf(defaultPayloadItf)
	defaultPayloadType := defaultPayloadValue.Type()
	for i := 0; i < defaultPayloadValue.NumField(); i++ {
		field := defaultPayloadType.Field(i).Tag.Get("json")
		_, exists := payloadDescription[field]
		if !exists {
			FatalIncompleteParametersDescription(a.Resource, a.Verb, a.Name, field)
		}

		// Automatically add datatype and name
		payloadDescription[field].DataType = defaultPayloadType.Field(i).Type.String()
		payloadDescription[field].Name = field
	}

	return payloadDescription
}

func NewPD(description string, mandatory bool) *ParameterDescription {
	return &ParameterDescription{
		Description: description,
		Mandatory:   mandatory,
	}
}
