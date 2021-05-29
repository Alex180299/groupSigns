package test

import (
	"github.com/Alex180299/groupSigns/service"
	"testing"
)

func TestParseParentheses(t *testing.T) {
	code := "()(())"
	lexer := &service.LexerService{}
	tokens := lexer.ParseCode(code)

	parser := &service.ParserService{}
	err := parser.ParseTokens(tokens)

	if err != nil {
		t.Error("Error parsing tokens:", err)
	}
}

func TestParseSquareBrackets(t *testing.T) {
	code := "[][[]]"
	lexer := &service.LexerService{}
	tokens := lexer.ParseCode(code)

	parser := &service.ParserService{}
	err := parser.ParseTokens(tokens)

	if err != nil {
		t.Error("Error parsing tokens:", err)
	}
}

func TestParseBrackets(t *testing.T) {
	code := "{}{{}}"
	lexer := &service.LexerService{}
	tokens := lexer.ParseCode(code)

	parser := &service.ParserService{}
	err := parser.ParseTokens(tokens)

	if err != nil {
		t.Error("Error parsing tokens:", err)
	}
}

func TestParseErrorOpen(t *testing.T) {
	code := "{}{{}}{"
	lexer := &service.LexerService{}
	tokens := lexer.ParseCode(code)

	parser := &service.ParserService{}
	err := parser.ParseTokens(tokens)

	if err == nil {
		t.Error("Expected error parsing tokens:", err)
	}
}

func TestParseErrorClose(t *testing.T) {
	code := "{}{{}}}"
	lexer := &service.LexerService{}
	tokens := lexer.ParseCode(code)

	parser := &service.ParserService{}
	err := parser.ParseTokens(tokens)

	if err == nil {
		t.Error("Expected error parsing tokens:", err)
	}
}

func TestParseSigns(t *testing.T) {
	code := "({}{{}})[({({[({}{()})]})})]"
	lexer := &service.LexerService{}
	tokens := lexer.ParseCode(code)

	parser := &service.ParserService{}
	err := parser.ParseTokens(tokens)

	if err != nil {
		t.Error("Expected error parsing tokens:", err)
	}
}
