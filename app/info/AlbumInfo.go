package info

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Album struct {
	AlbumId     primitive.ObjectID `bson:"_id,omitempty"` //
	UserId      primitive.ObjectID `bson:"UserId"`
	Name        string             `Name` // album name
	Type        int                `Type` // type, the default is image: 0
	Seq         int                `Seq`
	CreatedTime time.Time          `CreatedTime`
}
