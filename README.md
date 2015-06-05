# kace

    go get "github.com/codemodus/kace"

Package kace provides common case conversion functions which take into 
consideration common initialisms.

## Usage

```go
func Camel(s string, ucFirst bool) string
func Snake(s string) string
```

### Setup

```go
import (
	"fmt"

	"github.com/codemodus/kace"
)

func main() {
	s := "this is a test."

	fmt.Println(kace.Camel(s, false))
	fmt.Println(kace.Camel(s, true))

	fmt.Println(kace.Snake(s))

	// Output:
	// thisIsATest
	// ThisIsATest
	// this_is_a_test
}
```

## More Info

N/A

## Documentation

View the [GoDoc](http://godoc.org/github.com/codemodus/kace)

## Benchmarks

N/A
