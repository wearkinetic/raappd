package action

import "github.com/wearkinetic/gin"

type AuthValidator func(c *gin.Context, payloadItf interface{}, claimsItf interface{}) error
