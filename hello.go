package main

import (
    "fmt"
)

const helloPrefix = "Hello, "

func Hello(input string) string {
    return helloPrefix + input
}

func main() {
    fmt.Println(Hello("world"))
}
