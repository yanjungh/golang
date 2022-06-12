package main

import (
    "fmt"
)

func arrayType1() {
    var a [3]int = [3]int{2, 3, 5}

    // the ellipsis "..." means length is determined by the number of initializers
    b := [...]int{2, 3, 5, 7}

    fmt.Printf("a: %T\n", a)
    fmt.Printf("b: %T\n", b)

    // b = a
    // cannot use a (variable of type [3]int) as type [4]int in assignment
}

func arrayType2() {
    type Month int
    const (
        PADDING Month = iota
        JAN
        FEB
        MAR
        APR
        MAY
        JUN
    )

    // Note indices can appear in *ANY* order
    month := [...]string{JAN: "January", MAR: "March", MAY: "May", FEB: "February", APR: "April", JUN: "June"}
    fmt.Printf("Month: %v, month: %v\n", MAY, month[MAY])
}

func arrayType3() {
    // Indices not specified have value of "0"
    a := [...]int{5: 100, 7: 200, 9: 300}
    fmt.Printf("a: %v\n", a)
    // a: [0 0 0 0 0 100 0 200 0 300]
}

func main() {
    arrayType1()
    arrayType2()
    arrayType3()
}