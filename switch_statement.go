package main

import (
	"fmt"
)

func noFallthrough(score int) {
	var grade string
	switch {
	case score > 90:
		grade = "A"
	case score > 70:
		grade = "B"
	case score > 50:
		grade = "C"
	default:
		grade = "F"
	}
	fmt.Println("noFallthrough: ", grade)
}

func yesFallthrough(score int) {
	var grade string
	switch {
	case score > 90:
		grade = "A"
		fallthrough
	case score > 70:
		grade = "B"
		fallthrough
	case score > 50:
		grade = "C"
		fallthrough
	default:
		grade = "F"
	}
	fmt.Println("yesFallthrough: ", grade)
}

func main() {
	fmt.Println("Grade: ")

	score := 100
	yesFallthrough(score)
	noFallthrough(score)
}
