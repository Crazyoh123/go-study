package main

import "fmt"

func primeFactorization(n int) {
    fmt.Printf("Prime Factorization of %d: ", n)
    
    for i := 2; i <= n; i++ {
        for n%i == 0 {
            fmt.Printf("%d ", i)
            n /= i
        }
    }
    
    fmt.Println()
}

func main() {
    var n int
    fmt.Print("Enter a positive integer: ")
    fmt.Scan(&n)
    
    if n <= 0 {
        fmt.Println("Please enter a positive integer.")
        return
    }
    
    primeFactorization(n)
}
