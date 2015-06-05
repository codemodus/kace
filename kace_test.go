package kace_test

import (
	"fmt"
	"testing"

	"github.com/codemodus/kace"
)

func Example() {
	s := "this is a test."

	fmt.Println(kace.Camel(s, false))
	fmt.Println(kace.Camel(s, true))

	fmt.Println(kace.Snake(s))

	// Output:
	// thisIsATest
	// ThisIsATest
	// this_is_a_test
}

func TestCamel(t *testing.T) {
	var tests = []struct {
		i   string
		ucf bool
		o   string
	}{
		{"this is a test", false, "thisIsATest"},
		{"this is a test", true, "ThisIsATest"},
		{"this is a test3", true, "ThisIsATest3"},
		{"this is 4 test", true, "ThisIs4Test"},
		{"5this is a test", true, "5ThisIsATest"},
		{"this_is_a_test", false, "thisIsATest"},
		{"this_is_a_test", true, "ThisIsATest"},
		{"this is a test.", false, "thisIsATest"},
		{"This is a test.", true, "ThisIsATest"},
		{"this.is.a.Test", false, "thisIsATest"},
		{"This.is.a.Test", true, "ThisIsATest"},
		{"AndThisToo", false, "AndThisToo"},
		{"andThisToo", false, "andThisToo"},
		{"andThisToo", true, "AndThisToo"},
		{"AndThisToo", true, "AndThisToo"},
		{"this http conn", false, "thisHTTPConn"},
		{"this http conn", true, "ThisHTTPConn"},
		{"this_https_conn", false, "thisHTTPSConn"},
		{"this_https_conn", true, "ThisHTTPSConn"},
		{"willid mess it up", false, "willidMessItUp"},
		{"willid mess it up", true, "WillidMessItUp"},
		{"willid_mess_it_up", false, "willidMessItUp"},
		{"willid_mess_it_up", true, "WillidMessItUp"},
	}

	for _, v := range tests {
		want := v.o
		got := kace.Camel(v.i, v.ucf)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestSnake(t *testing.T) {
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
		{"ThisHTTPSConn", "this_https_conn"},
		{"willidMessItUp", "willid_mess_it_up"},
		{"WillidMessItUp", "willid_mess_it_up"},
	}

	for _, v := range tests {
		want := v.o
		got := kace.Snake(v.i)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func BenchmarkCamel4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = kace.Camel("this_is_a_test", true)
	}
}

func BenchmarkSnake4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = kace.Snake("ThisIsATest")
	}
}
