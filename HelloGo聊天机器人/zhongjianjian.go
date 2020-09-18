package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/go/src/pkg/net/http"
)

func AuthMiddleWare() gin.HandlerFunc{
	return  func (c *gin.Context) {
		token := c.Request.Header.Get("Authorizaation")
		authorized := check(token)
		if authorized {
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":"Unauthorized",
		})
		c.Abort()
		return
	}
}

func main() {

	r := gin.Default()

	r.GET("/path",AuthMiddleWare(),func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"data":"ok"})
	})
}
