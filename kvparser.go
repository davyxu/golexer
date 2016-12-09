package golexer

// 自定义的token id
const (
	pbtToken_EOF = iota
	pbtToken_WhiteSpace
	pbtToken_Identifier
	pbtToken_Numeral
	pbtToken_String
	pbtToken_Comma
	pbtToken_Unknown
)

// k:v 与pbt格式互换,详见kvparser_test.go

func ParseKV(str string, callback func(string, string) bool) (errRet error) {

	l := NewLexer()

	l.AddMatcher(NewNumeralMatcher(pbtToken_Numeral))
	l.AddMatcher(NewStringMatcher(pbtToken_String))

	l.AddIgnoreMatcher(NewWhiteSpaceMatcher(pbtToken_WhiteSpace))
	l.AddMatcher(NewSignMatcher(pbtToken_Comma, ":"))
	l.AddMatcher(NewIdentifierMatcher(pbtToken_Identifier))
	l.AddMatcher(NewUnknownMatcher(pbtToken_Unknown))

	l.Start(str)

	p := NewParser(l)

	defer ErrorCatcher(func(err error) {

		errRet = err

	})

	p.NextToken()

	for p.TokenID() != pbtToken_EOF {

		if p.TokenID() != pbtToken_Identifier {
			panic("expect identifier")
		}

		key := p.TokenValue()

		p.NextToken()

		if p.TokenID() != pbtToken_Comma {
			panic("expect comma")
		}

		p.NextToken()

		value := p.TokenValue()

		if !callback(key, value) {
			return nil
		}

		p.NextToken()

	}

	return nil
}
