/* Pointers in go, source code from:
    https://go.dev/tour/moretypes/1
 */
package main

import (
    "fmt"
)

func main() {
    i := 42
	describe(i)     // Output: (42, int)

    p := &i         // The & operator generates a pointer to its operand
	describe(*p)    // Same as describe(i); read value of i through a pointer
	describe(p)     // Output: (0x8812118, *int)

	*p = 21         // Set value of i through a pointer
	describe(*p)
	describe(p)     // No change, since we didn't change the pointer.

	j := 99
	p = &j
	describe(*p)    // Output: (99, int)
	describe(p)     // Output: (0x9012128, *int)
}

/* An empty interface doesn't specify any method, and we use it to handle values of unknown types.
   See explanations on:
   https://go.dev/tour/methods/14
*/
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
