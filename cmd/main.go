package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"yellowbear/pkg/api"
	"yellowbear/pkg/quizManage"
	"yellowbear/pkg/utils"
)

func main() {
	config, err := utils.ReadConfigYaml()
	if err != nil {
		fmt.Println("[main]", err)
		return
	}

	mc, err := utils.NewMongoDBClient(config.Database.Mongo.Uri)
	if err != nil {
		fmt.Println("[main] Failed to connect to mongodb client.", err)
	}
	defer func(mc *utils.MongoDBClient) {
		err := mc.Disconnect()
		if err != nil {
			fmt.Println("[main]", err)
		}
	}(mc)

	createSamplesOnly := flag.Bool("sample", false, "Construct 3 sample quizzes.")
	flag.Parse()
	if *createSamplesOnly {
		fmt.Println("--createSamplesOnly--")
		err = quizManage.Create3SamplePopularityQuizzes(mc, config)
		if err != nil {
			fmt.Println("[main]", err)
			return
		}
		return
	}

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
