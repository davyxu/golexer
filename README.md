# golexer

可自定义的词法解析器

# 特性

* 支持数值，字符串，注释，标识符等的内建匹配

* 可自定义匹配器来拾取需要的token

* 高性能并发匹配

```golang
	l := NewLexer()

	// 匹配顺序从高到低

	l.AddMatcher(new(NumeralMatcher))
	l.AddMatcher(new(StringMatcher))

	l.AddIgnoreMatcher(new(WhitespaceMatcher))
	l.AddIgnoreMatcher(new(LineEndMatcher))
	l.AddIgnoreMatcher(new(UnixStyleCommentMatcher))

	l.AddMatcher(new(IdentifierMatcher))

	l.AddMatcher(new(UnknownMatcher))

	l.Start(`"a" 
	123.3
	_id # comment
	
	'b'
	
	
	`)

	for {

		tk, err := l.Read()

		if err != nil {
			t.Error(err)
			break
		}

		if tk == nil {
			break
		}

		t.Log(tk.String())
	}
```


# 备注

感觉不错请fork和star, 谢谢!

博客: http://www.cppblog.com/sunicdavy

知乎: http://www.zhihu.com/people/xu-bo-62-87

邮箱: sunicdavy@qq.com
