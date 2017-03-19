package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/hippoai/raappd/action"
	"github.com/hippoai/raappd/responses"
)

// MakeNotGetHandler
func (r *Resource) MakeNotGetHandler(a *action.Action) func(c *gin.Context) {

	return func(c *gin.Context) {

		// 0 - Parse the payload first
		payloadItf, err := ParseBody(c, a.GetDefaultPayload)

		// 1 - If there is an error return it
		if err != nil {
			responses.RespondError(c, err)
			return
		}

		// 2 - If it worked, server it and let the action do the rest
		a.GetHandler(c, payloadItf)

	}

}

// MakeGetHandler
func (r *Resource) MakeGetHandler() func(c *gin.Context) {

	return func(c *gin.Context) {

		getParametersAsMap := ParseGETParameters(c)

		bestGet, err := r.FindGetAction(getParametersAsMap)
		if err != nil {
			responses.RespondError(c, err)
			return
		}

		// We found the best get parameter to satisfy this
		// Let's parse the payload now
		payloadItf, err := ReformatGetParameters(getParametersAsMap, bestGet.GetDefaultPayload)

		if err != nil {
			responses.RespondError(c, err)
			return
		}

		bestGet.GetHandler(c, payloadItf)

	}

}

// FindGetAction
func (r *Resource) FindGetAction(params map[string]interface{}) (*action.Action, error) {

	bestGets := []*action.Action{}
	maxNRequired := -1000
	maxNOptional := -1000
	minNAdditional := 1000

	// Loop through all the gest options
	for _, get := range r.Gets {

		// Count the number of required, optional and additional parameters we get
		nRequiredGot := 0
		nOptionalGot := 0
		nAdditionalGot := 0

		// Loop through the GET parameters
		for k, _ := range params {
			isRequired, err := get.ExpectedPayloadDescription.IsRequired(k)
			if err != nil {
				nAdditionalGot += 1
				continue
			}
			if isRequired {
				nRequiredGot += 1
			} else {
				nOptionalGot += 1
			}
		}

		// If you've got the right number of required, you are a potential candidate
		if nRequiredGot == get.ExpectedPayloadDescription.NRequired {

			switch isBetter(
				nRequiredGot, nOptionalGot, nAdditionalGot,
				maxNRequired, maxNOptional, minNAdditional,
			) {
			case GET_STRICTLY_BETTER:
				bestGets = []*action.Action{get}
			case GET_EQUAL:
				bestGets = append(bestGets, get)
			}

			// Update the best get estimate
			// Update the max and min values
			maxNRequired = intmax(maxNRequired, nRequiredGot)
			maxNOptional = intmax(maxNOptional, nOptionalGot)
			minNAdditional = intmin(minNAdditional, nAdditionalGot)

		}

	}

	if len(bestGets) == 0 {
		return nil, ErrNoGetActionFound()
	}

	if len(bestGets) > 1 {
		getsNames := []string{}
		for _, get := range bestGets {
			getsNames = append(getsNames, get.Name)
		}
		return nil, ErrMultipleGetsCanSatisfyThisQuery(getsNames)
	}

	// We found just one, hooray
	return bestGets[0], nil

}

func isBetter(
	nRequiredGot, nOptionalGot, nAdditionalGot,
	maxNRequired, maxNOptional, minNAdditional int,
) int {

	if nRequiredGot > maxNRequired {
		return GET_STRICTLY_BETTER
	}

	if nRequiredGot < maxNRequired {
		return GET_STRICTLY_WORSE
	}

	// Situation where nRequiredGot == maxNRequired
	// judge by highest number of optionals now
	if nOptionalGot > maxNOptional {
		return GET_STRICTLY_BETTER
	}

	if nOptionalGot < maxNOptional {
		return GET_STRICTLY_WORSE
	}

	// Situation where nOptionalGot == maxNOptional
	// judge by lowest number of additionals now
	if nAdditionalGot < minNAdditional {
		return GET_STRICTLY_BETTER
	}

	// By CONVENTION, if not strictly better fit, return false
	return GET_EQUAL

}

func intmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func intmin(a, b int) int {
	return -intmax(-a, -b)
}
