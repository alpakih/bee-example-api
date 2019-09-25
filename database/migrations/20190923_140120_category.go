package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Category_20190923_140120 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Category_20190923_140120{}
	m.Created = "20190923_140120"

	_ = migration.Register("Category_20190923_140120", m)
}

// Run the migrations
func (m *Category_20190923_140120) Up() {
	m.SQL("CREATE TABLE categories (" +
		"id SERIAL NOT NULL PRIMARY KEY, " +
		"code VARCHAR(20) NOT NULL, " +
		"name VARCHAR(50) NOT NULL, " +
		"created_at TIMESTAMP, " +
		"updated_at TIMESTAMP)")

}

// Reverse the migrations
func (m *Category_20190923_140120) Down() {
	m.SQL("DROP TABLE categories")

}
