package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	champPath := os.Getenv("CHAMPIONS_PATH")
	files, err := ioutil.ReadDir(champPath)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
	fmt.Printf("Total: %v champions", len(files))
}
