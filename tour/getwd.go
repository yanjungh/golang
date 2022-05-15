package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    cwd := cwd()
    fmt.Println("cwd is ", cwd)
}

func cwd() string {
    var err error
    var cwd string
    cwd, err = os.Getwd()
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
        cwd = ""
    }

    return cwd
}

