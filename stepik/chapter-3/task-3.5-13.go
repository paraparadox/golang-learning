package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func walkFn(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	f, _ := os.Open(path)
	defer f.Close()
	r := csv.NewReader(f)
	rows, _ := r.ReadAll()
	if len(rows) == 10 {
		fmt.Println(rows[4][2])
	}
	return nil
}

func main() {
	const root = "./task"
	err := filepath.Walk(root, walkFn)
	if err != nil {
		fmt.Println(err)
	}
}
