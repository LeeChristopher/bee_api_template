package controllers

import (
	"bee_api_template/packages"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type BasePreparer interface {
	Initialization()
}

type baseController struct {
	beego.Controller
	startTime time.Time
}

type Response struct {
	Runtime float64     `json:"runtime"`
	Code    uint64      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (m *baseController) Prepare() {
	m.startTime = packages.GetNow()

	if app, ok := m.AppController.(BasePreparer); ok {
		app.Initialization()
	}
}

func (m *baseController) Finish() {
	fmt.Println("end")
}

func (m *baseController) SetResponse(code uint64, message string, data interface{}) {
	if data == nil {
		data = make([]uint64, 0, 0)
	}

	responseData := &Response{
		Runtime: time.Since(m.startTime).Seconds(),
		Code:    code,
		Message: message,
		Data:    data,
	}
	_ = m.Ctx.Output.JSON(responseData, false, false)
	m.StopRun()
}
