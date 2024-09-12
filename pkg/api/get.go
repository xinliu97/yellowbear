package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sort"

	"yellowbear/pkg"
	"yellowbear/pkg/schema"
	"yellowbear/pkg/utils"
)

const (
	RANKSTARTER = 1
	dbName      = "yellowbear"
	collName    = "quizzes"
)

func ListAllQuizzes(mc *pkg.MongoDBClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		coll := mc.GetCollection(dbName, collName)
		emptyFilter := bson.D{{}}
		quizList := make(map[string]string)

		cursor, err := coll.Find(context.TODO(), emptyFilter)
		if err != nil {
			fmt.Println("[ListAllQuizzes]", err)
		}
		defer func(cursor *mongo.Cursor, ctx context.Context) {
			err := cursor.Close(ctx)
			if err != nil {
				fmt.Println("[ListAllQuizzes]", err)
			}
		}(cursor, context.TODO())

		for cursor.Next(context.TODO()) {
			var pq schema.PopularityCollRead
			err = cursor.Decode(&pq)
			if err != nil {
				fmt.Println("[ListAllQuizzes]", err)
			}
			quizList[pq.ID.Hex()] = pq.Question
		}

		err = utils.RespOkWithBody(c, quizList)
		if err != nil {
			fmt.Println("[ListAllQuizzes]", err)
			return
		}
	}
}

type HeatResult struct {
	QuizId   string `json:"quiz_id"`
	Question string `json:"question"`
	Heat     int    `json:"heat"`
}

func calcHeat(pq schema.PopularityCollRead) int {
	return pq.ParticipantCnt
}

type ByHeat []HeatResult

func (a ByHeat) Len() int {
	return len(a)
}

// Less Descending order
func (a ByHeat) Less(i, j int) bool {
	return a[i].Heat > a[j].Heat // from large to small
}

func (a ByHeat) Swap(i, j int) {
	a[i].QuizId, a[j].QuizId = a[j].QuizId, a[i].QuizId
	a[i].Question, a[j].Question = a[j].Question, a[i].Question
	a[i].Heat, a[j].Heat = a[j].Heat, a[i].Heat
}

func QuizzesInHeatOrder(mc *pkg.MongoDBClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		coll := mc.GetCollection(dbName, collName)
		emptyFilter := bson.D{{}}
		var orderedQuizzes []HeatResult

		cursor, err := coll.Find(context.TODO(), emptyFilter)
		if err != nil {
			fmt.Println("[ListAllQuizzes]", err)
		}
		defer func(cursor *mongo.Cursor, ctx context.Context) {
			err := cursor.Close(ctx)
			if err != nil {
				fmt.Println("[ListAllQuizzes]", err)
			}
		}(cursor, context.TODO())

		for cursor.Next(context.TODO()) {
			var pq schema.PopularityCollRead
			err = cursor.Decode(&pq)
			if err != nil {
				fmt.Println("[ListAllQuizzes]", err)
			}
			oneHeat := HeatResult{
				QuizId:   pq.ID.Hex(),
				Question: pq.Question,
				Heat:     calcHeat(pq),
			}
			orderedQuizzes = append(orderedQuizzes, oneHeat)
		}
		sort.Sort(ByHeat(orderedQuizzes))

		err = utils.RespOkWithBody(c, orderedQuizzes)
		fmt.Println(orderedQuizzes)
		if err != nil {
			fmt.Println("[ListAllQuizzes]", err)
			return
		}
	}
}
