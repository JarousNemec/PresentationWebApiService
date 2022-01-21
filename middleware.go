package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	user := "superuser"
	pass := "resurepus"
	return func(c *gin.Context) {
		username := c.GetHeader("User-Name")
		password := c.GetHeader("User-Pass")
		if username == user && password == pass {
			c.Status(http.StatusOK)
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
