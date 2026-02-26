package main

import (
	"fmt"
	"reflect"
)

func reflectFn(x interface{}) {
	v := reflect.ValueOf(x)

	kind := v.Kind()
	switch kind {
	case reflect.Int64:
		fmt.Println(v.Int())
	case reflect.Bool:
		fmt.Println(v.Bool())
	case reflect.Float32:
		fmt.Println(v.Float())
	case reflect.String:
		fmt.Println(v.String())
	default:
		fmt.Printf("暂未支持的类型: %v\n", kind)
	}
}
func reflectSetValue(x interface{}) {
	// 通过类型断言修改
	// v, _ := x.(*int64)
	// *v = 120

	// 通过反射修改
	v := reflect.ValueOf(x)
	// fmt.Println(v.Kind()) // ptr
	kind := v.Elem().Kind()
	switch kind {
	case reflect.Int64:
		if v.Elem().CanSet() == true {
			v.Elem().SetInt(120)
		}
	case reflect.String:
		if v.Elem().CanSet() == true {
			v.Elem().SetString("你好golang")
		}
	default:
		fmt.Println("暂不支持")
	}

}
func main() {
	// var a float32 = 3.14
	// var b int64 = 100
	// var c string = "Shanghai"
	// cp := &c
	// reflectFn(a)
	// reflectFn(b)
	// reflectFn(c)
	// reflectFn(cp)

	// 通过Setxx()修改变量值
	var a int64 = 100
	var s string = "golang"
	a1 := &a
	s1 := &s
	reflectSetValue(a1)
	reflectSetValue(s1)
	fmt.Println(int64(*a1))
	fmt.Println(string(*s1))
}
