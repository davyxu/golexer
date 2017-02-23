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

func isKeyword(r rune, index int) bool {
	basic := unicode.IsLetter(r) || r == '_'
	if index == 0 {
		return basic
	}

	return basic || unicode.IsDigit(r)
}

func (self *KeywordMatcher) String() string {
	return fmt.Sprintf("%s('%s')", reflect.TypeOf(self).Elem().Name(), string(self.word))
}

func (self *KeywordMatcher) Match(tz *Tokenizer) (*Token, error) {

	if (tz.Count() - tz.Index()) < len(self.word) {
		return nil, nil
	}

	var index int

	for _, c := range self.word {

		if !isKeyword(c, index) {
			return nil, nil
		}

		if tz.Peek(index) != c {
			return nil, nil
		}

		index++

	}

	pc := tz.Peek(index)

	var needParser bool

	tz.lex.VisitMatcher(func(m TokenMatcher) bool {

		km, ok := m.(*KeywordMatcher)
		if ok {

			if km.word[0] == pc {
				needParser = true
				return false
			}

		}

		return true
	})

	if isKeyword(pc, index) && !needParser {
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

	for i, c := range self.word {
		if !isKeyword(c, i) {
			panic("not keyword")
		}
	}

	return self
}
