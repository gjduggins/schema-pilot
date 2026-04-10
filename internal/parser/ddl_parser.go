// Code generated from ddl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ddl

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ddlParser struct {
	*antlr.BaseParser
}

var DdlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ddlParserInit() {
	staticData := &DdlParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "", "", "", "", "", "';'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "CREATE", "TABLE", "NUMBER", "VARCHAR2", "DATE", "SEMICOLON",
		"COMMA", "IDENTIFIER", "INT", "WS",
	}
	staticData.RuleNames = []string{
		"parse", "createTable", "tableName", "columnDefList", "columnDef", "dataType",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 12, 58, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 5, 0, 14, 8, 0, 10, 0, 12, 0, 17, 9, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 5, 3, 34, 8, 3, 10, 3, 12, 3, 37, 9, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 3, 5, 47, 8, 5, 1, 5, 3, 5, 50, 8, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 56, 8, 5, 1, 5, 0, 0, 6, 0, 2, 4, 6, 8, 10, 0, 0, 56,
		0, 15, 1, 0, 0, 0, 2, 20, 1, 0, 0, 0, 4, 28, 1, 0, 0, 0, 6, 30, 1, 0, 0,
		0, 8, 38, 1, 0, 0, 0, 10, 55, 1, 0, 0, 0, 12, 14, 3, 2, 1, 0, 13, 12, 1,
		0, 0, 0, 14, 17, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16,
		18, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 18, 19, 5, 0, 0, 1, 19, 1, 1, 0, 0,
		0, 20, 21, 5, 3, 0, 0, 21, 22, 5, 4, 0, 0, 22, 23, 3, 4, 2, 0, 23, 24,
		5, 1, 0, 0, 24, 25, 3, 6, 3, 0, 25, 26, 5, 2, 0, 0, 26, 27, 5, 8, 0, 0,
		27, 3, 1, 0, 0, 0, 28, 29, 5, 10, 0, 0, 29, 5, 1, 0, 0, 0, 30, 35, 3, 8,
		4, 0, 31, 32, 5, 9, 0, 0, 32, 34, 3, 8, 4, 0, 33, 31, 1, 0, 0, 0, 34, 37,
		1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 7, 1, 0, 0, 0,
		37, 35, 1, 0, 0, 0, 38, 39, 5, 10, 0, 0, 39, 40, 3, 10, 5, 0, 40, 9, 1,
		0, 0, 0, 41, 49, 5, 5, 0, 0, 42, 43, 5, 1, 0, 0, 43, 46, 5, 11, 0, 0, 44,
		45, 5, 9, 0, 0, 45, 47, 5, 11, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0,
		0, 0, 47, 48, 1, 0, 0, 0, 48, 50, 5, 2, 0, 0, 49, 42, 1, 0, 0, 0, 49, 50,
		1, 0, 0, 0, 50, 56, 1, 0, 0, 0, 51, 52, 5, 6, 0, 0, 52, 53, 5, 1, 0, 0,
		53, 54, 5, 11, 0, 0, 54, 56, 5, 2, 0, 0, 55, 41, 1, 0, 0, 0, 55, 51, 1,
		0, 0, 0, 56, 11, 1, 0, 0, 0, 5, 15, 35, 46, 49, 55,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ddlParserInit initializes any static state used to implement ddlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewddlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DdlParserInit() {
	staticData := &DdlParserStaticData
	staticData.once.Do(ddlParserInit)
}

// NewddlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewddlParser(input antlr.TokenStream) *ddlParser {
	DdlParserInit()
	this := new(ddlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &DdlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ddl.g4"

	return this
}

// ddlParser tokens.
const (
	ddlParserEOF        = antlr.TokenEOF
	ddlParserT__0       = 1
	ddlParserT__1       = 2
	ddlParserCREATE     = 3
	ddlParserTABLE      = 4
	ddlParserNUMBER     = 5
	ddlParserVARCHAR2   = 6
	ddlParserDATE       = 7
	ddlParserSEMICOLON  = 8
	ddlParserCOMMA      = 9
	ddlParserIDENTIFIER = 10
	ddlParserINT        = 11
	ddlParserWS         = 12
)

// ddlParser rules.
const (
	ddlParserRULE_parse         = 0
	ddlParserRULE_createTable   = 1
	ddlParserRULE_tableName     = 2
	ddlParserRULE_columnDefList = 3
	ddlParserRULE_columnDef     = 4
	ddlParserRULE_dataType      = 5
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllCreateTable() []ICreateTableContext
	CreateTable(i int) ICreateTableContext

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ddlParserRULE_parse
	return p
}

func InitEmptyParseContext(p *ParseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ddlParserRULE_parse
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ddlParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(ddlParserEOF, 0)
}

func (s *ParseContext) AllCreateTable() []ICreateTableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICreateTableContext); ok {
			len++
		}
	}

	tst := make([]ICreateTableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICreateTableContext); ok {
			tst[i] = t.(ICreateTableContext)
			i++
		}
	}

	return tst
}

