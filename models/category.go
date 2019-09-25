package models

import (
	"bee-example-api/mapper"
	"time"
)

type Category struct {
	Id        int       `orm:"auto"`
	Code      string    `orm:"column(code);size(20);unique"`
	Name      string    `orm:"column(name);size(50)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (c Category) Map() interface{} {
	categoryMapper := mapper.CategoryMapper{
		Id:   c.Id,
		Code: c.Code,
		Name: c.Name,
	}
	return categoryMapper
}

func (c Category) TableName() string {
	return "categories"
}
