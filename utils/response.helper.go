package utils

import "github.com/gin-gonic/gin"

func ProduceCreatedAPIResponse(c *gin.Context, response interface{}) {
	c.JSON(201, response)
}

func ProduceSuccessAPIResponse(c *gin.Context, response interface{}) {
	c.JSON(200, response)
}
