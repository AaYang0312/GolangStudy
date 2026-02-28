package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 1.读取文件 os.Open()为只读打开
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("出错:", err)
		return
	}
	defer file.Close()

	// 2.写入字符串
	lens := 0
	for i := 0; i < 10; i++ {
		tempLens, err := file.WriteString("223" + strconv.Itoa(i) + "\r\n") // 回车必须用\r\n
		if err != nil {
			fmt.Println(err)
			return
		}
		lens += tempLens
	}
	fmt.Println("成功写入", lens, "个字节！")
}
