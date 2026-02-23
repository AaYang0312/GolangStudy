package main

import "fmt"

type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

func (p *Phone) start() { // 指针类型接收者，如果是指针类型接收者必须用下面方法实现接口
	fmt.Printf("%v启动\n", p.Name)
}
func (p *Phone) stop() {
	fmt.Printf("%v关闭\n", p.Name)
}

func main() {
	// var p1 = Phone{
	// 	"小米手机",
	// }
	// var p2 Usber = p1
	// p2.start()

	var p3 = &Phone{
		"苹果手机",
	}
	var p4 Usber = p3
	p4.start()
}
