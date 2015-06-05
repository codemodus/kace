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
func Camel(s string, ucFirst bool) string {
	rs := []rune(s)
	buf := make([]rune, 0, len(rs))

	for i := 0; i < len(rs); i++ {
		if unicode.IsLetter(rs[i]) {
			if i == 0 || !unicode.IsLetter(rs[i-1]) {
				tmpBuf := make([]rune, 0, ciMaxLen)
				for n := i; n < len(rs) && n-i < ciMaxLen; n++ {
					tmpBuf = append(tmpBuf, unicode.ToUpper(rs[n]))

					if n < len(rs)-1 && !unicode.IsLetter(rs[n+1]) && !unicode.IsDigit(rs[n+1]) {
						break
					}
				}
				if ci[string(tmpBuf)] {
					buf = append(buf, tmpBuf...)
					i += len(tmpBuf)
					continue
				}
			}

			if i == 0 && ucFirst || i > 0 && !unicode.IsLetter(rs[i-1]) {
				buf = append(buf, unicode.ToUpper(rs[i]))
			} else {
				buf = append(buf, rs[i])
			}
		}

		if unicode.IsDigit(rs[i]) {
			buf = append(buf, rs[i])
		}
	}
	return string(buf)
}

// Snake returns a snake cased string.
func Snake(s string) string {
	rs := []rune(s)
	rsus := []rune{'_'}
	buf := make([]rune, 0, len(rs)*2)

	for i := len(rs); i > 0; i-- {
		if unicode.IsLetter(rs[i-1]) {
			if i < len(rs) && unicode.IsUpper(rs[i]) {
				if i > 1 && unicode.IsLower(rs[i-1]) || i < len(rs)-2 && unicode.IsLower(rs[i+1]) {
					buf = append(rsus, buf...)
				}
			}
			buf = append([]rune{unicode.ToLower(rs[i-1])}, buf...)
		} else if unicode.IsDigit(rs[i-1]) {
			if i == len(rs) || i == 1 || unicode.IsDigit(rs[i]) {
				buf = append([]rune{rs[i-1]}, buf...)
			} else {
				buf = append([]rune{rs[i-1], '_'}, buf...)
			}
		} else {
			if i == len(rs) {
				continue
			}
			buf = append(rsus, buf...)
		}
	}
	return string(buf)
}
