/*
Good morning! Here's your coding interview problem for today.

This problem was recently asked by Google.

Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?
*/

package main

import "fmt"

func sliceContainsTwoValuesThatSumToK(slice []int, k int) bool {
	var differences []int
	for _, value := range slice {
		if sliceContainsValue(differences, value) {
			return true
		} else {
			differences = append(differences, k-value)
		}
	}
	return false
}

func sliceContainsValue(slice []int, k int) bool {
	for _, value := range slice {
		if k == value {
			return true
		}
	}
	return false
}

func test(slice []int, k int, testID string, expectedResult bool) {
	if expectedResult == sliceContainsTwoValuesThatSumToK(slice, k) {
		fmt.Println("test " + testID + " passed")
	} else {
		fmt.Println("test " + testID + " failed")
	}
}

func main() {
	test([]int{10, 15, 3, 7}, 17, "1", true)
	test([]int{1, 2, 3}, 3, "2", true)
	test([]int{3}, 3, "3", false)
	test([]int{1, 1, 1}, 3, "4", false)
	test([]int{0, 0}, 0, "5", true)
	test([]int{}, 5, "6", false)
	test([]int{-3, 3, -3}, -6, "7", true)
	test([]int{-3, 6}, 3, "8", true)
	test([]int{3}, 6, "9", false)
	test([]int{3, 0}, 3, "10", true)
}
