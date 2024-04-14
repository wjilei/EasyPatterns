package main

import "fmt"

// 司机张三，李四
// 汽车宝马，奔驰
// 1 - 张三开奔驰
// 2 - 李四开宝马

type Benz struct{}

func (b *Benz) Run() {
	fmt.Println("benz is runing...")
}

type Bmw struct{}

func (b *Bmw) Run() {
	fmt.Println("bmw is runing ...")
}

type Zhang3 struct{}

func (z3 *Zhang3) DriveBenz(benz *Benz) {
	fmt.Println("zhang3 Drive benz")
	benz.Run()
}

type Li4 struct {
}

func (l4 *Li4) DriveBmw(bmw *Bmw) {
	fmt.Println("Li4 Drive bmw")
	bmw.Run()
}

func main() {
	benz := Benz{}
	zhangsan := Zhang3{}
	zhangsan.DriveBenz(&benz)

	bmw := Bmw{}
	lisi := Li4{}
	lisi.DriveBmw(&bmw)

}
