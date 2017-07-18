package resource

import (
	"strings"

	"github.com/wearkinetic/gin"
	"github.com/hippoai/goutil"
	"github.com/hippoai/raappd/constants"
)

// ParseBody for POST, PATCH, PUT, DELETE
func ParseBody(c *gin.Context, makeDefaultPayloadItf func() interface{}) (interface{}, error) {

	defaultPayloadItf := makeDefaultPayloadItf()
	err := c.BindJSON(defaultPayloadItf)
	if err != nil {
		return nil, ErrWrongPayload(defaultPayloadItf)
	}
	return defaultPayloadItf, nil

}

// ParseGETParameters, the values will either be strings or arrays of string
func ParseGETParameters(c *gin.Context) map[string]interface{} {
	m := map[string]interface{}{}

	queryParameters := c.Request.URL.Query()
	for k, v := range queryParameters {
		if len(v) > 1 {
			m[k] = strings.Join(v, constants.SPLIT_SLICE)
			continue
		}
		if len(v) == 1 {
			m[k] = v[0]
		}
	}

	return m
}

// ReformatGetParameters
func ReformatGetParameters(m map[string]interface{}, makeDefaultPayloadItf func() interface{}) (interface{}, error) {

	defaultPayloadItf := makeDefaultPayloadItf()
	err := goutil.JsonRestruct(m, defaultPayloadItf)
	if err != nil {
		return nil, ErrWrongPayload(defaultPayloadItf)
	}
	return defaultPayloadItf, nil

}
