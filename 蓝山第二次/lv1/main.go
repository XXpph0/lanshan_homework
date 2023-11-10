package main

import "fmt"

func Calculator(num1 int, num2 int, CMD func(int, int) int) int {
	return CMD(num1, num2)
}
func Add(a, b int) int {
	return a + b
}
func Minus(a, b int) int {
	return a - b
}
func main() {
	result := Calculator(5, 10, Add)
	fmt.Println("5 + 10 =", result)

	result = Calculator(10, 5, Minus)
	fmt.Println("10 - 5 =", result)
}
