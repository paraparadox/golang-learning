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
	date, _ := time.Parse("2006-01-02 15:04:05", dateString)
	h, m, s := date.Clock()
	if h > 13 {
		date = date.Add(time.Hour * 24)
	} else if h == 13 && (m > 0 || s > 0) {
		date = date.Add(time.Hour * 24)
	}
	fmt.Println(date.Format("2006-01-02 15:04:05"))
}

// 2020-05-15 08:00:00
