package main

import (
    "fmt"
)

type fruit struct {
    name string
    price float32
}

func main() {
    apple := fruit{name: "apple", price: 2.0}

    // Output:
    // Apple: {apple 2}, name: apple, price: 2.000000
    fmt.Printf("Apple: %v, name: %s, price: %f\n", apple, apple.name, apple.price)

    ap := &apple
    ap.name = "pineapple"   // equivalent of "(*ap).name"
    fmt.Printf("Apple: %v\n", *ap)  // Output: Apple: {pineapple 2}
}
