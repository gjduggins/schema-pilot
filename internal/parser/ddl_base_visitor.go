// Code generated from ddl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ddl

import (
	"github.com/antlr4-go/antlr/v4"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type BaseddlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseddlVisitor) VisitParse(ctx *ParseContext) interface{} {
	log.Log.Info("Visiting Parse 1")
	return v.VisitChildren(ctx)
}

func (v *BaseddlVisitor) VisitCreateTable(ctx *CreateTableContext) interface{} {
	log.Log.Info("VisitCreateTable 1")
	return v.VisitChildren(ctx)
}

func (v *BaseddlVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	log.Log.Info("VisitTableName 1")
	return v.VisitChildren(ctx)
}

func (v *BaseddlVisitor) VisitColumnDefList(ctx *ColumnDefListContext) interface{} {
	log.Log.Info("VisitColumnDefList 1")
	return v.VisitChildren(ctx)
}

func (v *BaseddlVisitor) VisitColumnDef(ctx *ColumnDefContext) interface{} {
	log.Log.Info("VisitColumnDef 1")
	return v.VisitChildren(ctx)
}

func (v *BaseddlVisitor) VisitDataType(ctx *DataTypeContext) interface{} {
	log.Log.Info("VisitDataType 1")
	return v.VisitChildren(ctx)
}
