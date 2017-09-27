package main

import (
	"fmt"
	"log"
)

const (
	minArraySize    = 1
	maxArraySize    = 1e5
	minCandleHeight = 1
	maxCandleHeight = 1e7
)

func main() {
	a, err := readArray()
	if err != nil {
		log.Fatal(err)
	}
	f, err := freqOfMax(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}

func readArray() ([]int, error) {
	size, err := readArraySize(minArraySize, maxArraySize)
	if err != nil {
		return nil, fmt.Errorf("reading array size: %s", err)
	}
	a, err := readArrayContents(size, minCandleHeight, maxCandleHeight)
	if err != nil {
		return nil, fmt.Errorf("reading array contents: %s", err)
	}
	return a, nil
}

func readArraySize(min, max uint) (uint, error) {
	var size uint
	if _, err := fmt.Scanf("%d\n", &size); err != nil {
		return 0, err
	}
	if err := checkInt(int(size), int(min), int(max)); err != nil {
		return 0, fmt.Errorf("array size: %v", err)
	}
	return size, nil
}

func checkInt(n, min, max int) error {
	if n < min {
		return fmt.Errorf("cannot be less than %d, was %d", min, n)
	}
	if n > max {
		return fmt.Errorf("cannot be bigger than %d, was %d", max, n)
	}
	return nil
}

func readArrayContents(size uint, min, max int) ([]int, error) {
	ns := make([]int, int(size))
	for i := 0; i < int(size); i++ {
		if _, err := fmt.Scanf("%d", &ns[i]); err != nil {
			return nil, fmt.Errorf("scanning element %d: %v", i, err)
		}
		if err := checkInt(ns[i], min, max); err != nil {
			return nil, fmt.Errorf("scanning element %d: %v", i, err)
		}
	}
	return ns, nil
}

func freqOfMax(a []int) (int, error) {
	if len(a) == 0 {
		return 0, fmt.Errorf("freqOfMax: the slice is empty")
	}
	max := a[0]
	count := 1
	for _, e := range a[1:] {
		if e > max {
			max = e
			count = 1
			continue
		}
		if e == max {
			count++
		}
	}
	return count, nil
}
