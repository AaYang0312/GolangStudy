package main

import "fmt"

// 接口是一个规范
type Usber interface {
	start()
	stop()
}
type Computer struct {
}

func (c Computer) work(usb Usber) {
	// 要判断 usb 的类型，此处判断 usb 对象是否为 Phone
	if _, ok := usb.(Phone); ok {
		usb.start()
	} else {
		usb.stop()
	}

}

// 如果接口中有方法，必须通过结构体或自定义类型实现
type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Printf("%v启动\n", p.Name)
}
func (p Phone) stop() {
	fmt.Printf("%v关机\n", p.Name)
}

func main() {
	var p1 = Phone{
		"Apple",
	}
	p1.start()

	// 实现 usb 接口
	var usbp1 Usber

	usbp1 = p1
	usbp1.stop()

	var c1 = Computer{}
	c1.work(p1) // p1 符合接口规范，可以调用
	// 只要变量含有接口类型的所有方法，那么变量就实现了这个接口
}
