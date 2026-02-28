package main

import (
	"fmt"
	"os"
)

func main() {
	// 1.读取文件 os.Open()为只读打开
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("出错:", err)
		return
	}
	defer file.Close()

	// 2. 写入文件
	var writeSlice []byte
	writeSlice = append(writeSlice, '1', '2', '3')
	lenb, err := file.Write(writeSlice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("成功写入%d个字节！", lenb)
}
