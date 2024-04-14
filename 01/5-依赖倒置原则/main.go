package main

import "fmt"

// 抽象层
type Car interface {
	Run()
}
type Driver interface {
	Drive(c Car)
}

// 实现层

type Benz struct {
	// ...
}

func (b *Benz) Run() {
	fmt.Println("Benz is running ...")
}

type Bmw struct {
	// ...
}

func (b *Bmw) Run() {
	fmt.Println("Bmw is running ...")
}

type Zhang3 struct{}

func (z3 *Zhang3) Drive(car Car) {
	fmt.Println("Zhang3 dirve car")
	car.Run()
}

type Li4 struct{}

func (l4 *Li4) Drive(car Car) {
	fmt.Println("Li4 dirve car")
	car.Run()
}

// 如果要增加一个王五开奔驰
type Wang5 struct{}

func (w5 *Wang5) Drive(car Car) {
	fmt.Println("Wang5 dirve car")
	car.Run()
}

// 业务逻辑层

func main() {
	// 张三开奔驰
	var benz Car
	benz = new(Benz)

	var zhangsan Driver
	zhangsan = new(Zhang3)
	zhangsan.Drive(benz)

	var bmw Car
	bmw = new(Bmw)
	var lisi Driver
	lisi = new(Li4)
	lisi.Drive(bmw)

	zhangsan.Drive(bmw)

	var wangwu Driver
	wangwu = new(Wang5)
	wangwu.Drive(benz)
}
