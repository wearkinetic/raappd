package action

import "github.com/gin-gonic/gin"

type AuthValidator func(c *gin.Context, payloadItf interface{}, claimsItf interface{}) error
