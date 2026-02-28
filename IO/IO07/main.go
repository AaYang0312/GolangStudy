package main

import (
	"fmt"
	"os"
)

func main() {
	str := "hello golang"
	err := os.WriteFile("./test.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}
