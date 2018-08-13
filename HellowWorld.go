package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func add2(x, y float32) float32 {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("hello world !", math.Pi, math.Nextafter(1, 2), add(5, 10))
	fmt.Println(split(50))
}

// http://go-tour-kr.appspot.com/#11
