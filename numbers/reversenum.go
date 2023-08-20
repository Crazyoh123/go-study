package main

import "fmt"

func main() {
    var n, rev int64
    fmt.Print("Enter an integer: ")
    fmt.Scan(&n)
    
    for n != 0 {
        rev = (rev * 10) + n % 10
        n = n / 10
    }
    
    fmt.Println("Reversed:", rev)
}
