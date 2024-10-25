package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Level_20241022_164329 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Level_20241022_164329{}
	m.Created = "20241022_164329"

	migration.Register("Level_20241022_164329", m)
}

// Run the migrations
func (m *Level_20241022_164329) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE Level(
		id SERIAL NOT NULL,
		level NUMERIC(1),
		zero_num NUMERIC(2),
		delete_flg boolean default false,
		PRIMARY KEY (id)
		)`)
}

// Reverse the migrations
func (m *Level_20241022_164329) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
