package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)


func scanFromStdin() string {
    fmt.Println("Please enter your full name, then press Ctrl+D: ")
	input := bufio.NewScanner(os.Stdin)
	names := make([]string, 0)
	for input.Scan() {
	    names = append(names, input.Text())
	}
	return strings.Join(names, " ")
}

func scanTwice() string {
    var first, last string

    fmt.Println("Your first name: ")
    fmt.Scanln(&first)          // pass by reference

    fmt.Println("Your last name: ")
    fmt.Scanln(&last)

    return fmt.Sprintf("%s %s", first, last)
}

func main() {
    name1 := scanFromStdin()
    fmt.Printf("Welcome to golang, %s!\n", name1)

    // Cannot re-use "name1"
    name2 := scanTwice()
    fmt.Printf("Welcome to golang, %s!\n", name2)
}