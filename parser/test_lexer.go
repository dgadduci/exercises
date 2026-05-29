// Code generated from Test.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

type TestLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var TestLexerLexerStaticData struct {
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

func testlexerLexerInit() {
	staticData := &TestLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'='", "", "", "", "'+'", "'-'", "'*'", "'/'",
	}
	staticData.SymbolicNames = []string{
		"", "", "ID", "FLOAT", "INT", "ADD", "MINUS", "MUL", "DIV", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "ID", "FLOAT", "INT", "ADD", "MINUS", "MUL", "DIV", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 9, 59, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1,
		1, 1, 1, 5, 1, 24, 8, 1, 10, 1, 12, 1, 27, 9, 1, 1, 2, 4, 2, 30, 8, 2,
		11, 2, 12, 2, 31, 1, 2, 1, 2, 4, 2, 36, 8, 2, 11, 2, 12, 2, 37, 1, 3, 4,
		3, 41, 8, 3, 11, 3, 12, 3, 42, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 8, 4, 8, 54, 8, 8, 11, 8, 12, 8, 55, 1, 8, 1, 8, 0, 0, 9, 1, 1,
		3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 1, 0, 4, 3, 0, 65,
		90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57,
		3, 0, 9, 10, 13, 13, 32, 32, 63, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 1, 19, 1, 0, 0, 0,
		3, 21, 1, 0, 0, 0, 5, 29, 1, 0, 0, 0, 7, 40, 1, 0, 0, 0, 9, 44, 1, 0, 0,
		0, 11, 46, 1, 0, 0, 0, 13, 48, 1, 0, 0, 0, 15, 50, 1, 0, 0, 0, 17, 53,
		1, 0, 0, 0, 19, 20, 5, 61, 0, 0, 20, 2, 1, 0, 0, 0, 21, 25, 7, 0, 0, 0,
		22, 24, 7, 1, 0, 0, 23, 22, 1, 0, 0, 0, 24, 27, 1, 0, 0, 0, 25, 23, 1,
		0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 4, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28,
		30, 7, 2, 0, 0, 29, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 29, 1, 0, 0,
		0, 31, 32, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 35, 5, 46, 0, 0, 34, 36,
		7, 2, 0, 0, 35, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0,
		37, 38, 1, 0, 0, 0, 38, 6, 1, 0, 0, 0, 39, 41, 7, 2, 0, 0, 40, 39, 1, 0,
		0, 0, 41, 42, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 8,
		1, 0, 0, 0, 44, 45, 5, 43, 0, 0, 45, 10, 1, 0, 0, 0, 46, 47, 5, 45, 0,
		0, 47, 12, 1, 0, 0, 0, 48, 49, 5, 42, 0, 0, 49, 14, 1, 0, 0, 0, 50, 51,
		5, 47, 0, 0, 51, 16, 1, 0, 0, 0, 52, 54, 7, 3, 0, 0, 53, 52, 1, 0, 0, 0,
		54, 55, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57, 1,
		0, 0, 0, 57, 58, 6, 8, 0, 0, 58, 18, 1, 0, 0, 0, 6, 0, 25, 31, 37, 42,
		55, 1, 6, 0, 0,
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

// TestLexerInit initializes any static state used to implement TestLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewTestLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TestLexerInit() {
	staticData := &TestLexerLexerStaticData
	staticData.once.Do(testlexerLexerInit)
}

// NewTestLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewTestLexer(input antlr.CharStream) *TestLexer {
	TestLexerInit()
	l := new(TestLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &TestLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Test.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TestLexer tokens.
const (
	TestLexerT__0  = 1
	TestLexerID    = 2
	TestLexerFLOAT = 3
	TestLexerINT   = 4
	TestLexerADD   = 5
	TestLexerMINUS = 6
	TestLexerMUL   = 7
	TestLexerDIV   = 8
	TestLexerWS    = 9
)
