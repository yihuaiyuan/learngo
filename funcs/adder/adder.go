package main

import "fmt"

func adder() func(v int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}
func main() {
	num := 10
	a := adder()
	for i := 0; i < num; i++ {
		fmt.Printf("0+1+2..+%d = %d\n", i, a(i))
	}
}
