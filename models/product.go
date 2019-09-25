package models

import (
	"bee-example-api/mapper"
	"time"
)

type Product struct {
	Id          int       `orm:"auto"`
	Name        string    `orm:"column(name);size(100)"`
	Description string    `orm:"column(description);size(255);null"`
	Price       int64     `orm:"column(price)"`
	Category    *Category `orm:"rel(fk);column(category_id)"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
}

func (p Product) Map() interface{} {
	productMapper := mapper.ProductMapper{
		Id:           p.Id,
		Name:         p.Name,
		Description:  p.Description,
		Price:        p.Price,
		CategoryId:   p.Category.Id,
		CategoryName: p.Category.Name,
	}

	return productMapper
}

func (p Product) TableName() string {
	return "products"
}
