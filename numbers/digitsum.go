package main

import "fmt"

func main(){
	var n int64
	fmt.Scan(&n)
	var sum int64
	for n != 0{
		sum += n % 10;
		n = n / 10;
	}
	fmt.Println("sum is ",sum)
}