package info

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 这里主要是为了统计每个tag的note数目
// 暂时没用
/*
type TagNote struct {
	TagId   primitive.ObjectID `bson:"_id,omitempty"` // 必须要设置bson:"_id" 不然mgo不会认为是主键
	UserId  primitive.ObjectID `bson:"UserId"`
	Tag   string        `Title`   // 标题
	NoteNum int           `NoteNum` // note数目
}
*/

// 每个用户一条记录, 存储用户的所有tags
type Tag struct {
	UserId primitive.ObjectID `bson:"_id"`
	Tags   []string           `Tags`
}

// v2 版标签
type NoteTag struct {
	TagId       primitive.ObjectID `bson:"_id"`
	UserId      primitive.ObjectID `UserId` // 谁的
	Tag         string             `Tag`    // UserId, Tag是唯一索引
	Usn         int                `Usn`    // Update Sequence Number
	Count       int                `Count`  // 笔记数
	CreatedTime time.Time          `CreatedTime`
	UpdatedTime time.Time          `UpdatedTime`
	IsDeleted   bool               `IsDeleted` // 删除位
}

type TagCount struct {
	TagCountId primitive.ObjectID `bson:"_id,omitempty"`
	UserId     primitive.ObjectID `UserId` // 谁的
	Tag        string             `Tag`
	IsBlog     bool               `IsBlog` // 是否是博客的tag统计
	Count      int                `Count`  // 统计数量
}

/*
type TagsCounts []TagCount
func (this TagsCounts) Len() int {
	return len(this)
}
func (this TagsCounts) Less(i, j int) bool {
	return this[i].Count > this[j].Count
}
func (this TagsCounts) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
*/
