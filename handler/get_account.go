package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/roundrobinquantum/account-api/response"
	client_for_account "github.com/roundrobinquantum/api-client/client"
	"net/http"
	"strconv"
)

var GetById = func(clt *client_for_account.Client, uri string) func(c *gin.Context) {
	return func(c *gin.Context) {
		accountId, err := strconv.Atoi(c.Param("accountId"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not parse accountId!"})
			return
		}

		if accountId < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id can not be lower than zero!"})
			return
		}

		req := client_for_account.Get(uri).Build()
		statusCode, respBody, err := clt.End(req)

		if err != nil || isSuccessStatusCode(statusCode) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not parse accountId!"}) //todo
			return
		}


		var resp response.AccountType
		err=json.Unmarshal(respBody,&resp)
		if err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not parse accountId!"}) //todo
			return
		}


		c.JSON(http.StatusOK, resp)
	}
}

func isSuccessStatusCode(statusCode int) bool {
	if int(statusCode/100) == 2 {
		return true
	}
	return false
}
