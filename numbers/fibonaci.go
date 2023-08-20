package main

import "fmt"

func fibonaci(n int)int {
	if n == 0{
		return 0;
	}else if n == 1{
		return 1;
	}else{
		return fibonaci(n - 1) + fibonaci(n - 2)
	}
}
func main(){
	var n int64
	fmt.Scan(&n)
	fmt.Println(fibonaci(int(n)))
}