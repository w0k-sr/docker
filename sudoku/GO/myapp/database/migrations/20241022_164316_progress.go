package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Progress_20241022_164316 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Progress_20241022_164316{}
	m.Created = "20241022_164316"

	migration.Register("Progress_20241022_164316", m)
}

// Run the migrations
func (m *Progress_20241022_164316) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE Progress(
		progress_id SERIAL NOT NULL,
		sudoku_id NUMERIC,
		progress_rate NUMERIC(3),
		sudoku_solve jsonb,
		delete_flg boolean default false,
		PRIMARY KEY (progress_id)
		)`)

}

// Reverse the migrations
func (m *Progress_20241022_164316) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
