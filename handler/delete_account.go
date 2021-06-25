package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	apiClient "github.com/roundrobinquantum/api-client/client"
	"net/http"
	"strconv"
)

const deletePath = "%s/accounts/%d"

var DeleteById = func(clt *apiClient.Client, address string) func(c *gin.Context) {
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

		uri := fmt.Sprintf(deletePath, address, accountId)
		req := apiClient.Delete(uri).Build()
		statusCode, _, err := clt.End(req)

		if err != nil || isSuccessStatusCode(statusCode) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}


		c.JSON(http.StatusOK,nil)
	}
}