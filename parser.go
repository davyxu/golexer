package golexer

import "errors"

type Parser struct {
	lexer *Lexer

	curr *Token

	errFunc func(error)
}

func (self *Parser) Lexer() *Lexer {
	return self.lexer
}

func (self *Parser) Expect(id int) Token {

	if self.TokenID() != id {
		panic(errors.New("Expect " + self.lexer.MatcherString(id)))
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

func (self *Parser) MatcherName() string {
	return self.curr.MatcherName()
}

func (self *Parser) MatcherString() string {
	return self.curr.MatcherString()
}

func (self *Parser) TokenRaw() string {

	return self.curr.Raw()
}

func (self *Parser) TokenPos() (int, int) {
	return self.curr.line, self.curr.index
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
	case nil:
	default:
		panic(err)
	}
}

func NewParser(l *Lexer) *Parser {

	return &Parser{
		lexer: l,
	}

}
