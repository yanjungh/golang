/* Go _method_ is a special type of function.
 */

package main

import (
    "fmt"
)

type rectangle struct {
    length, width float32
}

/* Method #1: receiver is a struct.
 * The receiver "(r rectangle)" appears between the "func" keyword and the method name "area()"
 */
func (r rectangle) area() float32 {
    return r.length * r.width
}

// Method #2, receiver is a pointer
func (r *rectangle) areaP() float32 {
    return r.length * r.width
}

// A regular function
func area2(r rectangle) float32 {
    return r.length * r.width
}

func main() {
    r := rectangle{3.0, 5.0}
    fmt.Printf("Area of r %v is %f\n", r, r.area())

    p := &r
    fmt.Printf("Area of r %v is %f\n", p, p.areaP())
    fmt.Printf("Area of r %v is %f\n", r, area2(r))
}