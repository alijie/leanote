package info

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 举报
type Report struct {
	ReportId primitive.ObjectID `bson:"_id"`
	NoteId   primitive.ObjectID `NoteId`

	UserId primitive.ObjectID `UserId` // UserId回复ToUserId
	Reason string             `Reason` // 评论内容

	CommentId primitive.ObjectID `CommendId,omitempty` // 对某条评论进行回复

	CreatedTime time.Time `CreatedTime`
}
