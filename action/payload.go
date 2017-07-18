package action

import (
	"reflect"

	"github.com/wearkinetic/gin"
)

type PayloadParser func(c *gin.Context) (payloadItf interface{}, err error)
type PayloadDescriptor func(payloadItf interface{}) map[string]*ParameterDescription

type PayloadDescription struct {
	Parameters map[string]*ParameterDescription `json:"parameters"`
	NRequired  int                              `json:"nOptional"`
	NOptional  int                              `json:"nRequired"`
}

func NewPayloadDescription() *PayloadDescription {
	return &PayloadDescription{
		Parameters: map[string]*ParameterDescription{},
	}
}

func (payloadDescription *PayloadDescription) Add(name, dataType, description string, required bool) *PayloadDescription {
	payloadDescription.Parameters[name] = NewParameterDescription(name, dataType, description, required)
	if required {
		payloadDescription.NRequired += 1
	} else {
		payloadDescription.NOptional += 1
	}
	return payloadDescription
}

func (payloadDescription *PayloadDescription) IsRequired(key string) (bool, error) {
	param, ok := payloadDescription.Parameters[key]
	if !ok {
		return false, ErrParameterNotRequiredNorOptional(key)
	}
	return param.Required, nil
}

type ParameterDescription struct {
	Name        string `json:"name"`
	DataType    string `json:"datatype"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

// NewParameterDescription instanciates
func NewParameterDescription(name, dataType, description string, required bool) *ParameterDescription {
	return &ParameterDescription{
		Name:        name,
		DataType:    dataType,
		Description: description,
		Required:    required,
	}
}

type GetDefaultPayload func() interface{}

// ExtractExpectedPayload
func ExtractExpectedPayload(defaultPayloadItf interface{}) (*PayloadDescription, error) {

	payloadDescription := NewPayloadDescription()

	// Now verify all the fields are there
	defaultPayloadValue := reflect.ValueOf(defaultPayloadItf).Elem()
	defaultPayloadType := defaultPayloadValue.Type()
	for i := 0; i < defaultPayloadValue.NumField(); i++ {
		field := defaultPayloadType.Field(i).Tag.Get("json")
		description := defaultPayloadType.Field(i).Tag.Get("description")
		required := defaultPayloadType.Field(i).Tag.Get("binding") == "required"

		if description == "" {
			return nil, ErrIncompleteParametersDescription(field)
		}

		// Extract the information from struct definition
		payloadDescription.Add(
			field,
			defaultPayloadType.Field(i).Type.String(),
			description, required,
		)

	}

	return payloadDescription, nil
}
