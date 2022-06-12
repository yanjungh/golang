package main

import (
    "fmt"
)

/* Because a slice contains a pointer to an element of an array,
   passing a slice to a function means passing a pointer.
   This allows us to swap values in place.
 */
func reverse(s []int) {
    n := len(s)
    for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1  {
        s[i], s[j] = s[j], s[i]
    }
}

func main() {
    a := [...]int{1, 2, 3, 4, 5}
    reverse(a[:])   // Reverse the whole array
    fmt.Printf("a: %v\n", a)
}