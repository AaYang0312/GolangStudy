package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读取文件方法2
func main() {
	// 打开文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 读取文件
	var strString string
	strReader := bufio.NewReader(file) // *os.file实现了io.Reader接口
	for {
		str, err := strReader.ReadString('\n')
		if err == io.EOF {
			strString += str
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		strString += str
	}
	fmt.Println(strString)

}
