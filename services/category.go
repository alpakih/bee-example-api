package services

import (
	"bee-example-api/dto"
	"bee-example-api/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func StoreCategory(c dto.CategoryDto) (category models.Category, err error) {

	o := orm.NewOrm()
	category.Code = c.Code
	category.Name = c.Name

	_, e := o.Insert(&category)
	if e != nil {
		err = e
	}
	return category, err
}
func GetCategoryById(id int) (category models.Category, err error) {
	o := orm.NewOrm()
	category.Id = id
	e := o.Read(&category)
	if e != nil {
		err = e
	}

	return category, err
}

func GetAllCategory() (resultData interface{}, err error) {
	o := orm.NewOrm()
	var c [] models.Category

	_, e := o.QueryTable("categories").All(&c)
	if e != nil {
		err = e
	}

	var result []beego.M
	data := make(beego.M)
	for _, v := range c {
		data = make(beego.M)
		data["id"] = v.Id
		data["code"] = v.Code
		data["name"] = v.Name
		result = append(result, data)
	}

	resultData = result
	return resultData, err
}

func UpdateCategory(id int, cDto dto.CategoryDto) (category models.Category, error error) {
	o := orm.NewOrm()
	category.Id = id
	if o.Read(&category) == nil {
		category.Name = cDto.Name
		category.Code = cDto.Code
		if num, err := o.Update(&category); err == nil {
			fmt.Println(num)
		}
	}
	return category, error
}

func DeleteCategory(id int) (res bool, err error) {
	o := orm.NewOrm()
	if num, erD := o.Delete(&models.Category{Id: id}); erD == nil {
		res = true
		fmt.Println(num)
	} else {
		err = erD
		res = false
	}
	return res, err
}
