package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Order_20190923_140213 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Order_20190923_140213{}
	m.Created = "20190923_140213"

	migration.Register("Order_20190923_140213", m)
}

// Run the migrations
func (m *Order_20190923_140213) Up() {
	m.SQL("CREATE TABLE orders (" +
		"id SERIAL NOT NULL," +
		"transaction_code VARCHAR(100) NOT NULL UNIQUE," +
		"customer_id INT NOT NULL," +
		"product_id INT NOT NULL," +
		"transaction_date TIMESTAMP NOT NULL," +
		"total BIGINT NOT NULL," +
		"created_at TIMESTAMP," +
		"updated_at TIMESTAMP," +
		"PRIMARY KEY (id)," +
		"FOREIGN KEY(customer_id) REFERENCES customers(id)," +
		"FOREIGN KEY (product_id) REFERENCES products(id))")
}

// Reverse the migrations
func (m *Order_20190923_140213) Down() {
	m.SQL("DROP TABLE orders")
}
