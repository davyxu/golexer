package golexer

import (
	"errors"
)

type Lexer struct {
	matchers []matcherMeta

	comm chan tokenAndError

	running bool
}

type matcherMeta struct {
	m      TokenMatcher
	ignore bool
}

type tokenAndError struct {
	tk  *Token
	err error
}

var eofToken = tokenAndError{NewToken(nil, nil, "EOF", ""), nil}

func (self *Lexer) VisitMatcher(callback func(TokenMatcher) bool) {

	for _, m := range self.matchers {
		if !callback(m.m) {
			return
		}
	}
}

func (self *Lexer) MatcherString(id int) string {
	m := self.MatcherByID(id)
	if m != nil {
		return m.String()
	}

	return ""
}

func (self *Lexer) MatcherByID(id int) TokenMatcher {
	for _, m := range self.matchers {
		if m.m.ID() == id {
			return m.m
		}
	}

	return nil
}

// 添加一个匹配器，如果结果匹配，返回token
func (self *Lexer) AddMatcher(m TokenMatcher) {
	self.matchers = append(self.matchers, matcherMeta{
		m:      m,
		ignore: false,
	})
}

// 添加一个匹配器，如果结果匹配，直接忽略匹配内容
func (self *Lexer) AddIgnoreMatcher(m TokenMatcher) {
	self.matchers = append(self.matchers, matcherMeta{
		m:      m,
		ignore: true,
	})
}

func (self *Lexer) Start(src string) {

	if self.comm != nil {
		close(self.comm)
	}

	self.comm = make(chan tokenAndError)

	self.running = true

	go self.tokenWorker(src)
}

func (self *Lexer) Read() (*Token, error) {

	if !self.running {
		return eofToken.tk, nil
	}

	if self.comm == nil {
		return eofToken.tk, errors.New("call 'Start' first")
	}

	te := <-self.comm

	if te.err != nil || te.tk.MatcherID() == 0 {
		self.running = false
	}

	return te.tk, te.err
}

func (self *Lexer) tokenWorker(src string) {

	tz := NewTokenizer(src, self)

	if len(self.matchers) > 0 {

		for !tz.EOF() {

			for _, mm := range self.matchers {

				token, err := mm.m.Match(tz)

				if err != nil {
					self.comm <- tokenAndError{NewToken(nil, tz, err.Error(), ""), err}
					return
				}

				if token == nil {
					continue
				}

				if mm.ignore {
					break
				}

				self.comm <- tokenAndError{token, nil}

				// 重新从matcher开始检查
				break

			}
		}
	}

	self.comm <- eofToken
}

func NewLexer() *Lexer {

	return &Lexer{}

}
