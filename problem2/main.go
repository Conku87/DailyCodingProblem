/*
Good morning! Here's your coding interview problem for today.

This problem was asked by Uber.

Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?
*/

package main

import "fmt"

func calculateProductListExceptSameIndex(slice []int) []int {
	workingSlice := makeStartingSlice(len(slice))
	leftProduct := 1
	for index, value := range slice {
		workingSlice[index] *= leftProduct
		leftProduct *= value
	}
	rightProduct := 1
	for i := len(slice) - 1; i >= 0; i-- {
		workingSlice[i] *= rightProduct
		rightProduct *= slice[i]
	}
	return workingSlice
}

func makeStartingSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = 1
	}
	return slice
}

func areSlicesEqual(slice1 []int, slice2 []int) bool {
	if len(slice1) == len(slice2) {
		for i := range slice1 {
			if slice1[i] != slice2[i] {
				return false
			}
		}
		return true
	}
	return false
}

func test(input []int, expectedOutput []int, testNumber int) {
	var result string
	if areSlicesEqual(expectedOutput, calculateProductListExceptSameIndex(input)) {
		result = fmt.Sprintf("test %d passed", testNumber)
	} else {
		result = fmt.Sprintf("test %d failed - expected %v, was %v", testNumber, expectedOutput, calculateProductListExceptSameIndex(input))
	}
	fmt.Println(result)
}

func main() {
	test([]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}, 1)
	test([]int{3, 2, 1}, []int{2, 3, 6}, 2)
	test([]int{5}, []int{1}, 3)
	test([]int{}, []int{}, 4)
}
