package api

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"yellowbear/pkg"
	"yellowbear/pkg/schema"
	"yellowbear/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type rcvdAnswers struct {
	QuizID string   `json:"_id" binding:"required"`
	Names  []string `json:"names" binding:"required"`
}

func getDocByStringID(strID string, coll *mongo.Collection, tarDoc interface{}) error {
	objectID, err := primitive.ObjectIDFromHex(strID)
	if err != nil {
		fmt.Println("[getDocByStringID]", err)
		return err
	}

	filter := bson.D{{Key: "_id", Value: objectID}}
	err = coll.FindOne(context.TODO(), filter).Decode(tarDoc)
	if err != nil {
		fmt.Println("[getDocByStringID]", err)
		return err
	}
	return nil
} // TODO return cursor to avoid multiple Finds

func recordVotes(rcvd rcvdAnswers, coll *mongo.Collection) error {
	var quiz schema.PopularityCollRead
	err := getDocByStringID(rcvd.QuizID, coll, &quiz)
	if err != nil {
		fmt.Println("[recordVotes]", err)
		return err
	}

	for _, name := range rcvd.Names {
		if _, exists := quiz.VoteCount[name]; exists {
			quiz.VoteCount[name]++
		}
	}

	// write to db
	objId, err := primitive.ObjectIDFromHex(rcvd.QuizID)
	if err != nil {
		fmt.Println("[recordVotes]", err)
		return err
	}
	filter := bson.D{{Key: "_id", Value: objId}}
	content := bson.D{{Key: "vote_count", Value: quiz.VoteCount}}
	updateResult, err := utils.MongoUpdate(coll, filter, content)
	if err != nil {
		fmt.Println("[recordVotes]", err)
		return err
	}
	fmt.Println("[recordVotes]", updateResult)

	return nil
}

func addParticipantCnt(rcvd rcvdAnswers, coll *mongo.Collection) error {
	var quiz schema.PopularityCollRead
	err := getDocByStringID(rcvd.QuizID, coll, &quiz)
	if err != nil {
		fmt.Println("[recordVotes]", err)
		return err
	}

	objId, err := primitive.ObjectIDFromHex(rcvd.QuizID)
	if err != nil {
		fmt.Println("[recordVotes]", err)
		return err
	}
	filter := bson.D{{Key: "_id", Value: objId}}
	content := bson.D{{Key: "participant_cnt", Value: quiz.ParticipantCnt + 1}}
	updateResult, err := utils.MongoUpdate(coll, filter, content)
	if err != nil {
		fmt.Println("[recordVotes]", err)
		return err
	}
	fmt.Println("[recordVotes]", updateResult)

	return nil
}

func countVoteRate(vc int, total int) float32 {
	return float32(vc) / float32(total)
}

type rate struct {
	Name     string  `json:"name"`
	VoteRate float32 `json:"vote_rate"`
}

type ByRate []rate

func (a ByRate) Len() int {
	return len(a)
}

func (a ByRate) Less(i, j int) bool {
	return a[i].VoteRate > a[j].VoteRate // from large to small
}

func (a ByRate) Swap(i, j int) {
	a[i].Name, a[j].Name = a[j].Name, a[i].Name
	a[i].VoteRate, a[j].VoteRate = a[j].VoteRate, a[i].VoteRate
}

func sortVotes(rcvd rcvdAnswers, coll *mongo.Collection) ([]rate, error) {
	var quiz schema.PopularityCollRead
	err := getDocByStringID(rcvd.QuizID, coll, &quiz)
	if err != nil {
		fmt.Println("[recordVotes]", err)
		return nil, err
	}

	var rankedRate []rate

	for name, vc := range quiz.VoteCount {
		r := rate{
			Name:     name,
			VoteRate: countVoteRate(vc, quiz.ParticipantCnt),
		}
		rankedRate = append(rankedRate, r)
	}

	sort.Sort(ByRate(rankedRate))
	return rankedRate, nil
}

func HandleAnswers(mc *pkg.MongoDBClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var rcvd rcvdAnswers
		err := utils.ReadPostBody(c, &rcvd)
		if err != nil {
			fmt.Println("[handleAnswers]", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vote."})
		}

		coll := mc.GetCollection("yellowbear", "quizzes")

		err = addParticipantCnt(rcvd, coll)
		if err != nil {
			fmt.Println("[handleAnswers]", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vote."})
		}

		err = recordVotes(rcvd, coll)
		if err != nil {
			fmt.Println("[handleAnswers]", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vote."})
		}

		rankedVoteRate, err := sortVotes(rcvd, coll)
		if err != nil {
			fmt.Println("[handleAnswers]", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vote."})
		}
		fmt.Println("[HandleAnswers]", rankedVoteRate)

		err = utils.RespOkWithBody(c, rankedVoteRate)
		if err != nil {
			fmt.Println("[HandleAnswers]", err)
			return
		}
	}
}
