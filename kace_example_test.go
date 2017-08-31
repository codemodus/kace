package kace_test

import (
	"fmt"

	"github.com/codemodus/kace"
)

func Example() {
	s := "this is a test."

	fmt.Println(kace.Camel(s))
	fmt.Println(kace.Pascal(s))

	fmt.Println(kace.Snake(s))
	fmt.Println(kace.SnakeUpper(s))

	fmt.Println(kace.Kebab(s))
	fmt.Println(kace.KebabUpper(s))

	// Output:
	// thisIsATest
	// ThisIsATest
	// this_is_a_test
	// THIS_IS_A_TEST
	// this-is-a-test
	// THIS-IS-A-TEST
}
