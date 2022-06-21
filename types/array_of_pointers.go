package main

import (
    "fmt"
//     "testing"
)

type City struct {
    name    string
    state   string
}

func printBySlice(c []City) {
    fmt.Printf("Func: printBySlice\n")
    for _, city := range(c) {
        fmt.Printf("City: %s, State: %s\n", city.name, city.state)
    }
}

func printByPointer(c []*City) {
    fmt.Printf("Func: printByPointer\n")
    for _, city := range(c) {
        fmt.Printf("City: %s, State: %s\n", city.name, city.state)
    }
}

func main() {
    c1 := City{name: "New York", state: "New York"}
    c2 := City{name: "Mountain View", state: "California"}

    c := [...]City{c1, c2}
    printBySlice(c[:])

    fmt.Println()
    cities := []*City{{name: "New York", state: "New York"}, {name: "Boston", state: "Massachusetts"}}
    printByPointer(cities)

    fmt.Printf("Type of cities: %T\n", cities)      // []*main.City

    fmt.Println()
    c3 := &City{name: "Anchorage", state: "Alaska"}
    fmt.Printf("Value of c3: %v, type of c3: %T\n", c3, c3)
    // Value of c3: &{Anchorage Alaska}, type of c3: *main.City
}
