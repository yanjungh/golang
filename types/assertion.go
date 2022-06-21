package main

import (
    "fmt"
)

func main() {
    // From https://stackoverflow.com/questions/24492868/what-is-the-meaning-of-dot-parenthesis-syntax
    var i interface{}
    i = int(42)
    j := 42

    fmt.Printf("i == j: %v\n", i == j)      // i == j: true

    a, ok := i.(int)
    fmt.Printf("a: %d, ok: %v\n", a, ok)    // a: 42, ok: true

    b, ok := i.(string)
    fmt.Printf("b: %v, ok: %v\n", b, ok)    // b: , ok: false

    if _, ok := i.(int); ok {
        fmt.Printf("Idiom names: TryAssert and 'Comma ok'\n")
    }
}