package main

import "fmt"

func main() {
	var head1 []int = []int{4, 8, 15, 19}
	var head2 []int = []int{7, 9, 10, 16}

	fmt.Println(merged_sort(&head1, &head2))
	fmt.Println(merge_sort(head1, head2))
}

func merge_sort(head1, head2 []int) []int {
	head1 = append(head1, head2...)
	return sort(head1)
}

func sort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	return a
}

func merged_sort(head1, head2 *[]int) *[]int {
	var head []int = []int{}
	var i, j int = 0, 0
	for i < len(*head1) && j < len(*head2) {
		if (*head1)[i] < (*head2)[j] {
			head = append(head, (*head1)[i])
			i++
		} else {
			head = append(head, (*head2)[j])
			j++
		}
	}

	for i < len(*head1) {
		head = append(head, (*head1)[i])
		i++
	}

	for j < len(*head2) {
		head = append(head, (*head2)[j])
		j++
	}

	return &head
}
