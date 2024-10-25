package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateSudokuGenerate_20241022_153030 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateSudokuGenerate_20241022_153030{}
	m.Created = "20241022_153030"

	migration.Register("CreateSudokuGenerate_20241022_153030", m)
}

// Run the migrations
func (m *CreateSudokuGenerate_20241022_153030) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE sudoku_generate(
		id SERIAL NOT NULL,
		sudoku_level NUMERIC(1),
		sudoku_zero_presence jsonb,
		sudoku_zero_absence jsonb,
		delete_flg boolean default false,
		PRIMARY KEY (id)
		)`)

}

// Reverse the migrations
func (m *CreateSudokuGenerate_20241022_153030) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
