package resource

import "reflect"

type ResponseFormat struct {
	Description string `json:"description"`
	Parameters  []ParameterDescription
}

// NewResponseFormat creates one from a description and defaultResponse
func NewResponseFormat(description string, defaultResponseItf interface{}) *ResponseFormat {
	return &ResponseFormat{
		Description: description,
		Parameters:  ExtractParametersDescription(defaultResponseItf),
	}
}

type ParameterDescription struct {
	Name        string `json:"name"`
	DataType    string `json:"datatype"`
	Description string `json:"description"`
}

// ExtractParametersDescription from an object
func ExtractParametersDescription(defaultResponseItf interface{}) []ParameterDescription {
	pd := []ParameterDescription{}

	// Now verify all the fields are there
	defaultResponseValue := reflect.ValueOf(defaultResponseItf).Elem()
	defaultResponseType := defaultResponseValue.Type()
	for i := 0; i < defaultResponseValue.NumField(); i++ {
		field := defaultResponseType.Field(i).Tag.Get("json")
		description := defaultResponseType.Field(i).Tag.Get("description")
		dt := defaultResponseType.Field(i).Type.String()

		// Extract the information from struct definition
		pd = append(pd, ParameterDescription{
			Name:        field,
			DataType:    dt,
			Description: description,
		})

	}

	return pd
}

func (r *Resource) AddResponseFormat(description string, defaultResponseItf interface{}) {
	r.ResponseFormats = append(r.ResponseFormats,
		NewResponseFormat(description, defaultResponseItf),
	)
}
