package main

import "fmt"

type ClothesShop struct {
}

func (c *ClothesShop) Style() {
	fmt.Println("逛街的装扮")
}

type ClothesWork struct{}

func (c *ClothesWork) Style() {
	fmt.Println("工作的装扮")
}

func main() {
	cw := ClothesWork{}
	cw.Style()

	cs := ClothesShop{}
	cs.Style()
}

// type Clothes struct {
// }

// func (c *Clothes) OnWork() {
// 	fmt.Println("工作的装扮")
// }

// func (c *Clothes) OnShop() {
// 	fmt.Println("逛街的装扮")
// }

// func main() {
// 	c := Clothes{}

// 	// 工作的业务
// 	fmt.Println("在工作")
// 	c.OnWork()

// 	fmt.Println("在逛街")
// 	c.OnShop()
// }
