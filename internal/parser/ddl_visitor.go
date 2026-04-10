// Code generated from ddl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ddl

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ddlParser.
type ddlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ddlParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by ddlParser#createTable.
	VisitCreateTable(ctx *CreateTableContext) interface{}

	// Visit a parse tree produced by ddlParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}

	// Visit a parse tree produced by ddlParser#columnDefList.
	VisitColumnDefList(ctx *ColumnDefListContext) interface{}

	// Visit a parse tree produced by ddlParser#columnDef.
	VisitColumnDef(ctx *ColumnDefContext) interface{}

	// Visit a parse tree produced by ddlParser#dataType.
	VisitDataType(ctx *DataTypeContext) interface{}
}
