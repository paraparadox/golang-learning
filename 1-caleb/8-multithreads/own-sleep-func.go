package main

import (
	"fmt"
	"time"
)

func Sleep(n int) {
	amt := time.Duration(n)
	<-time.After(time.Second * amt)
}

func main() {
	fmt.Println("Start\n")
	n := 2
	fmt.Println("Waiting for", n, "seconds")
	Sleep(n)
	fmt.Println("Done!\n")
	n = 3
	fmt.Println("Waiting for", n, "seconds")
	Sleep(n)
	fmt.Println("Done!\n")
	n = 1
	fmt.Println("Waiting for", n, "seconds")
	Sleep(n)
	fmt.Println("Done!\n")
	n = 4
	fmt.Println("Waiting for", n, "seconds")
	Sleep(n)
	fmt.Println("Done!\n")

}
