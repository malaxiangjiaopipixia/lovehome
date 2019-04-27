package controllers

import (
	"encoding/json"
	"lovehome/model"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *UserController) Register() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	//获取前端传来的数据
	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)

	beego.Info(`resp["mobile"]=`, resp["mobile"])
	beego.Info(`resp["password"]=`, resp["password"])
	beego.Info(`resp["sms_code"]=`, resp["sms_code"])

	//插入数据库

	o := orm.NewOrm()
	user := model.User{}
	user.Name = resp["mobile"].(string)
	user.Passwd = resp["password"].(string)

	id, err := o.Insert(user)
	if err != nil {
		resp["error"] = 4002
		resp["errmsg"] = "注册失败......"
		return
	}
	beego.Info("Regist Success!!!id=", id)
	resp["error"] = 0
	resp["errmsg"] = "注册成功......"
	return
}
