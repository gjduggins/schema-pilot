package model

// Column represents a table column.
type Column struct {
	Name string
	Type string
}

// Table represents a database table.
type Table struct {
	Name    string
	Columns []Column
}
