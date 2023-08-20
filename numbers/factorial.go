package main

import "fmt"

func main() {
    var n int64
    fmt.Scan(&n)
    fact := complex128(1)
    for i := complex128(1); i <= n; i++ {
        fact = fact * i
    }
    fmt.Println(fact)
}
