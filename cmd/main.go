package main

import (
	// "go.mongodb.org/mongo-driver/mongo/options"
	"yellowbear/pkg"
	"yellowbear/pkg/quizManage"
	// "context"
	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/mongo"
	"fmt"
	"yellowbear/pkg/api"
)

func create3SamplePopularityQuizzes(mc *pkg.MongoDBClient) {
	quizManage.CreatePopularity(mc, "/root/yellowbear/pkg/schema/samplePopularityInput1.json")
	quizManage.CreatePopularity(mc, "/root/yellowbear/pkg/schema/samplePopularityInput2.json")
	quizManage.CreatePopularity(mc, "/root/yellowbear/pkg/schema/samplePopularityInputN.json")
}


func main() {
	/*
	 	create 3 popularity-type quizzes from json files into mongodb
	*/
	mc, err := pkg.NewMongoDBClient()
	if err != nil {
		fmt.Println("[Disconnect] Failed to connect to mongodb client.", err)
	}
	defer mc.Disconnect()

	// create3SamplePopularityQuizzes(mc)

	/*
		handle requests
	*/
	//create router
	router := gin.Default()
	// respond to http requests
	router.GET("quiz/all", api.ListAllQuizzes(mc))
	router.POST("quiz/submit", api.HandleAnswers(mc))
	// start to listen
	router.Run()
}



// // define a sample doc
	// sampleDoc := Institution {
	// 	Name: "东北大学",
	// 	Location: "辽宁省",
	// 	VoteCount: 0,
	// }
	// // insert it
	// insertResult, err := institutionsColl.InsertOne(context.TODO(), sampleDoc)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Inserted document with _id: %v\n", insertResult.InsertedID)