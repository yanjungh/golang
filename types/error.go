package main

import (
    "errors"
    "fmt"
    "strconv"
)

var NaN = errors.New("Not a number")

func convStr(s string) (int, error) {
    intVar, err := strconv.Atoi(s)
    if err != nil {
        return 0, NaN
    }
    return intVar, nil
}

func ConvStr(s string) {
    i, err := convStr(s)
    if err != nil {
        switch {
        case errors.Is(err, NaN):
            fmt.Printf("Input str: %s is not a number\n", s)
        default:
            fmt.Printf("Impossible error: %s\n", err)
        }
    } else {
        fmt.Printf("s: %s, i: %d\n", s, i)
    }
}

func main() {
    s := "29"
    ConvStr(s)

    t := "XY"
    ConvStr(t)
}



