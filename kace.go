// Package kace provides common case conversion functions which take into
// consideration common initialisms.
package kace

import (
	"unicode"
)

const (
	kebabDelim = '-'
	snakeDelim = '_'
)

var (
	ciTrie = newTrie(ciMap)
)

// Camel returns a camelCased string.
func Camel(s string) string {
	return camel(s, false)
}

// Pascal returns a PascalCased string.
func Pascal(s string) string {
	return camel(s, true)
}

func camel(s string, ucFirst bool) string {
	tmpBuf := make([]rune, 0, ciTrie.maxDepth)
	buf := make([]rune, 0, len(s))

	for i := 0; i < len(s); i++ {
		tmpBuf = tmpBuf[:0]
		if unicode.IsLetter(rune(s[i])) {
			if i == 0 || !unicode.IsLetter(rune(s[i-1])) {
				for n := i; n < len(s) && n-i < ciTrie.maxDepth; n++ {
					tmpBuf = append(tmpBuf, unicode.ToUpper(rune(s[n])))
					if n < len(s)-1 && !unicode.IsLetter(rune(s[n+1])) && !unicode.IsDigit(rune(s[n+1])) {
						break
					}
				}
				if ((i == 0 && ucFirst) || i > 0) && ciTrie.find(tmpBuf) {
					buf = append(buf, tmpBuf...)
					i += len(tmpBuf)
					continue
				}
			}

			if i == 0 && ucFirst || i > 0 && !unicode.IsLetter(rune(s[i-1])) {
				buf = append(buf, unicode.ToUpper(rune(s[i])))
			} else if i == 0 && !ucFirst {
				buf = append(buf, unicode.ToLower(rune(s[i])))
			} else {
				buf = append(buf, rune(s[i]))
			}
		}

		if unicode.IsDigit(rune(s[i])) {
			buf = append(buf, rune(s[i]))
		}
	}
	return string(buf)
}

// Snake returns a snake_cased string with all lowercase letters.
func Snake(s string) string {
	return delimitedCase(s, snakeDelim, false)
}

// SnakeUpper returns a SNAKE_CASED string with all upper case letters.
func SnakeUpper(s string) string {
	return delimitedCase(s, snakeDelim, true)
}

// Kebab returns a kebab-cased string with all lowercase letters.
func Kebab(s string) string {
	return delimitedCase(s, kebabDelim, false)
}

// KebabUpper returns a KEBAB-CASED string with all upper case letters.
func KebabUpper(s string) string {
	return delimitedCase(s, kebabDelim, true)
}

func delimitedCase(s string, delim rune, upper bool) string {
	buf := make([]rune, 0, len(s)*2)

	for i := len(s); i > 0; i-- {
		switch {
		case unicode.IsLetter(rune(s[i-1])):
			if i < len(s) && unicode.IsUpper(rune(s[i])) {
				if i > 1 && unicode.IsLower(rune(s[i-1])) || i < len(s)-2 && unicode.IsLower(rune(s[i+1])) {
					buf = append(buf, delim)
				}
			}

			buf = appendCased(buf, upper, rune(s[i-1]))

		case unicode.IsDigit(rune(s[i-1])):
			if i == len(s) || i == 1 || unicode.IsDigit(rune(s[i])) {
				buf = append(buf, rune(s[i-1]))
				continue
			}

			buf = append(buf, delim, rune(s[i-1]))

		default:
			if i == len(s) {
				continue
			}

			buf = append(buf, delim)
		}
	}

	reverse(buf)

	return string(buf)
}

func appendCased(rs []rune, upper bool, r rune) []rune {
	if upper {
		rs = append(rs, unicode.ToUpper(r))
		return rs
	}

	rs = append(rs, unicode.ToLower(r))

	return rs
}

func reverse(s []rune) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

var (
	// github.com/golang/lint/blob/master/lint.go
	ciMap = map[string]bool{
		"ACL":   true,
		"API":   true,
		"ASCII": true,
		"CPU":   true,
		"CSS":   true,
		"DNS":   true,
		"EOF":   true,
		"GUID":  true,
		"HTML":  true,
		"HTTP":  true,
		"HTTPS": true,
		"ID":    true,
		"IP":    true,
		"JSON":  true,
		"LHS":   true,
		"QPS":   true,
		"RAM":   true,
		"RHS":   true,
		"RPC":   true,
		"SLA":   true,
		"SMTP":  true,
		"SQL":   true,
		"SSH":   true,
		"TCP":   true,
		"TLS":   true,
		"TTL":   true,
		"UDP":   true,
		"UI":    true,
		"UID":   true,
		"UUID":  true,
		"URI":   true,
		"URL":   true,
		"UTF8":  true,
		"VM":    true,
		"XML":   true,
		"XMPP":  true,
		"XSRF":  true,
		"XSS":   true,
	}
)
