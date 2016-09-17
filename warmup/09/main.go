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

const (
	minArraySize  = 1
	maxArraySize  = 100000
	minArrayValue = 1
	maxArrayValue = 100000
	minRotations  = 1
	maxRotations  = 100000
	minQueries    = 1
	maxQueries    = 500
)

func main() {
	// there are some really long lines in the test so we will use a raw
	// bufio.Reader instead of a bufio.Scanner.
	br := bufio.NewReader(os.Stdin)

	arraySize, rotations, nQueries, err := readFirstLine(br)
	if err != nil {
		log.Fatal(err)
	}

	a, err := readArray(arraySize, br)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < nQueries; i++ {
		index, err := readQuery(br)
		if err != nil {
			log.Fatal(err)
		}
		index = indexBeforeRotations(index, arraySize, rotations)
		fmt.Println(a[index])
	}
}

func readFirstLine(br *bufio.Reader) (int, int, int, error) {
	line, err := br.ReadString('\n')
	if err != nil {
		return 0, 0, 0, err
	}
	var arraySize, rotations, nQueries int
	_, err = fmt.Sscanf(line, "%d %d %d",
		&arraySize, &rotations, &nQueries)
	if err != nil {
		return 0, 0, 0, err
	}

	if arraySize < minArraySize {
		return 0, 0, 0, fmt.Errorf("arraySize too small")
	}
	if arraySize > maxArraySize {
		return 0, 0, 0, fmt.Errorf("arraySize too big")
	}

	if rotations < minRotations {
		return 0, 0, 0, fmt.Errorf("rotations too small")
	}
	if rotations > maxRotations {
		return 0, 0, 0, fmt.Errorf("rotations too big")
	}

	if nQueries < minQueries {
		return 0, 0, 0, fmt.Errorf("nQueries too small")
	}
	if nQueries > maxQueries {
		return 0, 0, 0, fmt.Errorf("nQueries too big")
	}

	return arraySize, rotations, nQueries, err
}

func readArray(sz int, br *bufio.Reader) ([]int, error) {
	line, err := readLongLine(br)
	if err != nil {
		return nil, err
	}
	line = strings.TrimSpace(line)
	words := strings.Split(line, " ")
	if len(words) != sz {
		return nil, fmt.Errorf(
			"wrong number of elements: %d was read, but %d was expected",
			len(words), sz)
	}
	nums := make([]int, sz)
	for i, w := range words {
		nums[i], err = strconv.Atoi(w)
		if err != nil {
			return nil, fmt.Errorf(
				"invalid format element %d: %s",
				i, err)
		}
	}
	return nums, nil
}

func readLongLine(br *bufio.Reader) (string, error) {
	var buf []byte
	var chunk []byte
	pendingData := true
	var err error
	for pendingData {
		chunk, pendingData, err = br.ReadLine()
		if err != nil {
			return "", err
		}
		buf = append(buf, chunk...)
	}
	return string(buf), nil
}

func readQuery(br *bufio.Reader) (int, error) {
	line, err := br.ReadString('\n')
	if err != nil && err != io.EOF {
		return 0, err
	}
	line = strings.TrimSpace(line)
	n, err := strconv.Atoi(line)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func indexBeforeRotations(after, sz, rot int) int {
	before := after - (rot % sz)
	if before < 0 {
		before = sz + before
	}
	return before
}
