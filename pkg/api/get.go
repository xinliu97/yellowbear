package api

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"fmt"

	"yellowbear/pkg/schema"
	"yellowbear/pkg"
	"yellowbear/pkg/utils"
)

const (
	RANKSTARTER = 1
	dbName = "yellowbear"
	collName = "quizzes"
)

func ListAllQuizzes(mc *pkg.MongoDBClient) gin.HandlerFunc {
	return func(c *gin.Context){
		coll := mc.GetCollection(dbName, collName)
		emptyFilter := bson.D{{}}
		quizList := make(map[string]string)

		cursor, err := coll.Find(context.TODO(), emptyFilter)
		if err != nil {
			fmt.Println("[ListAllQuizzes]", err)
		}
		defer cursor.Close(context.TODO())
		
		for cursor.Next(context.TODO()) {
			var pq schema.PopularityCollRead
			err = cursor.Decode(&pq)
			if err != nil {
				fmt.Println("[ListAllQuizzes]", err)
			}
			quizList[pq.ID.String()] = pq.Question
		}

		utils.RespOkWithBody(c, quizList)
	}
}