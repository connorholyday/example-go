package sum

import (
    "testing"
    "reflect"
    "fmt"
)

func TestSum(t *testing.T) {

    t.Run("collection of varying numbers", func(t *testing.T) {
        numbers := []int{1, 2, 3}

        got := Sum(numbers)
        want := 6

        if got != want {
            t.Errorf("got '%d' but want '%d', given %v", got, want, numbers)
        }
    })
}

func checkSum(t *testing.T, got, want []int) {
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got '%d' but want '%d'", got, want)
    }
}

func TestSumAll(t *testing.T) {
    t.Run("sum of multiple collections of varying numbers", func(t *testing.T) {
        first := []int{1, 2, 3}
        second := []int{1, 2, 3, 4}

        got := SumAll(first, second)
        want := []int{6, 10}

        checkSum(t, got, want)
    })
}

func ExampleSumAll() {
    got := SumAll([]int{1, 2}, []int{3, 4})
    fmt.Println(got)
    // Output: [3 7]
}

func TestSumAllTails(t *testing.T) {

    t.Run("sum of slice tails", func(t *testing.T) {
        got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
        want := []int{5, 11}

        checkSum(t, got, want)
    })

    t.Run("safely sum empty slices", func(t *testing.T) {
        got := SumAllTails([]int{}, []int{4, 5, 6})
        want := []int{0, 11}

        checkSum(t, got, want)
    })
}
