package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func cal(x int, y int) (plus int, minus int, multiple int, division float64) {
	plus = x + y
	minus = x - y
	multiple = x * y
	division = float64(x) / float64(y)
	return
}

// inner function
func noNameFunc() {
	f := func(x int) {
		fmt.Println(x)
	}
	f(1)

	func(x int) {
		fmt.Println(x)
	}(5)
}

// closure
func calArea(pi float64) func(radius int) float64 {
	return func(radius int) float64 {
		return pi * float64(radius*radius)
	}
}

// closure
func calFib(n int) func() int {
	c1, c2 := 0, 1
	return func() int {
		for index := 0; index < n; index++ {
			c1, c2 = c2, c1+c2
		}
		return c1
	}
}

// 可変長引数
func calSum(args ...int) (sum int) {
	for _, arg := range args {
		sum += arg
	}
	return
}

func main() {
	fmt.Println(cal(2, 3))
	noNameFunc()
	c := calArea(3.14)
	fmt.Println(c(2))

	f := calFib(4)
	fmt.Println(f())

	fmt.Println(calSum(2, 4, 6, 8, 10))
	s := []int{1, 3, 5, 7, 9}
	fmt.Println(calSum(s...))
}
