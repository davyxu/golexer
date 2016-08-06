package golexer

import (
	"fmt"
	"reflect"
	"strconv"
)

type Token struct {
	value   string
	matcher TokenMatcher
}

func (self *Token) ToFloat32() float32 {
	v, err := strconv.ParseFloat(self.value, 32)

	if err != nil {
		return 0
	}

	return float32(v)
}

func (self *Token) MatcherName() string {
	return reflect.TypeOf(self.matcher).Elem().Name()
}

func (self *Token) String() string {
	return fmt.Sprintf("matcher: %s  value:%s", self.MatcherName(), self.value)
}

func NewToken(m TokenMatcher, v string) *Token {

	return &Token{
		value:   v,
		matcher: m,
	}
}
