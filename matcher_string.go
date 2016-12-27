package golexer

import "bytes"

// 字符串
type StringMatcher struct {
	baseMatcher
	builder bytes.Buffer
}

func (self *StringMatcher) Match(tz *Tokenizer) (*Token, error) {

	if tz.Current() != '"' && tz.Current() != '\'' {
		return nil, nil
	}

	beginChar := tz.Current()

	begin := tz.Index()

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
			case '"', '\'':
				self.builder.WriteRune(tz.Current())
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

		if !escaping && tz.Current() == beginChar {
			break
		}

		if tz.Current() == '\n' ||
			tz.Current() == 0 {
			break
		}

	}

	tz.ConsumeOne()

	return NewToken(self, tz, self.builder.String(), tz.StringRange(begin, tz.Index())), nil
}

func NewStringMatcher(id int) TokenMatcher {
	return &StringMatcher{
		baseMatcher: baseMatcher{id},
	}
}
