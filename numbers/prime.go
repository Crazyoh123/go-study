package main

import (
    "fmt"
    "math"
)

func main() {
    var n int64
    var k bool = true
    fmt.Scan(&n)
    
    for i := int64(2); i <= int64(math.Sqrt(float64(n))); i++ {
        if n%i == 0 {
            k = false
            break
        }
    }
    
    if k && n > 1 {
        fmt.Println("IT IS PRIME NUMBER")
    } else {
        fmt.Println("IT IS NOT PRIME NUMBER")
    }
}
