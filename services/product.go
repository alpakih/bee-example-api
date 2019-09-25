package services

import (
	"bee-example-api/dto"
	"bee-example-api/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func StoreProduct(c dto.ProductDto) (product models.Product, err error) {

	o := orm.NewOrm()

	var category models.Category
	category.Id = c.CategoryId

	product.Name = c.Name
	product.Description = c.Description
	product.Price = c.Price
	eC := o.Read(&category)

	if eC != nil {
		beego.Error(eC)
	} else {
		product.Category = &category
	}
	beego.Info("Category", category)
	_, e := o.Insert(&product)
	if e != nil {
		err = e
	}
	beego.Info(err)
	return product, err
}

func GetProductById(id int) (product models.Product, err error) {
	o := orm.NewOrm()
	product.Id = id
	e := o.Read(&product)
	if e != nil {
		err = e
	}

	return product, err
}

func GetAllProduct() (i interface{}, err error) {
	o := orm.NewOrm()
	var c [] models.Product

	_, e := o.QueryTable("products").All(&c)
	if e != nil {
		err = e
	}

	var result []map[string]interface{}
	data := make(map[string]interface{})
	for _, v := range c {
		data = make(map[string]interface{})
		data["id"] = v.Id
		data["name"] = v.Name
		data["description"] = v.Description
		data["price"] = v.Price
		data["category_id"] = v.Category.Id
		data["category_name"] = v.Category.Name

		result = append(result, data)
	}

	i = result
	return i, err
}

func UpdateProduct(id int, cDto dto.ProductDto) (product models.Product, error error) {
	o := orm.NewOrm()
	product.Id = id

	var category models.Category
	category.Id = cDto.CategoryId
	eC := o.Read(&category)

	if eC != nil {
		beego.Error("Cannot find category")
	}
	if o.Read(&product) == nil {
		product.Name = cDto.Name
		product.Description = cDto.Description
		product.Price = cDto.Price
		product.Category = &category
		if num, err := o.Update(&product); err == nil {
			fmt.Println(num)
		}
	}
	return product, error
}

func DeleteProduct(id int) (res bool, err error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&models.Product{Id: id}); err == nil {
		res = true
		fmt.Println(num)
	}
	return res, err
}
