package main

import "fmt"

// 抽象的银行业务员
type AbstractBanker interface {
	DoBusi() // 抽象的处理业务的接口
}

// 实现银行业务员接口
type SaveBanker struct{}

func (sb *SaveBanker) DoBusi() {
	fmt.Println("进行了存款")
}

// 转账的业务员
type TransferBanker struct {
}

func (tb *TransferBanker) DoBusi() {
	fmt.Println("进行了转账")
}

// 股票的业务员
type SharesBanker struct {
}

func (sb *SharesBanker) DoBusi() {
	fmt.Println("进行了股票交易")
}

// 实现一个架构层（基于抽象层进行业务封装-针对interface接口进行封装）
func BankBusiness(banker AbstractBanker) {
	banker.DoBusi()
}

func main() {
	// 存款业务
	// sb := SaveBanker{}
	// sb.DoBusi()

	// 转账业务
	// tb := TransferBanker{}
	// tb.DoBusi()

	// 股票业务
	// sb1 := SharesBanker{}
	// sb1.DoBusi()

	// 存款业务
	BankBusiness(&SaveBanker{})
	// 转账业务
	BankBusiness(&TransferBanker{})
	// 股票业务
	BankBusiness(&SharesBanker{})

}