func (s *ParseContext) CreateTable(i int) ICreateTableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateTableContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateTableContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ddlListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ddlListener); ok {
		listenerT.ExitParse(s)
	}
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ddlVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ddlParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ddlParserRULE_parse)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ddlParserCREATE {
		{
			p.SetState(12)
			p.CreateTable()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(18)
		p.Match(ddlParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICreateTableContext is an interface to support dynamic dispatch.
type ICreateTableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CREATE() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	TableName() ITableNameContext
	ColumnDefList() IColumnDefListContext
	SEMICOLON() antlr.TerminalNode

	// IsCreateTableContext differentiates from other interfaces.
	IsCreateTableContext()
}

type CreateTableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateTableContext() *CreateTableContext {
	var p = new(CreateTableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ddlParserRULE_createTable
	return p
}

func InitEmptyCreateTableContext(p *CreateTableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ddlParserRULE_createTable
}

func (*CreateTableContext) IsCreateTableContext() {}

func NewCreateTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateTableContext {
	var p = new(CreateTableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ddlParserRULE_createTable

	return p
}

func (s *CreateTableContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateTableContext) CREATE() antlr.TerminalNode {
	return s.GetToken(ddlParserCREATE, 0)
}

func (s *CreateTableContext) TABLE() antlr.TerminalNode {
	return s.GetToken(ddlParserTABLE, 0)
}

func (s *CreateTableContext) TableName() ITableNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *CreateTableContext) ColumnDefList() IColumnDefListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnDefListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnDefListContext)
}

func (s *CreateTableContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ddlParserSEMICOLON, 0)
}

func (s *CreateTableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateTableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateTableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ddlListener); ok {
		listenerT.EnterCreateTable(s)
	}
}

func (s *CreateTableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ddlListener); ok {
		listenerT.ExitCreateTable(s)
	}
}

func (s *CreateTableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ddlVisitor:
		return t.VisitCreateTable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ddlParser) CreateTable() (localctx ICreateTableContext) {
	localctx = NewCreateTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ddlParserRULE_createTable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Match(ddlParserCREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(21)
		p.Match(ddlParserTABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(22)
		p.TableName()
	}
	{
		p.SetState(23)
		p.Match(ddlParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(24)
		p.ColumnDefList()
	}
	{
		p.SetState(25)
		p.Match(ddlParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(26)
		p.Match(ddlParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsTableNameContext differentiates from other interfaces.
	IsTableNameContext()
}

type TableNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameContext() *TableNameContext {
	var p = new(TableNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ddlParserRULE_tableName
	return p
}

func InitEmptyTableNameContext(p *TableNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ddlParserRULE_tableName
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ddlParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ddlParserIDENTIFIER, 0)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ddlListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ddlListener); ok {
		listenerT.ExitTableName(s)
	}
}

func (s *TableNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ddlVisitor:
		return t.VisitTableName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ddlParser) TableName() (localctx ITableNameContext) {
	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ddlParserRULE_tableName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Match(ddlParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumnDefListContext is an interface to support dynamic dispatch.
type IColumnDefListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColumnDef() []IColumnDefContext
	ColumnDef(i int) IColumnDefContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsColumnDefListContext differentiates from other interfaces.
	IsColumnDefListContext()
}

type ColumnDefListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnDefListContext() *ColumnDefListContext {
	var p = new(ColumnDefListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ddlParserRULE_columnDefList
	return p
}

func InitEmptyColumnDefListContext(p *ColumnDefListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ddlParserRULE_columnDefList
}

func (*ColumnDefListContext) IsColumnDefListContext() {}

func NewColumnDefListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnDefListContext {
	var p = new(ColumnDefListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ddlParserRULE_columnDefList

	return p
}

func (s *ColumnDefListContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnDefListContext) AllColumnDef() []IColumnDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumnDefContext); ok {
			len++
		}
	}

	tst := make([]IColumnDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumnDefContext); ok {
			tst[i] = t.(IColumnDefContext)
			i++
		}
	}

	return tst
}

func (s *ColumnDefListContext) ColumnDef(i int) IColumnDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnDefContext)
}

func (s *ColumnDefListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ddlParserCOMMA)
}

