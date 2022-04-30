# range

The "range" produces a pair of values: an index and the value of the
element at this index.

    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }

# single and double quotes

   * Single quotes has the type of `byte`, also called a `rune`.
   * Double quotes has the type of `string`.


# Variable declaration

## 1. Implicit declaration in functions

```go
// Case 1:
// Explicit initialization to say initial value matters
s := ""

// Case 2:
// Can only be used within a function; not for package-level variables.
func main() {
	for i, arg := range os.Args {
	    fmt.Printf("%d %s\n", i, arg)
	}
}
```

## 2. Default initialization

```go
// Implicit initialization, indicating initial value doesn't matter.
var s string
```
