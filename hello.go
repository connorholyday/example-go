package main

import (
    "fmt"
)

const helloPrefix = "Hello, "

func Hello(input string) string {
    suffix := "world"

    if input != "" {
        suffix = input
    }

    return helloPrefix + suffix
}

func main() {
    fmt.Println(Hello("world"))
}
