package info

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 随机token
// 验证邮箱
// 找回密码

// token type
const (
	TokenPwd = iota
	TokenActiveEmail
	TokenUpdateEmail
)

// 过期时间
const (
	PwdOverHours         = 2.0
	ActiveEmailOverHours = 48.0
	UpdateEmailOverHours = 2.0
)

type Token struct {
	UserId      primitive.ObjectID `bson:"_id"`
	Email       string             `Email`
	Token       string             `Token`
	Type        int                `Type`
	CreatedTime time.Time          `CreatedTime`
}
