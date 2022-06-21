package main

import (
    "fmt"
)

// Reference: https://go.dev/blog/slices

/* Best practice: use a slice instead of an array or a pointer.

   A slice has empty size, e.g. []int instead [5]int.

   Because a slice contains a pointer to an element of an array,
   passing a slice to a function means passing a pointer.
   This allows us to swap values in place.
 */
func reverse(s []int) {
    n := len(s)
    for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1  {
        s[i], s[j] = s[j], s[i]
    }
}

// Passing a pointer.
func reverseByPointerToArray(a *[5]int) {
    n := len(a)
    for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1  {
        a[i], a[j] = a[j], a[i]
    }
}

// Passing a pointer to a slice
func reverseByPointerToSlice(slicePtr *[]int) {
    a := *slicePtr
    n := len(a)
    for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1  {
        a[i], a[j] = a[j], a[i]
    }
}

func main() {
    a := [...]int{1, 2, 3, 4, 5}

    fmt.Printf("Original array, a: %v\n", a)

    // We pass a slice, instead an array that has type [5]int
    reverse(a[:])
    fmt.Printf("Reverse once, a: %v\n", a)

    reverseByPointerToArray(&a)
    fmt.Printf("Reverse twice, a: %v\n", a)

    /* &(a[:]) gives out below error:
        invalid operation: cannot take address of a[:] (value of type []int)
     */
    s := a[:]
    reverseByPointerToSlice(&s)
    fmt.Printf("Reverse thrice, a: %v\n", a)
}
