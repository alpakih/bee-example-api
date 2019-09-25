package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type OrderItems_20190924_122324 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &OrderItems_20190924_122324{}
	m.Created = "20190924_122324"

	migration.Register("OrderItems_20190924_122324", m)
}

// Run the migrations
func (m *OrderItems_20190924_122324) Up() {
	m.SQL("CREATE TABLE order_items (" +
		"id SERIAL NOT NULL," +
		"order_id INT," +
		"price BIGINT NOT NULL," +
		"qty INT NOT NULL," +
		"sub_total BIGINT NOT NULL," +
		"created_at TIMESTAMP," +
		"updated_at TIMESTAMP," +
		"PRIMARY KEY (id)," +
		"FOREIGN KEY (order_id) REFERENCES orders(id))")
}

// Reverse the migrations
func (m *OrderItems_20190924_122324) Down() {
	m.SQL("DROP TABLE order_items")
}
