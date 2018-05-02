package iteration

import (
    "testing"
    "fmt"
    "strings"
)

func TestRepeat(t *testing.T) {
    char := "a"
    amount := 10

    repeated := Repeat(char, amount)
    got := strings.Count(repeated, char)

    if got != amount {
        t.Errorf("Got '%d' want '%d'", got, amount)
    }
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 10);
    }
}

func ExampleRepeat() {
    got := Repeat("a", 10)
    fmt.Println(got)
    // Output: aaaaaaaaaa
}
