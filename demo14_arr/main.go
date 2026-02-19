package main

func main() {

	// 1、数组长度也是类型的一部分
	// var arr1 [3]int
	// var arr2 [4]int
	// var strArr [3]string

	// fmt.Printf("arr1:%T, arr2:%T, strArr:%T", arr1, arr2, strArr)

	// 2、数组的初始化
	// var arr1 [3]int
	// var strArr [3]string
	// fmt.Println(arr1)   // [0 0 0]
	// fmt.Println(strArr) // [  ]

	// 第一种方法
	// var arr1 [3]int
	// arr1[0] = 1
	// arr1[1] = 2
	// arr1[2] = 4
	// fmt.Println(arr1)

	// 第二种方法

	// var arr1 = [3]int{12, 21, 4}
	// fmt.Println(arr1)

	// 第三种

	// var numArr = [...]int{1, 2, 3, 4, 5}
	// fmt.Println(len(numArr))

	// arr1 := [...]string{"php", "nodejs", "golang"}
	// fmt.Println(len(arr1))

	// 第四种

	// a := [...]string{1: "php", 3: "java"}

	// fmt.Printf("type of a:%T\n", a)
	// fmt.Println(a)

	// 第五种，循环遍历

	// var arr1 = [3]int{23, 34, 5}
	// for i := 0; i < len(arr1); i++ {
	// 	arr1[i] = 250
	// }
	// fmt.Println(arr1)

	// var arr2 = [...]string{"php", "java", "c++", "golang"}
	// for k, v := range arr2 {
	// 	fmt.Printf("k=%v, v=%v\n", k, v)
	// }

	// 求数组里所有数的和和平均值，用for和for range实现
	var arr = [...]int{1, 56, 23, 45, 65, 45}

	sum := 0
	// for i := 0; i < len(arr); i++ {
	// 	sum += arr[i]
	// }
	// ave = sum / len(arr)
	// fmt.Printf("sum=%v, average=%v", sum, ave)

	// for _, v := range arr {
	// 	sum += v
	// }
	// fmt.Printf("sum=%v, ave=%.2f", sum, float64(sum)/float64(len(arr)))

}
