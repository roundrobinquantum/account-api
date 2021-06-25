package handler

import (
	"github.com/gin-gonic/gin"
	clt "github.com/roundrobinquantum/api-client/client"
)

func InitializeRoutes(e *gin.Engine, c *clt.Client, uri string) {
	e.GET("accounts/id/:accountId", GetById(c, uri))
	e.POST("accounts/id:accountId", CreateAccount(c, uri))
	e.DELETE("accounts/id:accountId", DeleteById(c, uri))
}

