package main

import "fmt"

type Banker struct{}

// 存款业务
func (b *Banker) Save() {
	fmt.Println("进行了存款业务")
}

// 转账业务
func (b *Banker) Transfer() {
	fmt.Println("进行了转账业务")
}

// 支付业务
func (b *Banker) Pay() {
	fmt.Println("进行了支付业务")
}

// 增加了一个功能，股票
func (b *Banker) Gupiao() {
	fmt.Println("股票交易")
}
