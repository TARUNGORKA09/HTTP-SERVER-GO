package main

import "fmt"

func main() {
	//var sum int = 1
	sum := 1
	for sum < 100 {
		sum += sum
		fmt.Println(sum)
	}
}
