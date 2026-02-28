package main

import (
	"bufio"
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

	// 2.写入bufio
	len := 0
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		lentemp, err := writer.WriteString("写入第" + strconv.Itoa(i) + "行\r\n")
		writer.Flush() // 将缓存中的内容写入到文件中
		if err != nil {
			fmt.Println(err)
			return
		}
		len += lentemp
	}
	fmt.Printf("成功写入%d个字节", len)
}
