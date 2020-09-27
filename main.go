package main

import (
	_ "bee_api_template/commonds"
	"bee_api_template/controllers"
	"bee_api_template/packages"
	_ "bee_api_template/routers"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	initConfig()
}

func main() {
	err := packages.GetDbClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = packages.GetRedisClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	//toolbox.StartTask()
	defer func() {
		_ = packages.Db.Close()
		_ = packages.Redis.Close()
		//toolbox.StopTask()
	}()
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

/*
 * 初始化配置
 */
func initConfig() {
	runMode := beego.AppConfig.DefaultString("runmode", "prod")
	err := beego.LoadAppConfig("ini", fmt.Sprintf("conf/%s.app.conf", runMode))
	if err != nil {
		fmt.Println(err)
		return
	}

	beego.BConfig.ServerName = "nginx"
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.EnableErrorsShow = false
	beego.BConfig.EnableErrorsRender = false
	beego.BConfig.WebConfig.AutoRender = false
	return
}
