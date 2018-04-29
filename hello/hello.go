package main

import (
    "fmt"
)

const englishPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(input string, language string) string {
    suffix := "world"

    if input != "" {
        suffix = input
    }


    return greetingPrefix(language) + suffix
}

func greetingPrefix(language string) (prefix string) {
    switch language {
        case "Spanish":
            prefix = spanishHelloPrefix
        case "French":
            prefix = frenchHelloPrefix
        default:
            prefix = englishPrefix
    }
    return
}

func main() {
    fmt.Println(Hello("world", "English"))
}
