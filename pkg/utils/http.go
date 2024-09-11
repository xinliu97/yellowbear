package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespOkWithBody(c *gin.Context, body interface{}) {
	bodyJson, err := json.MarshalIndent(body, "", "    ")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode data"})
			return
		}
	
		c.Header("Content-Type", "application/json")
		c.Data(http.StatusOK, "application/json", bodyJson)
}