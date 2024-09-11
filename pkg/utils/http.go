package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespOkWithBody(c *gin.Context, bodyStruct interface{}) error {
	bodyJson, err := json.MarshalIndent(bodyStruct, "", "    ")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode data"})
			return err
		}
	
		c.Header("Content-Type", "application/json")
		c.Data(http.StatusOK, "application/json", bodyJson)

		return nil
}

func ReadPostBody(c *gin.Context, bodyStruct interface{}) error {
	if err := c.ShouldBindJSON(bodyStruct); err != nil {  
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	return nil
}