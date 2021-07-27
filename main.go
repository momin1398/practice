package main

import "fmt"

func main() {
	var c int
	fmt.Scan(&c)
	if c%2 == 0 {
		fmt.Println("even number")
	}

}
