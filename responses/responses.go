package responses

import "github.com/gin-gonic/gin"

func RespondError(c *gin.Context, err error) {
	c.JSON(CODE_BAD_REQUEST, err)
}

func RespondAccessNotGranted(c *gin.Context, err error) {
	c.JSON(CODE_ACCESS_NOT_GRANTED, err)
}

func RespondNoHandler(c *gin.Context) {
	c.JSON(CODE_BAD_REQUEST, ErrNoHandler())
}

func RespondObject(c *gin.Context, o interface{}) {
	c.JSON(CODE_SUCCESS, o)
}
