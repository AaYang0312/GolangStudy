package main

import "fmt"

func main() {
	// var arr1 []int
	// var arr2 = []int{1, 2, 4, 6}
	// fmt.Println(arr1)        // []
	// fmt.Println(arr1 == nil) // true 切片默认值为nli

	// fmt.Println(arr2 == nil) // false

	// 基于数组定义切片
	// a := [5]int{1, 2, 3, 4, 5}
	// var b []int

	// b = a[:]                   // 获取a中所有值 [1 2 3 4 5]
	// c := a[1:3]                // [2,3]
	// d := a[2:3]                // [3]
	// e := a[2:]                 // [3,4,5]
	// f := a[:3]                 // [1,2,3]
	// fmt.Println(b, c, d, e, f) //
	// // 类型：[]int-切片 [n]int-数组

	// // 切片的长度和容量len(),cap(),长度为元素个数，容量是底层数组的长度减去切片的起始索引
	// s1 := []int{2, 3, 4, 5, 6, 7}
	// s2 := s1[2:]  // [4,5,6,7]
	// s3 := s1[1:3] // [3,4]
	// fmt.Printf("s2长度为%v,容量为%v\ns3长度为%v,容量为%v\n", len(s2), cap(s2), len(s3), cap(s3))

	// // 声明切片时可以使用下标索引
	// var arr1 []int
	// var arr2 = []int{1, 2, 3, 4}
	// arr3 := []int{1, 2, 3, 4}
	// arr4 := []int{1: 2, 2: 3, 4: 6} // [0 2 3 0 6]
	// fmt.Println(arr1)
	// fmt.Println(arr2)
	// fmt.Println(arr3)
	// fmt.Println(arr4, cap(arr4))

	// 使用 make 函数构造切片
	// // make([]T, size, cap) T:切片的元素类型, size:切片中元素的数量, cap:切片容量
	// a1 := make([]int, 4, 8)
	// fmt.Println(a1) // [0 0 0 0]
	// a1[2] = 1
	// fmt.Println(a1) // [0 0 1 0]
	// a1 = append(a1, 1)
	// fmt.Println(a1, len(a1), cap(a1)) // [0 0 1 0 1] 5 8
	// // 容量用来限制切片在内存中持有的容量

	// //a3 := []int{4, 5, 7}
	// a1 = append(a1, a1...)
	// fmt.Println(a1, len(a1), cap(a1)) // [0 0 1 0 1 0 0 1 0 1] 10 16

	// 切片的扩容策略
	var sliceA []int
	for i := 0; i < 10; i++ {
		sliceA = append(sliceA, i)
		fmt.Printf("%v 长度:%d 容量:%d\n", sliceA, len(sliceA), cap(sliceA))
		/*
			[0] 长度:1 容量:1
			[0 1] 长度:2 容量:2
			[0 1 2] 长度:3 容量:4
			[0 1 2 3] 长度:4 容量:4
			[0 1 2 3 4] 长度:5 容量:8
			[0 1 2 3 4 5] 长度:6 容量:8
			[0 1 2 3 4 5 6] 长度:7 容量:8
			[0 1 2 3 4 5 6 7] 长度:8 容量:8
			[0 1 2 3 4 5 6 7 8] 长度:9 容量:16
			[0 1 2 3 4 5 6 7 8 9] 长度:10 容量:16
		*/
		// 如果新容量大于两倍旧容量，则最终容量为新申请的容量，
		// 前1024元素采取两倍扩容，之后循环增加原容量的1/4，直到大于等于新申请的容量
	}
}
