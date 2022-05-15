package main

import (
    "fmt"
)

type price struct {
    unit interface{}
    total float32
}

func main() {
    m := make(map[string]price)         // key: string, value: price
    m["apple"] = price{3, 3.99}         // 3 apples for $3.99
    fmt.Printf("Apple: %v\n", m["apple"])

    m["cherry"] = price{"1LB", 5.99}     // $5.99 per pound
    fmt.Printf("Apple: %v\n", m["cherry"])
}
