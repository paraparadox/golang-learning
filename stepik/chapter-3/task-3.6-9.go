package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	GlobalId int `json:"global_id"`
}

func main() {
	f, err := os.Open("./data.json")
	defer f.Close()
	if err != nil {
		return
	}

	items := make([]Item, 10)
	dec := json.NewDecoder(f)
	dec.Decode(&items)
	sum := 0
	for _, item := range items {
		sum += item.GlobalId
	}
	fmt.Println(sum)
}
