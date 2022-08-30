package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("task-3.5-14-data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	i := 0
	for {
		s, err := reader.ReadString(';')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}
		i++
		if s == "0;" {
			break
		}
	}
	fmt.Println(i)
}
