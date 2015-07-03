# kace

    go get "github.com/codemodus/kace"

Package kace provides common case conversion functions which take into 
consideration common initialisms.

## Usage

```go
func Camel(s string, ucFirst bool) string
func Kebab(s string) string
func KebabUpper(s string) string
func Snake(s string) string
func SnakeUpper(s string) string
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
```

## More Info

### TODO

#### Reduce allocations in camel func.

I'd like to make use of a tree of some sort to look up initialisms so that 
allocations can be avoided.  This will need to be created from the intialisms 
map during init and be tested against using an unexported function.

## Documentation

View the [GoDoc](http://godoc.org/github.com/codemodus/kace)

## Benchmarks

    benchmark                 iter       time/iter   bytes alloc        allocs
    ---------                 ----       ---------   -----------        ------
    BenchmarkCamel4        1000000   1285.00 ns/op      144 B/op   7 allocs/op
    BenchmarkSnake4        2000000    696.00 ns/op      128 B/op   2 allocs/op
    BenchmarkSnakeUpper4   2000000    679.00 ns/op      128 B/op   2 allocs/op
    BenchmarkKebab4        2000000    691.00 ns/op      128 B/op   2 allocs/op
    BenchmarkKebabUpper4   2000000    677.00 ns/op      128 B/op   2 allocs/op
