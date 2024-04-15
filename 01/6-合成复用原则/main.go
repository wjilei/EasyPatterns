package main

import "fmt"

type Cat struct{}

func (c *Cat) Eat() {
	fmt.Println("小猫吃饭")
}

// 给小猫增加一个可以睡觉的方法（使用继承实现）
type CatB struct {
	Cat
}

func (b *CatB) Sleep() {
	fmt.Println("小猫睡觉")
}

// 给小猫增加一个可以睡觉的方法（使用组合的方式实现）
type CatC struct {
	C *Cat
}

func (cc *CatC) Eat() {
	cc.C.Eat()
}

func (cc *CatC) Sleep() {
	fmt.Println("小猫睡觉")
}

func main() {
	c := Cat{}
	c.Eat()

	fmt.Println("----------")
	cb := CatB{}
	cb.Eat()
	cb.Sleep()

	fmt.Println("-----------")
	cc := CatC{
		// C: &Cat{},
	}
	if cc.C != nil {
		cc.C.Eat()
	}
	cc.Eat()
	cc.Sleep()
}
