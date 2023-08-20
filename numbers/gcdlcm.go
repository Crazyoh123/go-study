package main

import (
    "fmt"
    "math"
)

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}

func lcm(a, b int) int {
    return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}

func main() {
    var n, n1 int
    fmt.Print("Enter two integers: ")
    fmt.Scan(&n, &n1)

    fmt.Println("GCD:", gcd(n, n1))
    fmt.Println("LCM:", lcm(n, n1))
}
