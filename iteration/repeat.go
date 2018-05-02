package iteration

func Repeat(x string, amount int) string {
    var repeated string
    for i := 0; i < amount; i++ {
        repeated += x
    }
    return repeated
}
