package schema

type Institution struct {
	Name string `bson:"name"`
	Location string `bson:"location"`
	VoteCount int32 `bson:"voteCount"`
}