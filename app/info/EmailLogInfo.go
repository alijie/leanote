package info

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 发送邮件
type EmailLog struct {
	LogId primitive.ObjectID `bson:"_id"`

	Email   string `Email`   // 发送者
	Subject string `Subject` // 主题
	Body    string `Body`    // 内容
	Msg     string `Msg`     // 发送失败信息
	Ok      bool   `Ok`      // 发送是否成功

	CreatedTime time.Time `CreatedTime`
}
