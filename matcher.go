package golexer

import (
	"bytes"
	"errors"

	"unicode"
)

// 未知字符
type UnknownMatcher struct {
}

func (self *UnknownMatcher) Match(tz *Tokenizer) (*Token, error) {

	if tz.Current() == 0 {
		return nil, nil
	}

	begin := tz.Index()

	tz.ConsumeOne()

	return NewToken(self, tz, tz.StringRange(begin, tz.Index())), nil
}

// 空白字符
type WhitespaceMatcher struct {
}

func isWhiteSpace(c rune) bool {
	return c == ' ' || c == '\t'
}

func (self *WhitespaceMatcher) Match(tz *Tokenizer) (*Token, error) {

	var count int
	for {

		c := tz.Peek(count)

		if !isWhiteSpace(c) {
			break
		}

		count++

	}

	if count == 0 {
		return nil, nil
	}

	tz.ConsumeMulti(count)

	return NewToken(self, tz, ""), nil
}

// 行结束
type LineEndMatcher struct {
}

func (self *LineEndMatcher) Match(tz *Tokenizer) (*Token, error) {

	var count int
	for {

		c := tz.Peek(count)

		if c == '\n' {
			tz.increaseLine()
		} else if c == '\r' {

		} else {
			break
		}

		count++

	}

	if count == 0 {
		return nil, nil
	}

	tz.ConsumeMulti(count)

	return NewToken(self, tz, ""), nil
}

// 整形，浮点数
type NumeralMatcher struct {
}

func (self *NumeralMatcher) Match(tz *Tokenizer) (*Token, error) {

	if !unicode.IsDigit(tz.Current()) && tz.Current() != '-' {
		return nil, nil
	}

	begin := tz.Index()

	var maybeFloat bool

	for {

		tz.ConsumeOne()

		if !unicode.IsDigit(tz.Current()) {

			if tz.Current() == '.' {
				maybeFloat = true
				break
			}
		}

	}

	if maybeFloat {
		for i := 0; ; i++ {

			tz.ConsumeOne()

			if !unicode.IsDigit(tz.Current()) {

				// .之后的第一个字符居然不是数字
				if i == 0 {
					return nil, errors.New("invalid numeral")
				}

				break

			}

		}
	}

	return NewToken(self, tz, tz.StringRange(begin, tz.Index())), nil
}

// 字符串
type StringMatcher struct {
	builder bytes.Buffer
}

func (self *StringMatcher) Match(tz *Tokenizer) (*Token, error) {

	if tz.Current() != '"' && tz.Current() != '\'' {
		return nil, nil
	}

	tz.ConsumeOne()

	var escaping bool

	self.builder.Reset()

	for {

		if escaping {
			switch tz.Current() {
			case 'n':
				self.builder.WriteRune('\n')
			case 'r':
				self.builder.WriteRune('\r')
			default:
				self.builder.WriteRune('\\')
				self.builder.WriteRune(tz.Current())
			}

			escaping = false
		} else {
			if tz.Current() == '\\' {
				escaping = true
			} else {
				self.builder.WriteRune(tz.Current())
			}
		}

		tz.ConsumeOne()

		if tz.Current() == '\n' ||
			tz.Current() == 0 ||
			tz.Current() == '"' ||
			tz.Current() == '\'' {
			break
		}

	}

	tz.ConsumeOne()

	return NewToken(self, tz, self.builder.String()), nil
}

// 标识符
type IdentifierMatcher struct {
}

func (self *IdentifierMatcher) Match(tz *Tokenizer) (*Token, error) {

	if !unicode.IsLetter(tz.Current()) && tz.Current() != '_' {
		return nil, nil
	}

	begin := tz.Index()

	for {

		tz.ConsumeOne()

		if !unicode.IsLetter(tz.Current()) && tz.Current() != '_' {
			break
		}

	}

	return NewToken(self, tz, tz.StringRange(begin, tz.index)), nil
}

// 操作符，分隔符，关键字
type SignMatcher struct {
	word []rune
}

func (self *SignMatcher) Match(tz *Tokenizer) (*Token, error) {

	if (tz.Count() - tz.Index()) < len(self.word) {
		return nil, nil
	}

	for i, c := range self.word {

		if tz.Peek(i) != c {
			return nil, nil
		}

	}

	tz.ConsumeMulti(len(self.word))

	return NewToken(self, tz, string(self.word)), nil
}

func NewSignMatcher(word string) *SignMatcher {
	return &SignMatcher{
		word: []rune(word),
	}
}

// #开头的行注释
type UnixStyleCommentMatcher struct {
}

func (self *UnixStyleCommentMatcher) Match(tz *Tokenizer) (*Token, error) {
	if tz.Current() != '#' {
		return nil, nil
	}

	tz.ConsumeOne()

	begin := tz.Index()

	for {

		tz.ConsumeOne()

		if tz.Current() == '\n' || tz.Current() == 0 {
			break
		}

	}

	return NewToken(self, tz, tz.StringRange(begin, tz.index)), nil
}
