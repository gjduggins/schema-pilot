package manager

import (
	"fmt"
	"runtime/debug"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gjduggins/schemapilot-operator/internal/model"
	"github.com/gjduggins/schemapilot-operator/internal/parser"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// TableVisitor extracts table definitions from the parse tree.
type OracleVisitor struct {
	parser.BaseddlVisitor
	tbl model.Table
}

func NewOracleVisitor() *OracleVisitor {
	return &OracleVisitor{
		BaseddlVisitor: parser.BaseddlVisitor{
			BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
		},
	}
}

func (v *OracleVisitor) VisitCreateTable(ctx *parser.CreateTableContext) interface{} {
	tableName := ctx.TableName().IDENTIFIER().GetText()
	v.tbl.Name = tableName

	for _, colCtx := range ctx.ColumnDefList().(*parser.ColumnDefListContext).AllColumnDef() {
		col := colCtx.(*parser.ColumnDefContext).Accept(v).(model.Column)
		v.tbl.Columns = append(v.tbl.Columns, col)
	}

	return v.tbl
}

// VisitColumnDef extracts column name and type
func (v *OracleVisitor) VisitColumnDef(ctx *parser.ColumnDefContext) interface{} {

	name := ctx.IDENTIFIER().GetText()
	typ := ctx.DataType().GetText()
	return model.Column{Name: name, Type: typ}
}

func (v *OracleVisitor) VisitParse(ctx *parser.ParseContext) interface{} {
	tables := ctx.AllCreateTable()
	if len(tables) == 0 {
		return nil
	}

	// For now, just handle the first table
	return v.VisitCreateTable(tables[0].(*parser.CreateTableContext))
}

func ParseSQL(ddl string) (model.Table, error) {

	is := antlr.NewInputStream(ddl)
	lexer := parser.NewddlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewddlParser(stream)
	v := NewOracleVisitor()
	tree := p.Parse()
	ruleNames := p.GetRuleNames()
	s := antlr.TreesStringTree(tree, ruleNames, p)
	log.Log.Info("Parse Tree:\n" + s)
	fmt.Printf("ROOT TYPE = %T\n", tree)
	root := tree.(*parser.ParseContext)

	tables := root.AllCreateTable()
	fmt.Println("count =", len(tables))

	if tree == nil {
		return v.tbl, fmt.Errorf("parser returned nil tree")
	}

	defer func() {
		if r := recover(); r != nil {
			log.Log.Error(nil, "panic in visitor.Visit", "panic", r)
			debug.PrintStack()
		}
	}()

	log.Log.Info("Visiting parse tree...")
	tree.Accept(v)

	log.Log.Info("Table Name is " + v.tbl.Name)
	return v.tbl, nil
}
