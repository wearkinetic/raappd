package action

import "github.com/gin-gonic/gin"

type PayloadParser func(c *gin.Context) (payloadItf interface{}, err error)
