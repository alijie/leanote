package info

import "go.mongodb.org/mongo-driver/bson/primitive"

// 建议
type Suggestion struct {
	Id         primitive.ObjectID `bson:"_id"`
	UserId     primitive.ObjectID `UserId`
	Addr       string             `Addr`
	Suggestion string             `Suggestion`
}
