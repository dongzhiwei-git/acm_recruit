package middlewares

import (
	"acm_recruit/dao"
	"context"

	"github.com/gin-gonic/gin"
)

func DBMiddleware(c *gin.Context) {
	reqCtx, _ := context.WithCancel(context.Background())
	ctx := context.WithValue(c.Request.Context(), "DB", dao.GetDBInstance().DB.WithContext(reqCtx))
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
