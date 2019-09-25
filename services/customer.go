package services

import (
	"bee-example-api/dto"
	"bee-example-api/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

func StoreCustomer(c dto.CreateCustomerDto) (customer models.Customer, err error) {

	o := orm.NewOrm()
	customer.Name = c.Name
	customer.Email = c.Email
	customer.Address = c.Address
	customer.Phone = c.Phone

	_, e := o.Insert(&customer)
	if e != nil {
		err = e
	}
	return customer, err
}
func GetCustomerById(id int) (customer models.Customer, err error) {
	o := orm.NewOrm()
	customer.Id = id
	e := o.Read(&customer)
	if e != nil {
		err = e
	}

	return customer, err
}

func GetAllCustomer() (i interface{}, err error) {
	o := orm.NewOrm()
	var c [] models.Customer

	_, e := o.QueryTable("customers").All(&c)
	if e != nil {
		err = e
	}

	var result []map[string]interface{}
	data := make(map[string]interface{})
	for _, v := range c {
		data = make(map[string]interface{})
		data["id"] = v.Id
		data["name"] = v.Name
		data["email"] = v.Email
		data["address"] = v.Address
		data["phone"] = v.Phone

		result = append(result, data)
	}

	i = result
	return i, err
}

func UpdateCustomer(id int, cDto dto.UpdateCustomerDto) (customer models.Customer, error error) {
	o := orm.NewOrm()
	customer.Id = id
	if o.Read(&customer) == nil {
		customer.Name = cDto.Name
		customer.Email = cDto.Email
		customer.Phone = cDto.Phone
		customer.Address = cDto.Address
		if num, err := o.Update(&customer); err == nil {
			fmt.Println(num)
		}
	}
	return customer, error
}

func DeleteCustomer(id int) (res bool, err error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&models.Customer{Id: id}); err == nil {
		res = true
		fmt.Println(num)
	}
	return res, err
}
