package utils

import (
	"os"
	"fmt"
	"io"
	"encoding/json"
	"yellowbear/pkg/schema"
)

func ReadPopularityCreationJson(filePath string, pc *schema.PopularityCreation) error {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("[CreatePopularity] Failed to open json file.", err)
		return err
	}
	defer file.Close()

	rawJson, err := io.ReadAll(file)
    if err != nil {
        fmt.Println("[CreatePopularity] Failed to read json file.", err)
		return err
    }

	err = json.Unmarshal(rawJson, pc)
	if err != nil {
        fmt.Println("[CreatePopularity] Failed to read json file.", err)
		return err
    }

	return nil
}

// func getFieldByName(s interface{}, fieldName string) interface{} {
//     v := reflect.ValueOf(s)

//     if v.Kind() == reflect.Struct {
//         field := v.FieldByName(fieldName)
//         if field.IsValid() && field.CanInterface() {
//             return field.Interface()
//         }
//     }

//     return nil
// }

func getCandidatePrimaryAttrValue(candidate map[string]string, popColl *schema.PopularityColl) string {
	return candidate[popColl.PrimaryAttr]
}

func initVoteCnt(popColl *schema.PopularityColl) error {
	popColl.VoteCount = make(map[string]int)
	for _, candidate := range popColl.Candidates {
		popColl.VoteCount[getCandidatePrimaryAttrValue(candidate, popColl)] = 0
	}

	return nil
}

func ConstructPopularityCollection(pc schema.PopularityCreation, popColl *schema.PopularityColl) {
	popColl.Question = pc.Question
	popColl.Attributes = pc.Attributes
	popColl.PrimaryAttr = pc.PrimaryAttr
	popColl.Candidates = pc.Candidates
	popColl.ParticipantCnt = 0
	initVoteCnt(popColl)
}