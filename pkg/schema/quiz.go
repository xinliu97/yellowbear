package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

// for mongodb driver, maping mongo collection to Go struct

// Free input to hit candidates, no correct answer, only show statistics
type PopularityColl struct {
	Question string `bson:"question"`
	Attributes []string `bson:"attributes"`
	PrimaryAttr string `bson:"primary_attr"`
	Candidates []map[string]string `bson:"candidates"`
	ParticipantCnt int `bson:"participantcnt"`
	VoteCount map[string]int `bson:"vote_count"`  // PrimaryAttr: vote count number
}

type PopularityCollRead struct {
	ID primitive.ObjectID `bson:"_id"`
	Question string `bson:"question"`
	Attributes []string `bson:"attributes"`
	PrimaryAttr string `bson:"primary_attr"`
	Candidates []map[string]string `bson:"candidates"`
	ParticipantCnt int `bson:"participantcnt"`
	VoteCount map[string]int `bson:"vote_count"`  // PrimaryAttr: vote count number
}

type PopularityCreation struct {
	Question string `json:"question"`
    Attributes []string `json:"attributes"`
    PrimaryAttr string `json:"primary_attr"`
    Candidates []map[string]string `json:"candidates"`
}

// other types of quizzes
type XXXColl struct {}
