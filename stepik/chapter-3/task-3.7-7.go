package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	dateString, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	dateString = strings.Trim(dateString, "\n")
	dates := strings.Split(dateString, ",")
	first, _ := time.Parse("02.01.2006 15:04:05", dates[0])
	second, _ := time.Parse("02.01.2006 15:04:05", dates[1])
	var result time.Duration
	if first.Before(second) {
		result = second.Sub(first)
	} else {
		result = first.Sub(second)
	}
	fmt.Println(result.Round(time.Second))
}
