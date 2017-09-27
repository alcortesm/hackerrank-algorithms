package main

import "fmt"

const size = 5

func main() {
	inputs := readInts(size)
	sums := allPossibleSums(inputs)
	min, max := minAndMax(sums)
	fmt.Println(min, max)
}

func readInts(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &r[i])
	}
	return r
}

func allPossibleSums(ns []int) []int {
	r := make([]int, size)
	for ignore := 0; ignore < size; ignore++ {
		r[ignore] = sumIgnoring(ns, ignore)
	}
	return r
}

func sumIgnoring(ns []int, ignore int) int {
	sum := 0
	for i := 0; i < size; i++ {
		if i != ignore {
			sum += ns[i]
		}
	}
	return sum
}

func minAndMax(s []int) (min, max int) {
	if len(s) < 1 {
		panic("cannot get minimum of empty slice")
	}
	min, max = s[0], s[0]
	for _, e := range s[1:] {
		if e < min {
			min = e
			continue
		}
		if e > max {
			max = e
		}
	}
	return min, max
}
