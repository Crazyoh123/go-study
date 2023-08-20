package main

import "fmt"

func power(a, b int) int {
    var res int64 = 1 // Initialize result to 1
    for i := 1; i <= b; i++ {
        res *= int64(a) // Multiply by 'a' in each iteration
    }
    return int(res)
}

func main() {
    var a, b int64
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Println(power(int(a), int(b)))
}
