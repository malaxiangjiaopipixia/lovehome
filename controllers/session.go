package controllers

import (
	"lovehome/model"

	"github.com/astaxie/beego"
)

type SessionController struct {
	beego.Controller
}

func (this *SessionController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *SessionController) GetSessionData() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	user := model.User{}
	user.Name = "tyj"
	resp["error"] = 0
	resp["errmsg"] = "ok"
	resp["data"] = user

}

func (this *SessionController) DeleteSessionData() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	this.DelSession("name")
	resp["error"] = 0
	resp["errmsg"] = "ok"

}
