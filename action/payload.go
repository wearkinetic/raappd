package action

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/hippoai/goutil"
	"github.com/wearkinetic/beasag/berrors"
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
	defaultPayloadValue := reflect.ValueOf(defaultPayloadItf).Elem()
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

// PayloadParser
func (a *Action) ParsePayload(c *gin.Context) (interface{}, error) {

	// If get
	if c.Request.Method == "GET" {
		return a.ParseQueryParameters(c)
	}

	return a.ParseBody(c)

}

func (a *Action) ParseBody(c *gin.Context) (interface{}, error) {
	defaultPayloadItf := a.GetDefaultPayload()

	err := c.BindJSON(defaultPayloadItf)
	if err != nil {
		return nil, berrors.ErrWrongPayload(defaultPayloadItf)
	}
	return defaultPayloadItf, nil
}

// ParseQueryParameters
func (a *Action) ParseQueryParameters(c *gin.Context) (interface{}, error) {

	queryParameters := c.Request.URL.Query()

	defaultPayloadItf := a.GetDefaultPayload()

	defaultPayloadValue := reflect.ValueOf(defaultPayloadItf).Elem()
	defaultPayloadType := defaultPayloadValue.Type()
	queryParametersAsMapStrItf := map[string]interface{}{}
	for i := 0; i < defaultPayloadValue.NumField(); i++ {
		field := defaultPayloadType.Field(i).Tag.Get("json")
		required := defaultPayloadType.Field(i).Tag.Get("binding") == "required"
		parameterValues := queryParameters[field]
		if required && ((len(parameterValues) == 0) || (parameterValues[0] == "")) {
			return nil, berrors.ErrWrongPayload(defaultPayloadItf)
		}

		// Populate the query parameters
		if len(queryParameters[field]) == 1 {
			queryParametersAsMapStrItf[field] = parameterValues[0]
		} else {
			queryParametersAsMapStrItf[field] = parameterValues
		}
	}

	err := goutil.JsonRestruct(queryParametersAsMapStrItf, defaultPayloadItf)

	if err != nil {
		return nil, berrors.ErrWrongPayload(defaultPayloadItf)
	}

	return defaultPayloadItf, nil

}
