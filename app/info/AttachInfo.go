package info

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Attach belongs to note
type Attach struct {
	AttachId     primitive.ObjectID `bson:"_id,omitempty"` //
	NoteId       primitive.ObjectID `bson:"NoteId"`        //
	UploadUserId primitive.ObjectID `bson:"UploadUserId"`  // 可以不是note owner, 协作者userId
	Name         string             `Name`                 // file name, md5, such as 13232312.doc
	Title        string             `Title`                // raw file name
	Size         int64              `Size`                 // file size (byte)
	Type         string             `Type`                 // file type, "doc" = word
	Path         string             `Path`                 // the file path such as: files/userId/attachs/adfadf.doc
	CreatedTime  time.Time          `CreatedTime`

	// FromFileId primitive.ObjectID `bson:"FromFileId,omitempty"` // copy from fileId, for collaboration
}
