package main

import "fmt"

func main() {
	getSliceDetail()
	loopZeroLengthSlice()
	loopSlice()
}

func getArray() [2]int {
	var a [2]int
	a[0] = 1
	a[1] = 2
	return a
}

func getSlice() []int {
	var a []int = []int{100, 200}
	a = append(a, 300)
	return a
}

func getSliceDetail() {
	a := make([]int, 3, 5)
	a = append(a, 0, 0, 0)
	printSliceDetail(a)

	a = append(a, 100, 200, 300, 400, 500)
	printSliceDetail(a)
}

func printSliceDetail(array []int) {
	fmt.Printf("slice: %v, len: %d, capacity: %d\n", array, len(array), cap(array))
}

func loopZeroLengthSlice() {
	s := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		s = append(s, i)
		printSliceDetail(s)
	}
}

func loopSlice() {
	s := make([]int, 5)
	for i := 0; i < 5; i++ {
		s = append(s, i)
		printSliceDetail(s)
	}
}
