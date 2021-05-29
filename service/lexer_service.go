package service

import (
	"github.com/Alex180299/groupSigns/model"
)

type LexerService struct{}

func (lexer *LexerService) ParseCode(code string) []model.TokenType {
	codeSize := len(code)
	tokens := make([]model.TokenType, 0)

	for i := 0; i < codeSize; i++ {
		if code[i] == 40 {
			tokens = append(tokens, model.LeftParentheses)
		} else if code[i] == 41 {
			tokens = append(tokens, model.RightParentheses)
		} else if code[i] == 91 {
			tokens = append(tokens, model.LeftSquareBracket)
		} else if code[i] == 93 {
			tokens = append(tokens, model.RightSquareBracket)
		} else if code[i] == 123 {
			tokens = append(tokens, model.LeftBracket)
		} else if code[i] == 125 {
			tokens = append(tokens, model.RightBracket)
		}
	}

	return tokens
}
