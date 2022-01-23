package main

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

var authorization map[string]interface{}

func AuthMiddleWare() gin.HandlerFunc {
	user := authorization["conn_user"]
	pass := authorization["conn_pass"]
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

func loadAuthorization() {
	client, err := storage.NewBasicClient(account, key)
	if err != nil {
		fmt.Printf("%s: \n", err)
	}
	tableCli = client.GetTableService()
	table := tableCli.GetTableReference(authtablename)
	entities, err := table.QueryEntities(30, fullmetadata, nil)
	authorization = entities.Entities[0].Properties
}
