package handler

import (

	"github.com/gin-gonic/gin"
	clt "github.com/roundrobinquantum/api-client/client"
)

func InitializeRoutes(e *gin.Engine,c  *clt.Client,uri string) {
	e.GET("accounts/accountId/:accountId", GetById(c,uri))
}
