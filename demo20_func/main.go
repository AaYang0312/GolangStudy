package main

// func sumFn(x int, y int) int {
// 	sum := x + y
// 	return sum
// }

// 函数参数的简写
// func sumFn(x, y int) int {
// 	sum := x + y
// 	return sum
// }

// 函数的可变参数
// func sumFn1(x ...int) {
// 	fmt.Printf("%v--%T", x, x) // [13 4 5 6]--[]int
// }
// func sumFn1(x ...int) int {
// 	sum := 0
// 	for _, v := range x {
// 		sum += v
// 	}
// 	return sum
// }

// func sumAveFn2(x ...int) (sum int, ave float64) {
// 	for _, v := range x {
// 		sum += v
// 	}
// 	return sum, float64(sum) / float64(len(x))
// }
// func main() {
// 	sum1, ave1 := sumAveFn2(13, 4, 5, 6)
// 	fmt.Println(sum1, ave1)
// }

// 练习：封装方法实现切片升序降序排序
// func sliceSortAsc(slice []int) []int {
// 	for i := 0; i < len(slice)-1; i++ {
// 		for j := i + 1; j < len(slice); j++ {
// 			if slice[i] > slice[j] {
// 				slice[i], slice[j] = slice[j], slice[i]
// 			}
// 		}
// 	}
// 	return slice
// }

// func sliceSortDesc(slice []int) []int {
// 	for i := 0; i < len(slice)-1; i++ {
// 		for j := i + 1; j < len(slice); j++ {
// 			if slice[i] < slice[j] {
// 				slice[i], slice[j] = slice[j], slice[i]
// 			}
// 		}
// 	}
// 	return slice
// }
// func main() {

// 	var sliceA = []int{12, 34, 37, 35, 556, 36, 2}
// 	sliceB := make([]int, len(sliceA))
// 	sliceSortAsc(sliceA)
// 	copy(sliceB, sliceA)
// 	SliceC := sliceSortDesc(sliceA)
// 	fmt.Println(sliceB)
// 	fmt.Println(SliceC)
// }

// 练习：封装方法按照key升序排序，最后输出字符串

// func main() {
// 	m1 := map[string]string{
// 		"username": "zhangsan",
// 		"age":      "20",
// 		"sex":      "男",
// 		"height":   "180",
// 	}
// 	fmt.Println(mapSort(m1))
// }
// func mapSort(map1 map[string]string) string {
// 	var sliceKey []string

// 	// 1、把 map 对象的 key 放在切片里
// 	for k, _ := range map1 {
// 		sliceKey = append(sliceKey, k)
// 	}
// 	// 2、对 key 升序排序，按ASCII码
// 	sort.Strings(sliceKey)

// 	var str string
// 	for _, v := range sliceKey {
// 		str += fmt.Sprintf("%v=>%v", v, map1[v])
// 	}
// 	return str
// }
