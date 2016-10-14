package golexer

import (
	"errors"
)

type Parser struct {
	lexer *Lexer

	curr *Token

	errFunc func(error)
}

var (
	ErrUnexpectToken = errors.New("ErrUnexpectToken")
)

func (self *Parser) Lexer() *Lexer {
	return self.lexer
}

func (self *Parser) Expect(id int) Token {

	if self.TokenID() != id {
		panic(ErrUnexpectToken)
	}

	t := *self.curr

	self.NextToken()

	return t
}

func (self *Parser) NextToken() {

	token, err := self.lexer.Read()

	if err != nil {
		panic(err)
	}

	self.curr = token
}

func (self *Parser) TokenID() int {
	return self.curr.MatcherID()
}

func (self *Parser) TokenValue() string {
	return self.curr.Value()
}

func ErrorCatcher(errFunc func(error)) {

	err := recover()

	switch err.(type) {
	// 运行时错误
	case interface {
		RuntimeError()
	}:

		// 继续外抛， 方便调试
		panic(err)

	case error:
		errFunc(err.(error))
	}
}

func NewParser(l *Lexer) *Parser {

	return &Parser{
		lexer: l,
	}

}
