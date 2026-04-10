// Code generated from ddl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ddlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var DdlLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ddllexerLexerInit() {
	staticData := &DdlLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "", "", "", "", "", "';'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "CREATE", "TABLE", "NUMBER", "VARCHAR2", "DATE", "SEMICOLON",
		"COMMA", "IDENTIFIER", "INT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "CREATE", "TABLE", "NUMBER", "VARCHAR2", "DATE", "SEMICOLON",
		"COMMA", "IDENTIFIER", "INT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 86, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9,
		5, 9, 70, 8, 9, 10, 9, 12, 9, 73, 9, 9, 1, 10, 4, 10, 76, 8, 10, 11, 10,
		12, 10, 77, 1, 11, 4, 11, 81, 8, 11, 11, 11, 12, 11, 82, 1, 11, 1, 11,
		0, 0, 12, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 1, 0, 17, 2, 0, 67, 67, 99, 99, 2, 0, 82, 82, 114,
		114, 2, 0, 69, 69, 101, 101, 2, 0, 65, 65, 97, 97, 2, 0, 84, 84, 116, 116,
		2, 0, 66, 66, 98, 98, 2, 0, 76, 76, 108, 108, 2, 0, 78, 78, 110, 110, 2,
		0, 85, 85, 117, 117, 2, 0, 77, 77, 109, 109, 2, 0, 86, 86, 118, 118, 2,
		0, 72, 72, 104, 104, 2, 0, 68, 68, 100, 100, 3, 0, 65, 90, 95, 95, 97,
		122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3, 0, 9, 10,
		13, 13, 32, 32, 88, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 1, 25, 1, 0, 0, 0, 3, 27, 1, 0, 0, 0, 5, 29,
		1, 0, 0, 0, 7, 36, 1, 0, 0, 0, 9, 42, 1, 0, 0, 0, 11, 49, 1, 0, 0, 0, 13,
		58, 1, 0, 0, 0, 15, 63, 1, 0, 0, 0, 17, 65, 1, 0, 0, 0, 19, 67, 1, 0, 0,
		0, 21, 75, 1, 0, 0, 0, 23, 80, 1, 0, 0, 0, 25, 26, 5, 40, 0, 0, 26, 2,
		1, 0, 0, 0, 27, 28, 5, 41, 0, 0, 28, 4, 1, 0, 0, 0, 29, 30, 7, 0, 0, 0,
		30, 31, 7, 1, 0, 0, 31, 32, 7, 2, 0, 0, 32, 33, 7, 3, 0, 0, 33, 34, 7,
		4, 0, 0, 34, 35, 7, 2, 0, 0, 35, 6, 1, 0, 0, 0, 36, 37, 7, 4, 0, 0, 37,
		38, 7, 3, 0, 0, 38, 39, 7, 5, 0, 0, 39, 40, 7, 6, 0, 0, 40, 41, 7, 2, 0,
		0, 41, 8, 1, 0, 0, 0, 42, 43, 7, 7, 0, 0, 43, 44, 7, 8, 0, 0, 44, 45, 7,
		9, 0, 0, 45, 46, 7, 5, 0, 0, 46, 47, 7, 2, 0, 0, 47, 48, 7, 1, 0, 0, 48,
		10, 1, 0, 0, 0, 49, 50, 7, 10, 0, 0, 50, 51, 7, 3, 0, 0, 51, 52, 7, 1,
		0, 0, 52, 53, 7, 0, 0, 0, 53, 54, 7, 11, 0, 0, 54, 55, 7, 3, 0, 0, 55,
		56, 7, 1, 0, 0, 56, 57, 5, 50, 0, 0, 57, 12, 1, 0, 0, 0, 58, 59, 7, 12,
		0, 0, 59, 60, 7, 3, 0, 0, 60, 61, 7, 4, 0, 0, 61, 62, 7, 2, 0, 0, 62, 14,
		1, 0, 0, 0, 63, 64, 5, 59, 0, 0, 64, 16, 1, 0, 0, 0, 65, 66, 5, 44, 0,
		0, 66, 18, 1, 0, 0, 0, 67, 71, 7, 13, 0, 0, 68, 70, 7, 14, 0, 0, 69, 68,
		1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0,
		72, 20, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 76, 7, 15, 0, 0, 75, 74, 1,
		0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78,
		22, 1, 0, 0, 0, 79, 81, 7, 16, 0, 0, 80, 79, 1, 0, 0, 0, 81, 82, 1, 0,
		0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85,
		6, 11, 0, 0, 85, 24, 1, 0, 0, 0, 4, 0, 71, 77, 82, 1, 6, 0, 0,
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

// ddlLexerInit initializes any static state used to implement ddlLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewddlLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func DdlLexerInit() {
	staticData := &DdlLexerLexerStaticData
	staticData.once.Do(ddllexerLexerInit)
}

// NewddlLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewddlLexer(input antlr.CharStream) *ddlLexer {
	DdlLexerInit()
	l := new(ddlLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &DdlLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "ddl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ddlLexer tokens.
const (
	ddlLexerT__0       = 1
	ddlLexerT__1       = 2
	ddlLexerCREATE     = 3
	ddlLexerTABLE      = 4
	ddlLexerNUMBER     = 5
	ddlLexerVARCHAR2   = 6
	ddlLexerDATE       = 7
	ddlLexerSEMICOLON  = 8
	ddlLexerCOMMA      = 9
	ddlLexerIDENTIFIER = 10
	ddlLexerINT        = 11
	ddlLexerWS         = 12
)
