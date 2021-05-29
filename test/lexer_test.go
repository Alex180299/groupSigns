package test

import (
	"github.com/Alex180299/groupSigns/model"
	"testing"

	"github.com/Alex180299/groupSigns/service"
)

func verifyTokens(expected, actual []model.TokenType) bool {
	size := len(expected)
	for i := 0; i < size; i++ {
		if actual[i] != expected[i] {
			return false
		}
	}

	return true
}

func TestLexerParenthesis(t *testing.T) {
	code := "()(())"
	lexer := &service.LexerService{}

	tokensResult := lexer.ParseCode(code)

	expectedLength := len(code)
	if len(tokensResult) != expectedLength {
		t.Error("Error in length of tokens")
	}

	expectedTokens := make([]model.TokenType, expectedLength)
	expectedTokens[0] = model.LeftParentheses
	expectedTokens[1] = model.RightParentheses
	expectedTokens[2] = model.LeftParentheses
	expectedTokens[3] = model.LeftParentheses
	expectedTokens[4] = model.RightParentheses
	expectedTokens[5] = model.RightParentheses

	if !verifyTokens(expectedTokens, tokensResult) {
		t.Error("Error to verify tokens")
	}
}

func TestLexerSquareBracket(t *testing.T) {
	code := "[][[]]"
	lexer := &service.LexerService{}

	tokensResult := lexer.ParseCode(code)

	expectedLength := len(code)
	if len(tokensResult) != expectedLength {
		t.Error("Error in length of tokens")
	}

	expectedTokens := make([]model.TokenType, expectedLength)
	expectedTokens[0] = model.LeftSquareBracket
	expectedTokens[1] = model.RightSquareBracket
	expectedTokens[2] = model.LeftSquareBracket
	expectedTokens[3] = model.LeftSquareBracket
	expectedTokens[4] = model.RightSquareBracket
	expectedTokens[5] = model.RightSquareBracket

	if !verifyTokens(expectedTokens, tokensResult) {
		t.Error("Error to verify tokens")
	}
}

func TestLexerBracket(t *testing.T) {
	code := "{}{{}}"
	lexer := &service.LexerService{}

	tokensResult := lexer.ParseCode(code)

	expectedLength := len(code)
	if len(tokensResult) != expectedLength {
		t.Error("Error in length of tokens")
	}

	expectedTokens := make([]model.TokenType, expectedLength)
	expectedTokens[0] = model.LeftBracket
	expectedTokens[1] = model.RightBracket
	expectedTokens[2] = model.LeftBracket
	expectedTokens[3] = model.LeftBracket
	expectedTokens[4] = model.RightBracket
	expectedTokens[5] = model.RightBracket

	if !verifyTokens(expectedTokens, tokensResult) {
		t.Error("Error to verify tokens")
	}
}

func TestLexerAllSigns(t *testing.T) {
	code := "([{}][()])"
	lexer := &service.LexerService{}

	tokensResult := lexer.ParseCode(code)

	expectedLength := len(code)
	if len(tokensResult) != expectedLength {
		t.Error("Error in length of tokens")
	}

	expectedTokens := make([]model.TokenType, expectedLength)
	expectedTokens[0] = model.LeftParentheses
	expectedTokens[1] = model.LeftSquareBracket
	expectedTokens[2] = model.LeftBracket
	expectedTokens[3] = model.RightBracket
	expectedTokens[4] = model.RightSquareBracket
	expectedTokens[5] = model.LeftSquareBracket
	expectedTokens[6] = model.LeftParentheses
	expectedTokens[7] = model.RightParentheses
	expectedTokens[8] = model.RightSquareBracket
	expectedTokens[9] = model.RightParentheses

	if !verifyTokens(expectedTokens, tokensResult) {
		t.Error("Error to verify tokens")
	}
}
