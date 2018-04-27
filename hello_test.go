package main

import (
    "testing"
)

func TestHello(t *testing.T) {
    input := "YOU"
    got := Hello(input)
    want := "Hello, " + input

    if got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}
