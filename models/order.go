package models

import "time"

type Order struct {
	Id              int     `orm:"auto"`
	TransactionCode string    `orm:"column(transaction_code);size(100)"`
	Customer        *Customer `orm:"rel(fk)"`
	Product         *Product  `orm:"rel(fk)"`
	Total           int64     `orm:"column(total)"`
	TransactionDate time.Time `orm:"type(datetime);column(transaction_date)"`
	CreatedAt       time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt       time.Time `orm:"auto_now_add;type(datetime)"`
}

func (c Order) TableName() string {
	return "orders"
}
