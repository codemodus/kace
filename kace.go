// Package kace provides common case conversion functions which take into
// consideration common initialisms.
package kace

import (
	"unicode"
)

var (
	ciMaxLen int
	// github.com/golang/lint/blob/master/lint.go
	// TODO: consider using tree for lookup to minimize allocations.
	ci = map[string]bool{
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
		"XSRF":  true,
		"XSS":   true,
	}
)

func init() {
	for k := range ci {
		if len(k) > ciMaxLen {
			ciMaxLen = len(k)
		}
	}
}

// Camel returns a camel cased string.
func Camel(s string, ucFirst bool) string {
	tmpBuf := make([]rune, 0, ciMaxLen)
	buf := make([]rune, 0, len(s))

	for i := 0; i < len(s); i++ {
		tmpBuf = tmpBuf[:0]
		if unicode.IsLetter(rune(s[i])) {
			if i == 0 || !unicode.IsLetter(rune(s[i-1])) {
				for n := i; n < len(s) && n-i < ciMaxLen; n++ {
					tmpBuf = append(tmpBuf, unicode.ToUpper(rune(s[n])))
					if n < len(s)-1 && !unicode.IsLetter(rune(s[n+1])) && !unicode.IsDigit(rune(s[n+1])) {
						break
					}
				}
				if ((i == 0 && ucFirst) || i > 0) && ci[string(tmpBuf)] {
					buf = append(buf, tmpBuf...)
					i += len(tmpBuf)
					continue
				}
			}

			if i == 0 && ucFirst || i > 0 && !unicode.IsLetter(rune(s[i-1])) {
				buf = append(buf, unicode.ToUpper(rune(s[i])))
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

// Snake returns a snake cased string.
func Snake(s string) string {
	return delimitedCase(s, '_', false)
}

// SnakeUpper returns a snake cased string with all upper case letters.
func SnakeUpper(s string) string {
	return delimitedCase(s, '_', true)
}

// Kebab returns a kebab cased string.
func Kebab(s string) string {
	return delimitedCase(s, '-', false)
}

// KebabUpper returns a kebab cased string with all upper case letters.
func KebabUpper(s string) string {
	return delimitedCase(s, '-', true)
}

// Snake returns a snake cased string.
func delimitedCase(s string, delim rune, upper bool) string {
	buf := make([]rune, 0, len(s)*2)

	for i := len(s); i > 0; i-- {
		if unicode.IsLetter(rune(s[i-1])) {
			if i < len(s) && unicode.IsUpper(rune(s[i])) {
				if i > 1 && unicode.IsLower(rune(s[i-1])) || i < len(s)-2 && unicode.IsLower(rune(s[i+1])) {
					buf = append(buf, delim)
				}
			}
			if upper {
				buf = append(buf, unicode.ToUpper(rune(s[i-1])))
			} else {
				buf = append(buf, unicode.ToLower(rune(s[i-1])))
			}
		} else if unicode.IsDigit(rune(s[i-1])) {
			if i == len(s) || i == 1 || unicode.IsDigit(rune(s[i])) {
				buf = append(buf, rune(s[i-1]))
			} else {
				buf = append(buf, delim, rune(s[i-1]))
			}
		} else {
			if i == len(s) {
				continue
			}
			buf = append(buf, delim)
		}
	}
	return string(reverse(buf))
}

func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
