package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const now = 1589570165

func main() {
	stdTime := time.Unix(now, 0)
	stdTime = stdTime.UTC()
	durationString, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	durationString = strings.Trim(durationString, "\n")
	durationString = strings.Replace(durationString, " мин. ", "m", 1)
	durationString = strings.Replace(durationString, " сек.", "s", 1)
	duration, _ := time.ParseDuration(durationString)
	stdTime = stdTime.Add(duration)
	fmt.Println(stdTime.Format(time.UnixDate))
}