func (s *ColumnDefListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ddlParserCOMMA, i)
}

func (s *ColumnDefListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnDefListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnDefListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ddlListener); ok {
		listenerT.EnterColumnDefList(s)
	}
}

func (s *ColumnDefListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ddlListener); ok {
		listenerT.ExitColumnDefList(s)
	}
}

func (s *ColumnDefListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ddlVisitor:
		return t.VisitColumnDefList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ddlParser) ColumnDefList() (localctx IColumnDefListContext) {
	localctx = NewColumnDefListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ddlParserRULE_columnDefList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.ColumnDef()
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ddlParserCOMMA {
		{
			p.SetState(31)
			p.Match(ddlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(32)
			p.ColumnDef()
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumnDefContext is an interface to support dynamic dispatch.
type IColumnDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	DataType() IDataTypeContext

	// IsColumnDefContext differentiates from other interfaces.
	IsColumnDefContext()
}

type ColumnDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnDefContext() *ColumnDefContext {
	var p = new(ColumnDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ddlParserRULE_columnDef
	return p
}

func InitEmptyColumnDefContext(p *ColumnDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ddlParserRULE_columnDef
}

func (*ColumnDefContext) IsColumnDefContext() {}

func NewColumnDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnDefContext {
	var p = new(ColumnDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ddlParserRULE_columnDef

	return p
}

func (s *ColumnDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ddlParserIDENTIFIER, 0)
}

func (s *ColumnDefContext) DataType() IDataTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *ColumnDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ddlListener); ok {
		listenerT.EnterColumnDef(s)
	}
}

func (s *ColumnDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ddlListener); ok {
		listenerT.ExitColumnDef(s)
	}
}

func (s *ColumnDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ddlVisitor:
		return t.VisitColumnDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ddlParser) ColumnDef() (localctx IColumnDefContext) {
	localctx = NewColumnDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ddlParserRULE_columnDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(ddlParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.DataType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataTypeContext is an interface to support dynamic dispatch.
type IDataTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	AllINT() []antlr.TerminalNode
	INT(i int) antlr.TerminalNode
	COMMA() antlr.TerminalNode
	VARCHAR2() antlr.TerminalNode

	// IsDataTypeContext differentiates from other interfaces.
	IsDataTypeContext()
}

type DataTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataTypeContext() *DataTypeContext {
	var p = new(DataTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ddlParserRULE_dataType
	return p
}

func InitEmptyDataTypeContext(p *DataTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ddlParserRULE_dataType
}

func (*DataTypeContext) IsDataTypeContext() {}

func NewDataTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataTypeContext {
	var p = new(DataTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ddlParserRULE_dataType

	return p
}

func (s *DataTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DataTypeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ddlParserNUMBER, 0)
}

func (s *DataTypeContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(ddlParserINT)
}

func (s *DataTypeContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(ddlParserINT, i)
}

func (s *DataTypeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ddlParserCOMMA, 0)
}

func (s *DataTypeContext) VARCHAR2() antlr.TerminalNode {
	return s.GetToken(ddlParserVARCHAR2, 0)
}

func (s *DataTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ddlListener); ok {
		listenerT.EnterDataType(s)
	}
}

func (s *DataTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ddlListener); ok {
		listenerT.ExitDataType(s)
	}
}

func (s *DataTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ddlVisitor:
		return t.VisitDataType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ddlParser) DataType() (localctx IDataTypeContext) {
	localctx = NewDataTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ddlParserRULE_dataType)
	var _la int

	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ddlParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Match(ddlParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ddlParserT__0 {
			{
				p.SetState(42)
				p.Match(ddlParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(43)
				p.Match(ddlParserINT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == ddlParserCOMMA {
				{
					p.SetState(44)
					p.Match(ddlParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(45)
					p.Match(ddlParserINT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(48)
				p.Match(ddlParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case ddlParserVARCHAR2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.Match(ddlParserVARCHAR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.Match(ddlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.Match(ddlParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.Match(ddlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
