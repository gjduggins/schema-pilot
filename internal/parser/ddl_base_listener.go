// Code generated from ddl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ddl

import "github.com/antlr4-go/antlr/v4"

// BaseddlListener is a complete listener for a parse tree produced by ddlParser.
type BaseddlListener struct{}

var _ ddlListener = &BaseddlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseddlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseddlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseddlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseddlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseddlListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseddlListener) ExitParse(ctx *ParseContext) {}

// EnterCreateTable is called when production createTable is entered.
func (s *BaseddlListener) EnterCreateTable(ctx *CreateTableContext) {}

// ExitCreateTable is called when production createTable is exited.
func (s *BaseddlListener) ExitCreateTable(ctx *CreateTableContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseddlListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseddlListener) ExitTableName(ctx *TableNameContext) {}

// EnterColumnDefList is called when production columnDefList is entered.
func (s *BaseddlListener) EnterColumnDefList(ctx *ColumnDefListContext) {}

// ExitColumnDefList is called when production columnDefList is exited.
func (s *BaseddlListener) ExitColumnDefList(ctx *ColumnDefListContext) {}

// EnterColumnDef is called when production columnDef is entered.
func (s *BaseddlListener) EnterColumnDef(ctx *ColumnDefContext) {}

// ExitColumnDef is called when production columnDef is exited.
func (s *BaseddlListener) ExitColumnDef(ctx *ColumnDefContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseddlListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseddlListener) ExitDataType(ctx *DataTypeContext) {}
