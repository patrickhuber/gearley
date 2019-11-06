package grammars

type TokenType interface {
	Id() string
}

type tokenType struct {
	id string
}

func NewTokenType(id string) TokenType {
	return &tokenType{
		id: id,
	}
}

func (t *tokenType) Id() string {
	return t.id
}
