package main

// json struct 相互转换
type Student struct {
	ID     int
	Gender string // 变量名必须大写开头，否则为私有变量，无法转换
	Name   string `json:"name"` // 或者指定一个json tag 结构体标签
	Sno    string
}

func main() {
	// 将 struct 转换为 json 方法

	// var p1 = Student{
	// 	ID:     12,
	// 	Gender: "male",
	// 	Name:   "张三",
	// 	Sno:    "s0001",
	// }
	// jsonByte, _ := json.Marshal(p1)

	// jsonstr := string(jsonByte) // 把 []byte 类型的切片转换为 string 类型

	// fmt.Printf("%v", jsonstr) // {"ID":12,"Gender":"male","Name":"张三","Sno":"s0001"}

	// 将 json 转换为 struct
	// var str = `{"ID":12,"Gender":"male","Name":"张三","Sno":"s0001"}` // 定义 json 字符串的方法
	// var s1 = Student{}
	// sp := &s1
	// var strbyte = make([]byte, 0)
	// strbyte = []byte(str)
	// err := json.Unmarshal(strbyte, sp) // 传入s1地址才能进行修改

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%#v\n", s1)

	// 将 struct 转换为 json 并保持变量小写
}
