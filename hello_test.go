package main

import (
    "testing"
)

func TestHello(t *testing.T) {

    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got '%s' want '%s'", got, want)
        }
    }

    t.Run("Saying hello to people", func(t *testing.T) {
        input := "YOU"
        got := Hello(input)
        want := "Hello, " + input
        assertCorrectMessage(t, got, want)
    })

    t.Run("Saying hello world when an empty string is supplied", func(t *testing.T) {
        input := ""
        got := Hello(input)
        want := "Hello, world"
        assertCorrectMessage(t, got, want)
    })

}
