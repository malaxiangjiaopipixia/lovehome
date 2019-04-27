package controllers

import (
	"github.com/astaxie/beego"
)

type HouseIndexController struct {
	beego.Controller
}

func (this *HouseIndexController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *HouseIndexController) GethouseIndex() {
	resp := make(map[string]interface{})
	resp["error"] = 0
	resp["errmsg"] = "查询成功"
	this.RetData(resp)
}
