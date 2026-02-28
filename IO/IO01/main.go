package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 1.读取文件
	file, err := os.Open("./main.go")

	if err != nil {
		fmt.Println("出错:", err)
		return
	}
	defer file.Close()

	// 2.读取内容
	var strSlice []byte
	var tempSlice = make([]byte, 128)
	for {
		n, err := file.Read(tempSlice)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取失败")
			return
		}
		// fmt.Printf("读取到了%d个字节\n", n)
		strSlice = append(strSlice, tempSlice[:n]...) // 注意写法，避免temp最后未被完全覆盖
	}
	fmt.Printf("读取到内容为: %v\n", string(strSlice))

}
