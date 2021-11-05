package main

func main() {
	var a = []int{5, 7, 1, 2, 8, 4, 3}
	findSumOfTwo(a, 19)
}

func findSumOfTwo(a []int, val int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == val {
				println(a[i], a[j], "true")
			}
		}
	}

	println("false")
}
