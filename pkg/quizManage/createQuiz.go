package quizmanage

import (
	"context"
	"yellowbear/pkg"
	"yellowbear/pkg/schema"
	"yellowbear/pkg/utils"
	"fmt"
)

const (
	dbName = "yellowbear"
	collName = "quizzes"
)


func CreatePopularity(mc *pkg.MongoDBClient, popularityFilePath string) error {
	var popCre schema.PopularityCreation
	err := utils.ReadPopularityCreationJson(popularityFilePath, &popCre)
	if err!=nil {
		fmt.Println("[CreatePopularity]", err)
		return nil
	}

	var popColl schema.PopularityColl
	utils.ConstructPopularityCollection(popCre, &popColl)
	coll := mc.GetCollection(dbName, collName)
	result, err := coll.InsertOne(context.TODO(), popColl)
	if err!=nil {
		fmt.Println("[CreatePopularity]", err)
		return err
	}
	fmt.Println("Successfully insert one popularity quiz mongo document, ID:", result.InsertedID)
	return nil
}