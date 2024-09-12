package main

import (
	"fmt"
	"yellowbear/pkg"
	"yellowbear/pkg/quizManage"
	// "context"
	"github.com/gin-gonic/gin"
	"yellowbear/pkg/api"
)

func create3SamplePopularityQuizzes(mc *pkg.MongoDBClient) {
	err := quizManage.CreatePopularity(mc, "/root/yellowbear/pkg/schema/samplePopularityInput1.json")
	if err != nil {
		return
	}
	err = quizManage.CreatePopularity(mc, "/root/yellowbear/pkg/schema/samplePopularityInput2.json")
	if err != nil {
		return
	}
	err = quizManage.CreatePopularity(mc, "/root/yellowbear/pkg/schema/samplePopularityInputN.json")
	if err != nil {
		return
	}
}

func main() {
	mc, err := pkg.NewMongoDBClient()
	if err != nil {
		fmt.Println("[Disconnect] Failed to connect to mongodb client.", err)
	}
	defer func(mc *pkg.MongoDBClient) {
		err := mc.Disconnect()
		if err != nil {
			fmt.Println("[main]", err)
		}
	}(mc)

	//create router
	router := gin.Default()
	// handle http requests
	router.GET("quiz/all", api.ListAllQuizzes(mc))
	router.GET("home/hot", api.QuizzesInHeatOrder(mc))
	router.POST("quiz/submit", api.HandleAnswers(mc))
	// start to listen
	err = router.Run()
	if err != nil {
		fmt.Println("[main]", err)
	}
}
