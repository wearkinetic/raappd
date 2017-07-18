package action

import (
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/wearkinetic/gin"
)

type ClaimsExtractor func(claimsAsJWTMapClaims jwt.MapClaims) (claimsItf interface{}, err error)

// ExtractClaims into jwt.Claims format
func ExtractClaims(c *gin.Context) (jwt.MapClaims, error) {

	// 1 - Extract JWT token from header
	tokenString := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", 1)
	token, err := ParseJWT(tokenString)
	if err != nil {
		return nil, err
	}

	// 2 - Now add the custom logic for this endpoint
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, ErrUnparsableJWT()
	}

	return claims, nil

}

// ParseJWT wraps jwt package with our secret key
func ParseJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return jwt_secret, nil
	})
}
