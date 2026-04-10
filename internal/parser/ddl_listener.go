// Code generated from ddl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ddl

import "github.com/antlr4-go/antlr/v4"

// ddlListener is a complete listener for a parse tree produced by ddlParser.
type ddlListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterCreateTable is called when entering the createTable production.
	EnterCreateTable(c *CreateTableContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterColumnDefList is called when entering the columnDefList production.
	EnterColumnDefList(c *ColumnDefListContext)

	// EnterColumnDef is called when entering the columnDef production.
	EnterColumnDef(c *ColumnDefContext)

	// EnterDataType is called when entering the dataType production.
	EnterDataType(c *DataTypeContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitCreateTable is called when exiting the createTable production.
	ExitCreateTable(c *CreateTableContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitColumnDefList is called when exiting the columnDefList production.
	ExitColumnDefList(c *ColumnDefListContext)

	// ExitColumnDef is called when exiting the columnDef production.
	ExitColumnDef(c *ColumnDefContext)

	// ExitDataType is called when exiting the dataType production.
	ExitDataType(c *DataTypeContext)
}
