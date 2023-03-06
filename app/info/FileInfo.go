package info

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	FileId         primitive.ObjectID `bson:"_id,omitempty"` //
	UserId         primitive.ObjectID `bson:"UserId"`
	AlbumId        primitive.ObjectID `bson:"AlbumId"`
	Name           string             `Name`  // file name
	Title          string             `Title` // file name or user defined for search
	Size           int64              `Size`  // file size (byte)
	Type           string             `Type`  // file type, "" = image, "doc" = word
	Path           string             `Path`  // the file path
	IsDefaultAlbum bool               `IsDefaultAlbum`
	CreatedTime    time.Time          `CreatedTime`

	FromFileId primitive.ObjectID `bson:"FromFileId,omitempty"` // copy from fileId, for collaboration
}
