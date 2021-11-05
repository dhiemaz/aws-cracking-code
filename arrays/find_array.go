package main

import "fmt"

func main() {
	var a []int = []int{5, 7, 8, 9, 4, 3, 2, 1}
	// sorted_array := sort(a)
	// fmt.Println(findMissingValue(sorted_array))

	n := len(a) + 1
	fmt.Println("n :", n)

	actual_sum := (n * (n + 1)) / 2

	fmt.Println("actual sum = ", actual_sum)
	fmt.Println(actual_sum - sum(a))
}

// func sort(a []int) []int {
// 	for i := 0; i < len(a); i++ {
// 		for j := i + 1; j < len(a); j++ {
// 			if a[i] > a[j] {
// 				a[i], a[j] = a[j], a[i]
// 			}
// 		}

// 		fmt.Println(a)
// 	}

// 	return a
// }

// findMissingValue in array
func sum(a []int) int {
	var result int
	for _, v := range a {
		result += v
	}

	return result
}
