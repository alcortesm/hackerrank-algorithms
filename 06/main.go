package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	a, err := readArray()
	if err != nil {
		log.Fatal(err)
	}

	report(a)
}

func readArray() ([]int, error) {
	s := bufio.NewScanner(os.Stdin)

	n, err := readArraySize(s)
	if err != nil {
		return nil, fmt.Errorf("reading array size: %s", err)
	}

	a, err := readArrayContents(s, n)
	if err != nil {
		return nil, fmt.Errorf("reading array contents: %s", err)
	}

	return a, nil
}

func readArraySize(s *bufio.Scanner) (int, error) {
	if !s.Scan() {
		return 0, s.Err()
	}
	n, err := strconv.Atoi(s.Text())
	if err != nil {
		return 0, err
	}
	if err = checkArraySize(n); err != nil {
		return 0, err
	}
	return n, nil
}

func checkArraySize(n int) error {
	if n < 0 {
		return fmt.Errorf("negative array size")
	}
	return nil
}

func readArrayContents(s *bufio.Scanner, n int) ([]int, error) {
	if !s.Scan() {
		return nil, s.Err()
	}
	line := s.Text()

	words := strings.Split(line, " ")
	if len(words) != n {
		return nil, fmt.Errorf(
			"bad number of elements: expected %d, was %d",
			n, len(words))
	}

	ints, err := stringsToInts(words)
	if err != nil {
		return nil, err
	}

	return ints, nil
}

func stringsToInts(ss []string) ([]int, error) {
	ii := make([]int, 0, len(ss))
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ii = append(ii, i)
	}
	return ii, nil
}

func report(a []int) {
	var pos, neg, zeros int
	for _, e := range a {
		switch {
		case e > 0:
			pos++
		case e < 0:
			neg++
		default:
			zeros++
		}
	}
	format := "%.6f\n"
	fmt.Printf(format, fraction(pos, len(a)))
	fmt.Printf(format, fraction(neg, len(a)))
	fmt.Printf(format, fraction(zeros, len(a)))
}

func fraction(n int, l int) float64 {
	return float64(n) / float64(l)
}
