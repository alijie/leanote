package controllers

import (
	"github.com/revel/revel"
	//	"encoding/json"
	//	"go.mongodb.org/mongo-driver/bson"
	. "leanote/app/lea"
	"leanote/app/lea/captcha"

	//	"leanote/app/types"
	//	"io/ioutil"
	//	"fmt"
	//	"math"
	//	"os"
	//	"path"
	//	"strconv"
	"io"
	"net/http"
)

// 验证码服务
type Captcha struct {
	BaseController
}

type Ca string

func (r Ca) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusOK, "image/png")
}

func (c Captcha) Get() revel.Result {
	c.Response.ContentType = "image/png"
	image, str := captcha.Fetch()
	out := io.Writer(c.Response.GetWriter())
	image.WriteTo(out)

	sessionId := c.GetSession("_ID")
	//	LogJ(c.Session)
	//	Log("------")
	//	Log(str)
	//	Log(sessionId)
	Log("..")
	sessionService.SetCaptcha(sessionId, str)

	return c.Render()
}
