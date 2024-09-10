package main

import (
	// "go.mongodb.org/mongo-driver/mongo/options"
	"yellowbear/pkg"
	"yellowbear/pkg/quizManage"
	// "context"
	// "github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/mongo"
	"fmt"
	// "yellowbear/pkg/api"
)



func main() {
	/* 
		creat a mongodb client and connect to mongo instance 
	*/
	// identify mongodb server instance
	// create a client
	mc, err := pkg.NewMongoDBClient()
	if err != nil {
		fmt.Println("[Disconnect] Failed to connect to mongodb client.", err)
	}
	defer mc.Disconnect()

	quizmanage.CreatePopularity(mc, "/root/yellowbear/pkg/schema/samplePopularityInput1.json")
	quizmanage.CreatePopularity(mc, "/root/yellowbear/pkg/schema/samplePopularityInput2.json")
	quizmanage.CreatePopularity(mc, "/root/yellowbear/pkg/schema/samplePopularityInputN.json")
	
	
	// /*
	// 	handle requests
	// */
	// //create router
	// router := gin.Default()
	// // respond to http requests
	// router.POST("/quiz/admin", api.HdlInputs(client))
	// router.POST("/quiz/admin/", api.HdlInputs(client))
	// // router.POST("/quiz/admin/result", api.DisplayResult(client))
	// // start to listen
	// router.Run()
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