package behavioral

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/**
模板模式，父类定义最佳实践流程，子类实现抽象方法。
适用于经过提炼之后的固定流程。
go只能通过组合实现，所以会看起来像策略模式哦
 */

/**
1 superClass 提供入口方法
2 interface 定义行为
3 触发时塞入实例
 */

type BankBusinessHandler interface {
  	// 排队拿号
	TakeRowNumber()

	WaitInHead()

	HandleBusiness()

	Commentate()

	//钩子方法
	CheckVipIdentity() bool
}

type BankBusinessExecutor struct {
	handler BankBusinessHandler
}

func NewBankBusinessExecutor(handler BankBusinessHandler) *BankBusinessExecutor {
	return &BankBusinessExecutor{
		handler,
	}
}
func (b *BankBusinessExecutor) ExecutorBankBusiness()  {
	// 适用于与客户端单次交互的流程
	// 如果需要与客户端多次交互才能完成整个流程，
	// 每次交互的操作去调对应模板里定义的方法就好，并不需要一个调用所有方法的模板方法
	b.handler.TakeRowNumber()
	if !b.handler.CheckVipIdentity() {
		b.handler.WaitInHead()
	}
	b.handler.HandleBusiness()
	b.handler.Commentate()
}

type BaseBusinessHandler struct {
}

func (p *BaseBusinessHandler) TakeRowNumber() {
	fmt.Println("请拿好您的取件码：" + strconv.Itoa(rand.Intn(100)) +
		" ，注意排队情况，过号后顺延三个安排")
}

func (p *BaseBusinessHandler) WaitInHead() {
	fmt.Println("排队等号中...")
	time.Sleep(1 * time.Second)
	fmt.Println("请去窗口xxx...")
}

func (p *BaseBusinessHandler) HandleBusiness() {
	fmt.Println("账户存储很多万人民币...")
}

func (p *BaseBusinessHandler) Commentate() {
	fmt.Println("请对我的服务作出评价，满意请按0，满意请按0，(～￣▽￣)～")
}

//定义一个存款业务流程
type DepositBusinessHandler struct {
	*BaseBusinessHandler
	userVip bool
}

//
//
func (p *DepositBusinessHandler) TakeRowNumber() {
	fmt.Println("请拿好您的取件码：" + strconv.Itoa(rand.Intn(100)) +
		" ，注意排队情况，过号后顺延六个安排")
}

//func (p *DepositBusinessHandler) WaitInHead() {
//	fmt.Println("排队等号中...")
//	time.Sleep(1 * time.Second)
//	fmt.Println("请去窗口xxx...")
//}
//
//func (p *DepositBusinessHandler) HandleBusiness() {
//	fmt.Println("账户存储很多万人民币...")
//}
//
//func (p *DepositBusinessHandler) Commentate() {
//	fmt.Println("请对我的服务作出评价，满意请按0，满意请按0，(～￣▽￣)～")
//}

func (p *DepositBusinessHandler) CheckVipIdentity() bool {
	return p.userVip
}



