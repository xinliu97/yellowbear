package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"fmt"
)

type answer struct {
	Name string `json:"name" binding:"required"`
}

func HdlInputs(client *mongo.Client) gin.HandlerFunc {
	fmt.Println("entered HdlInputs.")
    return func(c *gin.Context){
		var rcvdjson answer
		// bind JSON and check binding requirement
		if err := c.ShouldBindJSON(&rcvdjson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		// assign db name
		dbName := "yellowbear"
		// assign collection
		institutionsColl := client.Database(dbName).Collection("institutions")
	
		// assign search condition
		filter := bson.M{"name": rcvdjson.Name}
		// search by condition
		var searchResult bson.M  // TODO 这里如果改成Institution, 仍然能查到, 但是返回的doc其字段均为空值(除了filter指定的字段)
		err := institutionsColl.FindOne(context.TODO(), filter).Decode(&searchResult)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				fmt.Printf("No document was found with the title %s\n", rcvdjson.Name)
				c.JSON(http.StatusNotFound, gin.H{"error": "No document found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
            return
		}
		// add vote count
		update := bson.D{{Key: "$set", Value: bson.D{{Key: "voteCount", Value: searchResult["voteCount"].(int32)+1}}}}
		updateResult, err := institutionsColl.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vote."})
			return
		}
		fmt.Printf("update coll: %+v\n", updateResult)
		// encode searchResult bson into json and output it
		searchResult["voteCount"] = searchResult["voteCount"].(int32)+1
		responseJson, err := json.MarshalIndent(searchResult, "", "    ")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode data"})
			return
		}
		fmt.Printf("Received POST request, response: %s\n", responseJson)
	
		// OK response
		c.Header("Content-Type", "application/json")
		c.Data(http.StatusOK, "application/json", responseJson)
	}
}