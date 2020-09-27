package packages

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

/*
 * 统一响应结构
 */
type CommonResponse struct {
	RunTime float64     `json:"runtime"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetNow() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

/*
 * 获取缓存中用户登陆Key
 */
func GetLoginKey(userId uint64) (key string) {
	return fmt.Sprintf("%s:%d:auth:token", beego.BConfig.AppName, userId)
}
