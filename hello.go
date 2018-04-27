package main

import (
    "fmt"
)

const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(input string, language string) string {
    suffix := "world"

    if input != "" {
        suffix = input
    }
    
    if language == "Spanish" {
        return spanishHelloPrefix + suffix
    }

    return helloPrefix + suffix
}

func main() {
    fmt.Println(Hello("world", "English"))
}
