package golexer

import "reflect"

// #开头的行注释
type UnixStyleCommentMatcher struct {
	baseMatcher
}

func (self *UnixStyleCommentMatcher) String() string {
	return reflect.TypeOf(self).Elem().Name()
}

func (self *UnixStyleCommentMatcher) Match(tz *Tokenizer) (Token, error) {
	if tz.Current() != '#' {
		return EmptyToken, nil
	}

	tz.ConsumeOne()

	begin := tz.Index()

	for {

		if tz.Current() == '\n' || tz.Current() == '\r' || tz.Current() == 0 {
			break
		}

		tz.ConsumeOne()

	}

	return NewToken(self, tz, tz.StringRange(begin, tz.index), ""), nil
}

func NewUnixStyleCommentMatcher(id int) TokenMatcher {
	return &UnixStyleCommentMatcher{
		baseMatcher{id},
	}
}
