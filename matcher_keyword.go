package golexer

import (
	"fmt"
	"reflect"
	"unicode"
)

// 操作符，分隔符，关键字
type KeywordMatcher struct {
	baseMatcher
	word []rune
}

func isKeyword(r rune) bool {
	return unicode.IsLetter(r) ||
		r == '_'
}

func (self *KeywordMatcher) String() string {
	return fmt.Sprintf("%s(%s)", reflect.TypeOf(self).Elem().Name(), string(self.word))
}

func (self *KeywordMatcher) Match(tz *Tokenizer) (*Token, error) {

	if (tz.Count() - tz.Index()) < len(self.word) {
		return nil, nil
	}

	for i, c := range self.word {

		if !isKeyword(c) {
			return nil, nil
		}

		if tz.Peek(i) != c {
			return nil, nil
		}

	}

	pc := tz.Peek(len(self.word))

	var needParser bool

	tz.Lexer.VisitMatcher(func(m TokenMatcher) bool {

		km, ok := m.(*KeywordMatcher)
		if ok {

			if km.word[0] == pc {
				needParser = true
				return false
			}

		}

		return true
	})

	if isKeyword(pc) && !needParser {
		return nil, nil
	}

	tz.ConsumeMulti(len(self.word))

	return NewToken(self, tz, string(self.word), ""), nil
}

func NewKeywordMatcher(id int, word string) TokenMatcher {

	if len(word) == 0 {
		panic("empty string")
	}

	self := &KeywordMatcher{
		baseMatcher: baseMatcher{id},
		word:        []rune(word),
	}

	for _, c := range self.word {
		if !isKeyword(c) {
			panic("not keyword")
		}
	}

	return self
}
