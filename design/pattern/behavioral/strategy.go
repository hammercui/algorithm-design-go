package behavioral

import "fmt"

/**
策略模式
1 策略接口
2 Ctx持有参数和策略实例
 */
type PayBehavior interface{
	OrderPay(px *PayCtx)
}

type PayCtx struct {
	payBehavior PayBehavior
	payParams map[string]interface{}
}

func (p *PayCtx) setPayBehavior(behavior PayBehavior)  {
	p.payBehavior = behavior
}
func (p *PayCtx) Pay()  {
	p.payBehavior.OrderPay(p)
}
func NewPayCtx(p PayBehavior) *PayCtx  {
	// 支付参数，Mock数据
	params := map[string]interface{} {
		"appId": "hammer",
		"mchId": 123456,
	}
	return &PayCtx{
		payBehavior: p,
		payParams: params,
	}
}

//微信支付策略
type WXPay struct {
	
}
func (W WXPay) OrderPay(px *PayCtx) {
	fmt.Printf("I am wxpay %v \n",px.payParams)
}

//支付宝策略
type AliPay struct {

}

func (a AliPay) OrderPay(px *PayCtx) {
	fmt.Printf("i ma alipay %v \n",px.payParams)
}



