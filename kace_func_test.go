package kace_test

import (
	"testing"

	"github.com/codemodus/kace"
)

func TestFuncPascal(t *testing.T) {
	var data = []struct {
		i string
		o string
	}{
		{"This is a test sql", "ThisIsATestSQL"},
		{"this is a test sql", "ThisIsATestSQL"},
	}

	for _, v := range data {
		want := v.o
		got := kace.Pascal(v.i)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestFuncCamel(t *testing.T) {
	var data = []struct {
		i string
		o string
	}{
		{"this is a test sql", "thisIsATestSQL"},
		{"this_is_a_test sql", "thisIsATestSQL"},
	}

	for _, v := range data {
		want := v.o
		got := kace.Camel(v.i)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestFuncSnake(t *testing.T) {
	var data = []struct {
		i string
		o string
	}{
		{"thisIsATestSQL", "this_is_a_test_sql"},
		{"ThisIsATestSQL", "this_is_a_test_sql"},
	}

	for _, v := range data {
		want := v.o
		got := kace.Snake(v.i)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestFuncSnakeUpper(t *testing.T) {
	var data = []struct {
		i string
		o string
	}{
		{"thisIsATestSQL", "THIS_IS_A_TEST_SQL"},
		{"ThisIsATestSQL", "THIS_IS_A_TEST_SQL"},
	}

	for _, v := range data {
		want := v.o
		got := kace.SnakeUpper(v.i)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestFuncKebab(t *testing.T) {
	var data = []struct {
		i string
		o string
	}{
		{"thisIsATestSQL", "this-is-a-test-sql"},
		{"ThisIsATestSQL", "this-is-a-test-sql"},
	}

	for _, v := range data {
		want := v.o
		got := kace.Kebab(v.i)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestFuncKebabUpper(t *testing.T) {
	var data = []struct {
		i string
		o string
	}{
		{"thisIsATestSQL", "THIS-IS-A-TEST-SQL"},
		{"ThisIsATestSQL", "THIS-IS-A-TEST-SQL"},
	}

	for _, v := range data {
		want := v.o
		got := kace.KebabUpper(v.i)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
