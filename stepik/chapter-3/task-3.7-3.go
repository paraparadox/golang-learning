package main

import (
	"fmt"
	"time"
)

func main() {
	var s string
	fmt.Scanln(&s)
	firstTime, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return
	}
	fmt.Println(firstTime.Format(time.UnixDate))
}
