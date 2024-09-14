package quizManage

import (
	"context"
	"fmt"
	"yellowbear/pkg/schema"
	"yellowbear/pkg/utils"
)

const (
	dbName   = "yellowbear"
	collName = "quizzes"
)

func CreatePopularity(mc *utils.MongoDBClient, popularityFilePath string) error {
	var popCre schema.PopularityCreation
	err := utils.ReadPopularityCreationJson(popularityFilePath, &popCre)
	if err != nil {
		fmt.Println("[CreatePopularity]", err)
		return nil
	}

	var popColl schema.PopularityColl
	utils.ConstructPopularityCollection(popCre, &popColl)
	coll := mc.GetCollection(dbName, collName)
	result, err := coll.InsertOne(context.TODO(), popColl)
	if err != nil {
		fmt.Println("[CreatePopularity]", err)
		return err
	}
	fmt.Println("Successfully insert one popularity quiz mongo document, ID:", result.InsertedID)
	return nil
}

func Create3SamplePopularityQuizzes(mc *utils.MongoDBClient, cfg *utils.Config) error {
	err := CreatePopularity(mc, cfg.Database.Mongo.Sample1quizFp)
	if err != nil {
		return err
	}
	err = CreatePopularity(mc, cfg.Database.Mongo.Sample2quizFp)
	if err != nil {
		return err
	}
	err = CreatePopularity(mc, cfg.Database.Mongo.SampleNquizFp)
	if err != nil {
		return err
	}

	return nil
}
