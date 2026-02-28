package main

import (
	"fmt"
	"os"
)

func main() {
	readSlice, err := os.ReadFile("./main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(readSlice))
}
