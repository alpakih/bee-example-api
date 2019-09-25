package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Customer_20190923_140022 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Customer_20190923_140022{}
	m.Created = "20190923_140022"

	_ = migration.Register("Customer_20190923_140022", m)
}

// Run the migrations
func (m *Customer_20190923_140022) Up() {
	m.SQL("CREATE TABLE customers (" +
		"id SERIAL NOT NULL PRIMARY KEY, " +
		"name VARCHAR(100) NOT NULL," +
		"email VARCHAR(50) NOT NULL UNIQUE, " +
		"phone VARCHAR(13) NOT NULL UNIQUE," +
		"address VARCHAR(255), " +
		"created_at TIMESTAMP, " +
		"updated_at TIMESTAMP) ")
}

// Reverse the migrations
func (m *Customer_20190923_140022) Down() {
	m.SQL("DROP TABLE customers")

}
