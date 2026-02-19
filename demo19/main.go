package main

import (
	"fmt"
	"strings"
)

func main() {
	// 元素为 map 类型的切片
	//var userinfo = []string{"张三","李四"}

	// var userinfo = make([]map[string]string, 2)

	// fmt.Println(userinfo[0])        // map[]
	// fmt.Println(userinfo[0] == nil) // true map不初始化的默认值nil

	// if userinfo[0] == nil {
	// 	userinfo[0] = make(map[string]string)
	// 	userinfo[0]["username"] = "张三"
	// 	userinfo[0]["age"] = "20"
	// 	userinfo[0]["sex"] = "male"
	// }
	// if userinfo[1] == nil {
	// 	userinfo[1] = make(map[string]string)
	// 	userinfo[1]["username"] = "李四"
	// 	userinfo[1]["age"] = "21"
	// 	userinfo[1]["sex"] = "male"
	// }
	// // fmt.Println(userinfo[0]["username"])
	// for _, v := range userinfo {
	// 	// fmt.Println(v)
	// 	for key, value := range v {
	// 		fmt.Println(key, value)
	// 	}
	// }

	// 值为切片类型的 map
	// var userinfo = map[string][]string {
	// 	"hobby" : {"吃饭","睡觉","打游戏"},
	// }

	// var userinfo = make(map[string][]string)
	// userinfo["hobby"] = []string{
	// 	"吃饭",
	// 	"睡觉",
	// 	"打游戏",
	// }
	// userinfo["study"] = []string{
	// 	"golang",
	// 	"php",
	// 	"前端",
	// }
	// for _, v := range userinfo {
	// 	for _, value := range v {
	// 		fmt.Printf("%v\n", value)
	// 	}
	// }

	// var userinfo1 = make(map[string]string)

	// userinfo1["username"] = "张三"
	// userinfo1["age"] = "20"
	// userinfo2 := userinfo1
	// fmt.Println(userinfo1)
	// fmt.Println(userinfo2)
	// userinfo2["age"] = "50" // 影响userinfo1,map为引用数据类型
	// fmt.Println(userinfo1)
	// fmt.Println(userinfo2)

	// map1 := make(map[int]int, 10)
	// map1[10] = 100
	// map1[1] = 13
	// map1[4] = 56
	// map1[8] = 90

	// fmt.Println(map1)
	// for key, value := range map1 {
	// 	fmt.Printf("%v--%v\n", key, value)
	// }

	// 练习: 按key升序输出map里的key=>value (签名算法)
	/*
		1 13
		4 56
		10 100
		8 90
	*/
	// map1 := make(map[int]int, 10)
	// map1[10] = 100
	// map1[1] = 13
	// map1[4] = 56
	// map1[8] = 90

	// // 1、把map的key放在切片里
	// var keySlice []int
	// for key, _ := range map1 {
	// 	keySlice = append(keySlice, key)
	// }
	// fmt.Println(keySlice) // [4 8 10 1]

	// // 2、让key升序排序
	// sort.Ints(keySlice)

	// // 3、循环遍历key来输出map
	// for _, val := range keySlice {
	// 	fmt.Println(val, "--", map1[val])
	// }

	// 练习: 统计一个字符串中每个单词出现的次数
	strArr := "what is your name"
	var strSlice = strings.Split(strArr, " ")
	var wordCount = make(map[string]int)
	for _, v := range strSlice {
		wordCount[v] += 1
	}
	for key, val := range wordCount {
		fmt.Printf("%v:%v\n", key, val)
	}
}
