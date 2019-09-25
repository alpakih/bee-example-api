package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Product_20190923_140156 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Product_20190923_140156{}
	m.Created = "20190923_140156"

	_ = migration.Register("Product_20190923_140156", m)
}

// Run the migrations
func (m *Product_20190923_140156) Up() {
	 m.SQL("CREATE TABLE products (" +
	 	"id SERIAL NOT NULL, " +
	 	"name VARCHAR(100) NOT NULL, " +
	 	"description VARCHAR(255), " +
	 	"category_id INT, " +
	 	"price BIGINT NOT NULL," +
	 	"created_at TIMESTAMP," +
	 	"updated_at TIMESTAMP, " +
	 	"PRIMARY KEY (id), " +
	 	"FOREIGN KEY(category_id) REFERENCES categories(id))")
}

// Reverse the migrations
func (m *Product_20190923_140156) Down() {
	m.SQL("DROP TABLE products")
}
