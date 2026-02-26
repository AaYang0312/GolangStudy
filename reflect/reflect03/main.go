package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name" form:"usrname"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
	fmt.Println("写入完成...")
}
func (s Student) GetInfo() string {
	var str = fmt.Sprintf("姓名：%v, 年龄：%d, 分数：%d", s.Name, s.Age, s.Score)
	return str
}
func (s Student) print() {
	fmt.Println("这是一个打印方法...")
}

// 打印结构体字段
func PrintStructField(s interface{}) {
	// 判断是否为struct
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct && reflect.TypeOf(s).Elem().Kind() != reflect.Struct {
		fmt.Println("不是struct类型!")
		return
	}
	// 1. Field获取结构体字段
	field0 := t.Field(0)
	// fmt.Printf("%#v\n", field0)
	fmt.Println("字段名称：", field0.Name)
	fmt.Println("字段类型：", field0.Type)
	fmt.Println("字段标签1：", field0.Tag.Get("json"))
	fmt.Println("字段标签2：", field0.Tag.Get("form"))
	// 2. FieldByName获取结构体字段
	n := "Name"
	fieldName, iru := t.FieldByName(n)
	if iru == true {
		fmt.Println("字段名称：", fieldName.Name)
		fmt.Println("字段类型：", fieldName.Type)
		fmt.Println("字段标签1：", fieldName.Tag.Get("json"))
		fmt.Println("字段标签2：", fieldName.Tag.Get("form"))
	} else {
		fmt.Println("未找到字段", n)
	}
	// 3. NumField获取字段个数
	fieldCount := t.NumField()
	fmt.Printf("结构体有%v个属性\n", fieldCount)

	// 4. 值变量获取结构体属性对应的值
	v := reflect.ValueOf(s)
	// fmt.Printf("reflect.type: %#v, reflect.value: %#v\n", t, v)
	// index := []int{0}
	// fmt.Printf("%#v\n", v.FieldByIndex(index))
	v0 := v.Field(0)
	fmt.Println("结构体Name: ", v0)
	fmt.Println("结构体Age: ", v.FieldByName("Age"))
	fmt.Println("结构体Name: ", v.FieldByIndex([]int{2}))
}

// 打印结构体方法
func PrintStructFn(s interface{}) {
	t := reflect.TypeOf(s)
	// v := reflect.ValueOf(s)

	// 1. 类型变量Method获取结构体的方法
	// Method() 和结构体方法的顺序无关，和方法的 ASCII 码有关
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("第%d个方法为%v, 类型为%v\n", i, t.Method(i).Name, t.Method(i).Type)
	}
	// 第二种方法按名称
	ma, _ := t.MethodByName("Print")
	fmt.Printf("%v\n", ma)
	// 2. 类型变量获取结构体有多少个方法

	// 3. 调用方法

}

func main() {
	stu1 := Student{
		"Aemeath",
		18,
		100,
	}
	// PrintStructField(stu1)
	PrintStructFn(stu1)
}
