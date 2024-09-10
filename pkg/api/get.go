package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"fmt"
	"github.com/elliotchance/orderedmap/v2"

	"yellowbear/pkg/schema"
)

const RANKSTARTER = 1

type submittedNames struct {
	Names []string `json:"names" binding:"required"`
}

func sortVotes(coll *mongo.Collection) (*orderedmap.OrderedMap[schema.Institution, int], error) {
	ctx := context.TODO()

	findOptions := options.Find().SetSort(bson.D{{"voteCount", -1}})
	cursor, err := coll.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		fmt.Println("%s", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var instToRank = orderedmap.NewOrderedMap[schema.Institution, int]()
	rk := RANKSTARTER
	for cursor.Next(ctx) {
		var result schema.Institution
		if err = cursor.Decode(&result); err != nil {
			return nil, err
		}
		instToRank.Set(result, rk)
		rk++
	}

	return instToRank, nil
}

func getWholeRank() {
	for el := m.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Key, el.Value)
	}
}

func getTargetedRank() {

}

func DisplayResult(client *mongo.Client) gin.HandlerFunc {
	fmt.Println("entered DisplayResult.")
	return func(c *gin.Context){
		var rcvdjson submittedNames
		// bind JSON and check binding requirement
		if err := c.ShouldBindJSON(&rcvdjson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		// assign names
		dbName := "yellowbear"
		collName := "institutions"
		// assign collection
		coll := client.Database(dbName).Collection(collName)
		// sort
		instToRank, err := sortVotes(coll)
		// construct 1st part
		wholeRank := 
		// construct 2st part




		// assign search condition
		filter := bson.M{"name": rcvdjson.Name}
		// search by condition
		var searchResult bson.M  // TODO 这里如果改成Institution, 仍然能查到, 但是返回的doc其字段均为空值(除了filter指定的字段)
		err := coll.FindOne(context.TODO(), filter).Decode(&searchResult)
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
		updateResult, err := coll.UpdateOne(context.TODO(), filter, update)
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