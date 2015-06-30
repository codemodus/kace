// Package kace provides common case conversion functions which take into
// consideration common initialisms.
package kace

import (
	"unicode"
)

var (
	ciMaxLen int
	// github.com/golang/lint/blob/master/lint.go
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
func Camel(rs string, ucFirst bool) string {
	tmpBuf := make([]rune, 0, ciMaxLen)
	buf := make([]rune, 0, len(rs))

	for i := 0; i < len(rs); i++ {
		tmpBuf = tmpBuf[:0]
		if unicode.IsLetter(rune(rs[i])) {
			if i == 0 || !unicode.IsLetter(rune(rs[i-1])) {
				for n := i; n < len(rs) && n-i < ciMaxLen; n++ {
					tmpBuf = append(tmpBuf, unicode.ToUpper(rune(rs[n])))
					if n < len(rs)-1 && !unicode.IsLetter(rune(rs[n+1])) && !unicode.IsDigit(rune(rs[n+1])) {
						break
					}
				}
				if ci[string(tmpBuf)] {
					buf = append(buf, tmpBuf...)
					i += len(tmpBuf)
					continue
				}
			}

			if i == 0 && ucFirst || i > 0 && !unicode.IsLetter(rune(rs[i-1])) {
				buf = append(buf, unicode.ToUpper(rune(rs[i])))
			} else {
				buf = append(buf, rune(rs[i]))
			}
		}

		if unicode.IsDigit(rune(rs[i])) {
			buf = append(buf, rune(rs[i]))
		}
	}
	return string(buf)
}

// Snake returns a snake cased string.
func Snake(rs string) string {
	buf := make([]rune, 0, len(rs)*2)

	for i := len(rs); i > 0; i-- {
		if unicode.IsLetter(rune(rs[i-1])) {
			if i < len(rs) && unicode.IsUpper(rune(rs[i])) {
				if i > 1 && unicode.IsLower(rune(rs[i-1])) || i < len(rs)-2 && unicode.IsLower(rune(rs[i+1])) {
					buf = append(buf, '_')
				}
			}
			buf = append(buf, unicode.ToLower(rune(rs[i-1])))
		} else if unicode.IsDigit(rune(rs[i-1])) {
			if i == len(rs) || i == 1 || unicode.IsDigit(rune(rs[i])) {
				buf = append(buf, rune(rs[i-1]))
			} else {
				buf = append(buf, '_', rune(rs[i-1]))
			}
		} else {
			if i == len(rs) {
				continue
			}
			buf = append(buf, '_')
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
