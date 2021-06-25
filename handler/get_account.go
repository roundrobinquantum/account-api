package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/roundrobinquantum/account-api/response"
	apiClient "github.com/roundrobinquantum/api-client/client"
	"net/http"
	"strconv"
)

const getPath = "%s/accounts/%d"

var GetById = func(clt *apiClient.Client, address string) func(c *gin.Context) {
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

		uri := fmt.Sprintf(getPath, address, accountId)
		req := apiClient.Get(uri).Build()
		statusCode, respBody, err := clt.End(req)

		if err != nil || isSuccessStatusCode(statusCode) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var resp response.AccountType
		err = json.Unmarshal(respBody, &resp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
