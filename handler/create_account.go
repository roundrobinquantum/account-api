package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/roundrobinquantum/account-api/response"
	apiClient "github.com/roundrobinquantum/api-client/client"
	"net/http"
)

type Account struct {
	Id          int64
	TotalAmount float64
	Name        string
}

const createPath = "%s/accounts/"

var CreateAccount = func(clt *apiClient.Client, address string) func(c *gin.Context) {
	return func(c *gin.Context) {
		var accountModel *Account
		err := c.ShouldBindJSON(&accountModel)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not bind request"})
		}

		uri:=fmt.Sprintf(createPath,address)
		req := apiClient.Post(uri, accountModel).Build()
		statusCode, respBody, err := clt.End(req)

		if err != nil || isSuccessStatusCode(statusCode) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not parse accountId!"})
			return
		}

		var resp response.AccountType
		err = json.Unmarshal(respBody, &resp)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not parse accountId!"})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}
