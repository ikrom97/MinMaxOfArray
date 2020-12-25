package main

import "fmt"

func MinOfArray(array []int64) {
	min := array[0]
	for _, v := range array {
		if v < min {
			min = v
		}
	}
	fmt.Println("Минимальное значение массива =", min)
}
func MaxOfArray(array []int64) {
	max := array[0]
	for _, v := range array {
		if v > max {
			max = v
		}
	}
	fmt.Println("Максимальное значение массива =", max)
}
func main() {
	array := []int64{4, 2, 54, 11, 10, 45}
	MinOfArray(array)
	MaxOfArray(array)
}
