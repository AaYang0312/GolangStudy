package main

import "fmt"

type Clousre1 struct {
	n *int
}

func (c *Clousre1) Call() (id int) {
	id = *c.n
	(*c.n)++
	return
}

type Closure10 struct {
	n *int
}

func (c Closure10) Call() (id int) {
	id = *c.n
	(*c.n) += 10
	return
}
func newIDGenerators(start int) (
	*Clousre1,
	*Closure10,
) {
	n := start
	return &Clousre1{&n}, &Closure10{&n}
}
func main() {
	idGen1, idGen10 := newIDGenerators(1)
	fmt.Println(idGen1.Call())  // 1
	fmt.Println(idGen1.Call())  // 2
	fmt.Println(idGen10.Call()) // 3
	fmt.Println(idGen10.Call()) // 13
}
