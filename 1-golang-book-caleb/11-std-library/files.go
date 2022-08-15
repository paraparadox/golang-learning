package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		// handle the error
		return
	}
	//defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)

	file.Close()

	// Reading a file via io/ioutil
	bs, err = ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	str = string(bs)
	fmt.Println(str)

	file, err = os.Create("test.txt")
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}
	//defer file.Close()

	_, err = file.WriteString("test")
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}
	file.Close()

	// Listing a catalog
	dir, err := os.Open("../")
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}
	//defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name(), fi.Mode())
	}

	dir.Close()

	// Listing catalog, recursively
	err = filepath.Walk("../", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}
}