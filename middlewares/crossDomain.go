package middlewares

import (
	"github.com/gin-gonic/gin"
)

func CrossHandler(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
	c.Writer.Header().Set("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Writer.Header().Set("Content-Type", "application/json")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}
