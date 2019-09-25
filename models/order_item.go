package models

import "time"

type OrderItem struct {
	Id        int     `orm:"auto"`
	Order     *Order    `orm:"rel(fk)"`
	Price     int64     `orm:"column(price)"`
	Qty       int       `orm:"column(qty)"`
	SubTotal  int64     `orm:"column(sub_total)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (c OrderItem) TableName() string {
	return "order_items"
}
