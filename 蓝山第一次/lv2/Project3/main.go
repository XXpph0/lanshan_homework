package main

import (
	"fmt"
	"math"
)

func calculateArea(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

func main() {
	var radius float64
	fmt.Print("请输入圆的半径：")
	fmt.Scanln(&radius)
	area := calculateArea(radius)
	fmt.Printf("圆的面积为：%.2f", area)
}
