package tests

import (
	"leanote/app/service"
	"testing"

	"github.com/revel/revel"
)

func init() {
	revel.Init("dev", "leanote", "/Users/life/Documents/Go/package_base/src")
	service.InitService()
	service.ConfigS.InitGlobalConfigs()
}

// 测试登录
func TestSendMail(t *testing.T) {
	ok, err := service.EmailS.SendEmail("life@leanote.com", "你好", "你好吗")
	t.Log(ok)
	t.Log(err)
}
