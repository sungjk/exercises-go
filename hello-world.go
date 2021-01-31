package main

import "fmt"

func divide1(dividend, divisor int) (int, int) {
	var quotient = (int)(dividend / divisor)
	var remainder = dividend % divisor
	return quotient, remainder
}

func divide2(dividend, divisor int) (quotient, remainder int) {
	quotient = (int)(dividend / divisor)
	remainder = dividend % divisor
	return
}

func main() {
	var quotient, remainder int
	quotient, remainder = divide1(10, 3)
	fmt.Println("1: ", quotient, remainder)

	quotient, remainder = divide2(10, 3)
	fmt.Println("2: ", quotient, remainder)
}
