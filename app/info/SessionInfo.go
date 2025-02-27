package info

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// http://docs.mongodb.org/manual/tutorial/expire-data/
type Session struct {
	Id primitive.ObjectID `bson:"_id,omitempty"` // 没有意义

	SessionId string `bson:"SessionId"` // SessionId

	LoginTimes int    `LoginTimes` // 登录错误时间
	Captcha    string `Captcha`    // 验证码

	UserId string `UserId` // API时有值UserId

	CreatedTime time.Time `CreatedTime`
	UpdatedTime time.Time `UpdatedTime` // 更新时间, expire这个时间会自动清空
}
