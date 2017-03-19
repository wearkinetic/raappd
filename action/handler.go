package action

import (
	"github.com/gin-gonic/gin"
	"github.com/hippoai/raappd/responses"
)

type Handler func(c *gin.Context, payloadItf interface{}, claimsItf interface{})

// GetHandler returns the augmented handler for this action
func (a *Action) GetHandler(c *gin.Context, payloadItf interface{}) {

	if a.IsPrivate() {
		a.GetPrivateHandler(c, payloadItf)
		return
	}

	a.GetPublicHandler(c, payloadItf)

}

// GetPrivateHandler handles authorization
func (a *Action) GetPrivateHandler(c *gin.Context, payloadItf interface{}) {

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

	// 3 - Check the authorization
	errAuthorization := a.AuthValidator(c, payloadItf, claimsItf)
	if errAuthorization != nil {
		responses.RespondError(c, errAuthorization)
		return
	}

	// All verifications have passed - Handle the query
	a.Handler(c, payloadItf, claimsItf)

}

// GetPrivateHandler parses the payload
func (a *Action) GetPublicHandler(c *gin.Context, payloadItf interface{}) {

	a.Handler(c, payloadItf, nil)

}
