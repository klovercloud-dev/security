package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
)

var tokens map[string][]v1.Token

type tokenMock struct {
}

func (t tokenMock) GetByToken(token string) v1.Token {
	panic("implement me")
}

func (t tokenMock) Store(token v1.Token) error {
	if tokens == nil {
		tokens = make(map[string][]v1.Token)
		tokens[token.Uid] = append(tokens[token.Uid], token)
		return nil
	}
	if token.Type == enums.REGULAR_TOKEN {
		oldTokens := tokens[token.Uid]

		for i, each := range oldTokens {
			if each.Type == enums.REGULAR_TOKEN {
				each = token
				oldTokens[i] = each
				tokens[token.Uid] = oldTokens
				return nil
			}
		}
	}
	tokens[token.Uid] = append(tokens[token.Uid], token)
	return nil
}

func (t tokenMock) Delete(uid string) error {
	panic("implement me")
}

func (t tokenMock) Update(token string, refreshToken string, existingToken string) error {
	panic("implement me")
}

// NewTokenMock returns service.Token type service
func NewTokenMock() service.Token {
	return &tokenMock{}
}
