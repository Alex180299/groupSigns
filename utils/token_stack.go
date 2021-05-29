package utils

import "github.com/Alex180299/groupSigns/model"

type TokenStack []model.TokenType

func (s *TokenStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *TokenStack) Push(token model.TokenType) {
	*s = append(*s, token)
}

func (s *TokenStack) Pop() model.TokenType {
	if s.IsEmpty() {
		return model.None
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}
