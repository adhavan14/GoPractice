package main

import (
	"fmt"
	"math"
)

func calculate() {
	var num1 float64
	var num2 float64
	fmt.Print("Enter num1 = ")
	fmt.Scan(&num1)
	fmt.Print("Enter num2 = ")
	fmt.Scan(&num2)
	var sum = num1 + num2
	var sub float64= num1 - num2
	var pow = math.Pow(num1, num2)

	fmt.Print(sum, sub, pow)
}