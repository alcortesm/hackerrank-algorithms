package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	var line string
	_, err := fmt.Scan(&line)
	if err != nil {
		log.Fatal(err)
	}
	line = strings.TrimSpace(line)

	t, err := time.Parse("03:04:05PM", line)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t.Format("15:04:05"))
}
