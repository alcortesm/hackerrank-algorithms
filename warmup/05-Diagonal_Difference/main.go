package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	s, err := readSquare(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s.differenceOfDiagonalSums())
}

type square [][]int

func readSquare(r io.Reader) (square, error) {
	s := bufio.NewScanner(r)

	n, err := scanSquareSize(s)
	if err != nil {
		return nil, fmt.Errorf("reading square size: %s", err)
	}

	rows := make([][]int, n)

	for r := 0; r < n; r++ {
		rows[r] = make([]int, n)
		if err := scanRow(s, rows[r]); err != nil {
			return nil, fmt.Errorf(
				"scanning row %d: %s", r+1, err)
		}
	}

	return rows, nil
}

func scanSquareSize(s *bufio.Scanner) (int, error) {
	if !s.Scan() {
		return 0, s.Err()
	}
	n, err := strconv.Atoi(s.Text())
	if err != nil {
		return 0, fmt.Errorf("bad format: %s", err)
	}
	if err = checkSquareSize(n); err != nil {
		return 0, fmt.Errorf("invalid value: %s", err)
	}
	return n, nil
}

func checkSquareSize(n int) error {
	if n < 1 {
		return fmt.Errorf("too small")
	}
	return nil
}

func scanRow(s *bufio.Scanner, row []int) error {
	if !s.Scan() {
		return s.Err()
	}
	words := strings.Split(s.Text(), " ")
	if len(words) != len(row) {
		return fmt.Errorf(
			"unexpected number of elements, %d was expected, but %d were found",
			len(row), len(words))
	}
	for i, w := range words {
		n, err := strconv.Atoi(w)
		if err != nil {
			return fmt.Errorf(
				"element %d: invalid format: %s",
				i+1, err)
		}
		row[i] = n
	}
	return nil
}

func (s square) differenceOfDiagonalSums() int {
	sumOne := 0
	for i := 0; i < len(s); i++ {
		sumOne += s[i][i]
	}

	sumTwo := 0
	for i := 0; i < len(s); i++ {
		sumTwo += s[i][len(s)-1-i]
	}

	diff := sumOne - sumTwo
	if diff < 0 {
		diff = -diff
	}
	return diff
}
