package model

type Order struct {
	OrderId   int64
	OrderName string `orm:"size(10)"`
	OrderText string `orm:"size(100)"`
	User      *User  `orm:"rel(fk)"`
}

type User struct {
	Id     int64
	Name   string   `orm:"size(10)"`
	Passwd string   `orm:"size(10)"`
	Order  *[]Order `orm:"reverse(many)"`
}
