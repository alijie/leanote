package info

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 分组
type Group struct {
	GroupId     primitive.ObjectID `bson:"_id"` // 谁的
	UserId      primitive.ObjectID `UserId`     // 所有者Id
	Title       string             `Title`      // 标题
	UserCount   int                `UserCount`  // 用户数
	CreatedTime time.Time          `CreatedTime`

	Users []User `Users,omitempty` // 分组下的用户, 不保存, 仅查看
}

// 分组好友
type GroupUser struct {
	GroupUserId primitive.ObjectID `bson:"_id"` // 谁的
	GroupId     primitive.ObjectID `GroupId`    // 分组
	UserId      primitive.ObjectID `UserId`     //  用户
	CreatedTime time.Time          `CreatedTime`
}
