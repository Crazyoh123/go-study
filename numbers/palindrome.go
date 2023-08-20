package main

import "fmt"

func main(){
	var rev  int64
	var n int64
	fmt.Scan(&n)
	k :=  n 
	for n!=0{
		rev = (rev * 10) + n % 10;
		n = n /10;
	}
	if rev == k {
		fmt.Println("IT IS PALINDROME!")
	}else{
		fmt.Println("IT IS NOT PALINDROME!")
	}
}