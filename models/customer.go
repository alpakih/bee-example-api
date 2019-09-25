package models

import (
	"bee-example-api/mapper"
	"time"
)

type Customer struct {
	Id        int       `orm:"auto"`
	Name      string    `orm:"column(name);size(100)"`
	Email     string    `orm:"column(email);size(50);unique"`
	Address   string    `orm:"column(address);size(255);null"`
	Phone     string    `orm:"column(phone);size(13);unique"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (c Customer) Map() interface{} {
	customerMapper := mapper.CustomerMapper{
		Id:      c.Id,
		Name:    c.Name,
		Email:   c.Email,
		Address: c.Address,
		Phone:   c.Phone,
	}
	return customerMapper
}

func (c Customer) TableName() string {
	return "customers"
}
