package service

import (
	"errors"
	"github.com/Alex180299/groupSigns/model"
	"github.com/Alex180299/groupSigns/utils"
)

type ParserService struct{}

func (parser *ParserService) ParseTokens(tokens []model.TokenType) error {
	stack := utils.TokenStack{}
	for _, token := range tokens {
		if token == model.LeftParentheses || token == model.LeftSquareBracket || token == model.LeftBracket {
			stack.Push(token)
		} else if token == model.RightParentheses {
			tokenPopped := stack.Pop()
			if tokenPopped != model.LeftParentheses {
				return errors.New("error parenthesis expected")
			}
		} else if token == model.RightSquareBracket {
			tokenPopped := stack.Pop()
			if tokenPopped != model.LeftSquareBracket {
				return errors.New("error square bracket expected")
			}
		} else if token == model.RightBracket {
			tokenPopped := stack.Pop()
			if tokenPopped != model.LeftBracket {
				return errors.New("error bracket expected")
			}
		}
	}

	if stack.IsEmpty() {
		return nil
	} else {
		return errors.New("error stack isn`t empty")
	}
}