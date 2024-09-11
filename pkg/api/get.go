package api

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"fmt"
	"github.com/elliotchance/orderedmap/v2"

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