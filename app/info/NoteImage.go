package info

import "go.mongodb.org/mongo-driver/bson/primitive"

// 笔记内部图片
type NoteImage struct {
	NoteImageId primitive.ObjectID `bson:"_id,omitempty"` // 必须要设置bson:"_id" 不然mgo不会认为是主键
	NoteId      primitive.ObjectID `bson:"NoteId"`        // 笔记
	ImageId     primitive.ObjectID `bson:"ImageId"`       // 图片fileId
}
