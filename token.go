package golexer

import (
	"fmt"
	"reflect"
	"strconv"
)

type Token struct {
	value   string
	matcher TokenMatcher
	line    int
}

func (self *Token) MatcherID() int {

	if self == nil || self.matcher == nil {
		return 0
	}

	return self.matcher.ID()
}

func (self *Token) Value() string {
	if self == nil || self.matcher == nil {
		return ""
	}

	return self.value
}

func (self *Token) ToFloat32() float32 {
	v, err := strconv.ParseFloat(self.value, 32)

	if err != nil {
		return 0
	}

	return float32(v)
}

func (self *Token) MatcherName() string {
	if self == nil || self.matcher == nil {
		return ""
	}

	return reflect.TypeOf(self.matcher).Elem().Name()
}

func (self *Token) String() string {

	if self == nil {
		return ""
	}

	return fmt.Sprintf("line: %d id:%d matcher: %s  value:%s", self.line, self.MatcherID(), self.MatcherName(), self.value)
}

func NewToken(m TokenMatcher, tz *Tokenizer, v string) *Token {

	return &Token{
		value:   v,
		line:    tz.Line(),
		matcher: m,
	}
}
