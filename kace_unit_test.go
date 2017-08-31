package kace

import (
	"testing"
)

func TestUnitPascal(t *testing.T) {
	var tests = []struct {
		i string
		o string
	}{
		{"This is a test", "ThisIsATest"},
		{"this is a test", "ThisIsATest"},
		{"this is 4 test", "ThisIs4Test"},
		{"5this is a test", "5ThisIsATest"},
		{"this_is_a_test", "ThisIsATest"},
		{"This is a test.", "ThisIsATest"},
		{"This.is.a.Test", "ThisIsATest"},
		{"andThisToo", "AndThisToo"},
		{"AndThisToo", "AndThisToo"},
		{"this http conn", "ThisHTTPConn"},
		{"this_https_conn", "ThisHTTPSConn"},
		{"this_http_scan", "ThisHTTPScan"},
		{"willid mess it up", "WillidMessItUp"},
		{"willid_mess_it_up", "WillidMessItUp"},
		{"http_first_upper", "HTTPFirstUpper"},
		{"ahttp_upper", "AhttpUpper"},
	}

	for _, v := range tests {
		want := v.o
		got := Pascal(v.i)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestUnitCamel(t *testing.T) {
	var tests = []struct {
		i string
		o string
	}{
		{"this is a test", "thisIsATest"},
		{"this_is_a_test", "thisIsATest"},
		{"this is a test.", "thisIsATest"},
		{"this.is.a.Test", "thisIsATest"},
		{"AndThisToo", "andThisToo"},
		{"andThisToo", "andThisToo"},
		{"this http conn", "thisHTTPConn"},
		{"this_https_conn", "thisHTTPSConn"},
		{"this_http_scan", "thisHTTPScan"},
		{"willid mess it up", "willidMessItUp"},
		{"willid_mess_it_up", "willidMessItUp"},
		{"http_first_lower", "httpFirstLower"},
		{"ahttp_lower", "ahttpLower"},
	}

	for _, v := range tests {
		want := v.o
		got := Camel(v.i)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestUnitSnake(t *testing.T) {
	var tests = []struct {
		i string
		o string
	}{
		{"thisIsATest", "this_is_a_test"},
		{"ThisIsATest", "this_is_a_test"},
		{"ThisIsATest3", "this_is_a_test3"},
		{"ThisIs44Test", "this_is44_test"},
		{"5ThisIsATest", "5this_is_a_test"},
		{"this is a test", "this_is_a_test"},
		{"this_is_a_test", "this_is_a_test"},
		{"This is a test.", "this_is_a_test"},
		{"This.is.a.Test", "this_is_a_test"},
		{"thisHTTPSConn", "this_https_conn"},
		{"thisHTTPScan", "this_http_scan"},
		{"ThisHTTPSConn", "this_https_conn"},
		{"ThisHTTPScan", "this_http_scan"},
		{"willidMessItUp", "willid_mess_it_up"},
		{"WillidMessItUp", "willid_mess_it_up"},
	}

	for _, v := range tests {
		want := v.o
		got := Snake(v.i)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestUnitSnakeUpper(t *testing.T) {
	var tests = []struct {
		i string
		o string
	}{
		{"thisIsATest", "THIS_IS_A_TEST"},
		{"ThisIsATest", "THIS_IS_A_TEST"},
		{"ThisIsATest3", "THIS_IS_A_TEST3"},
		{"ThisIs44Test", "THIS_IS44_TEST"},
		{"5ThisIsATest", "5THIS_IS_A_TEST"},
		{"this is a test", "THIS_IS_A_TEST"},
		{"this_is_a_test", "THIS_IS_A_TEST"},
		{"This is a test.", "THIS_IS_A_TEST"},
		{"This.is.a.Test", "THIS_IS_A_TEST"},
		{"thisHTTPSConn", "THIS_HTTPS_CONN"},
		{"ThisHTTPSConn", "THIS_HTTPS_CONN"},
		{"willidMessItUp", "WILLID_MESS_IT_UP"},
		{"WillidMessItUp", "WILLID_MESS_IT_UP"},
	}

	for _, v := range tests {
		want := v.o
		got := SnakeUpper(v.i)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestUnitKebab(t *testing.T) {
	var tests = []struct {
		i string
		o string
	}{
		{"thisIsATest", "this-is-a-test"},
		{"ThisIsATest", "this-is-a-test"},
		{"ThisIsATest3", "this-is-a-test3"},
		{"ThisIs44Test", "this-is44-test"},
		{"5ThisIsATest", "5this-is-a-test"},
		{"this is a test", "this-is-a-test"},
		{"this_is_a_test", "this-is-a-test"},
		{"This is a test.", "this-is-a-test"},
		{"This.is.a.Test", "this-is-a-test"},
		{"thisHTTPSConn", "this-https-conn"},
		{"ThisHTTPSConn", "this-https-conn"},
		{"willidMessItUp", "willid-mess-it-up"},
		{"WillidMessItUp", "willid-mess-it-up"},
	}

	for _, v := range tests {
		want := v.o
		got := Kebab(v.i)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestUnitKebabUpper(t *testing.T) {
	var tests = []struct {
		i string
		o string
	}{
		{"thisIsATest", "THIS-IS-A-TEST"},
		{"ThisIsATest", "THIS-IS-A-TEST"},
		{"ThisIsATest3", "THIS-IS-A-TEST3"},
		{"ThisIs44Test", "THIS-IS44-TEST"},
		{"5ThisIsATest", "5THIS-IS-A-TEST"},
		{"this is a test", "THIS-IS-A-TEST"},
		{"this_is_a_test", "THIS-IS-A-TEST"},
		{"This is a test.", "THIS-IS-A-TEST"},
		{"This.is.a.Test", "THIS-IS-A-TEST"},
		{"thisHTTPSConn", "THIS-HTTPS-CONN"},
		{"ThisHTTPSConn", "THIS-HTTPS-CONN"},
		{"willidMessItUp", "WILLID-MESS-IT-UP"},
		{"WillidMessItUp", "WILLID-MESS-IT-UP"},
	}

	for _, v := range tests {
		want := v.o
		got := KebabUpper(v.i)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
