package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Customer), new(Category), new(Product))
}
