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
	s := bufio.NewScanner(os.Stdin)

	a, err := readRating(s)
	if err != nil {
		log.Fatalf("reading first rating: %s", err)
	}

	b, err := readRating(s)
	if err != nil {
		log.Fatalf("reading second rating: %s", err)
	}

	ca, cb := a.compare(b)

	fmt.Println(ca, cb)
}

type rating [3]int

func readRating(s *bufio.Scanner) (*rating, error) {
	if !s.Scan() {
		return nil, fmt.Errorf("reading score: %s", s.Err())
	}
	line := s.Text()
	words := strings.Split(line, " ")
	if len(words) != 3 {
		return nil, fmt.Errorf(
			"invalid score format: expected 3 words but %d was found",
			len(words))
	}

	r := &rating{}
	for i, w := range words {
		n, err := strconv.Atoi(w)
		if err != nil {
			return nil, fmt.Errorf(
				"cannot read score at word %d: %s",
				i+1, err)
		}
		if err = checkScore(n); err != nil {
			return nil, fmt.Errorf(
				"invalid score at word %d: %s",
				i+1, err)
		}
		r[i] = n
	}

	return r, nil
}

func checkScore(n int) error {
	if n < 1 {
		return fmt.Errorf("too small")
	}
	if n > 100 {
		return fmt.Errorf("too big")
	}
	return nil
}

func (a *rating) compare(b *rating) (int, int) {
	aRet := 0
	bRet := 0
	for i, sa := range a {
		sb := b[i]
		switch {
		case sa > sb:
			aRet++
		case sa < sb:
			bRet++
		default:
		}
	}
	return aRet, bRet
}
