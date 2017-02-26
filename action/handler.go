package action

import (
	"github.com/gin-gonic/gin"
	"github.com/hippoai/raappd/responses"
)

type Handler func(c *gin.Context, payloadItf interface{}, claimsItf interface{})

// GetHandler returns the augmented handler for this action
func (a *Action) GetHandler() func(c *gin.Context) {

	if a.IsPrivate() {
		return a.GetPrivateHandler()
	}

	return a.GetPublicHandler()

}

// GetPrivateHandler handles authorization
func (a *Action) GetPrivateHandler() func(c *gin.Context) {

	return func(c *gin.Context) {

		claimsJWTMapClaims, err := ExtractClaims(c)
		if err != nil {
			responses.RespondAccessNotGranted(c, ErrUnparsableJWT())
			return
		}
		claimsItf, err := a.ClaimsExtractor(claimsJWTMapClaims)
		if err != nil {
			responses.RespondAccessNotGranted(c, ErrUnparsableJWT())
			return
		}

		// 2 - Parse the payload and check the type
		payloadItf, err := a.PayloadParser(c)
		if err != nil {
			responses.RespondError(c, err)
			return
		}

		// 3 - Check the authorization
		errAuthorization := a.AuthValidator(c, payloadItf, claimsItf)
		if errAuthorization != nil {
			responses.RespondError(c, errAuthorization)
			return
		}

		// All verifications have passed - Handle the query
		a.Handler(c, payloadItf, claimsItf)

	}

}

// GetPrivateHandler parses the payload
func (a *Action) GetPublicHandler() func(c *gin.Context) {

	return func(c *gin.Context) {
		payloadItf, err := a.PayloadParser(c)
		if err != nil {
			responses.RespondError(c, err)
			return
		}

		a.Handler(c, payloadItf, nil)
	}

}
