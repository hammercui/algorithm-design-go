package behavioral

import "fmt"

/**
责任链模式，常用于中间件 拦截器
业务中：可以无痛额增加业务流程
特点：链条是单向的，也可以为洋葱模型
v1: 面向对象的方式实现
 */

type patient struct {
	Name string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
}

//定义action接口
type PatientHandler interface{
	Execute(*patient) error
	SetNext(PatientHandler) PatientHandler
	Do(*patient) error
}

//定义抽象父类
type Next struct {
	nextHandler PatientHandler
}
func (n *Next) SetNext(handler PatientHandler) PatientHandler{
	n.nextHandler = handler
	return handler
}
func (n *Next) Execute(patient *patient) (err error) {
	if n.nextHandler == nil{
		return
	}
	if err = n.nextHandler.Do(patient); err != nil{
		return err
	}
	return n.nextHandler.Execute(patient)
}

//定义挂号处理
type Reception struct {
	Next
}
func (r *Reception) Do(p *patient) (err error) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		return
	}
	fmt.Println("Reception registering patient")
	p.RegistrationDone = true
	return
}

// Clinic 诊室处理器--用于医生给病人看病
type Clinic struct {
	Next
}
func (d *Clinic) Do(p *patient) (err error) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		return
	}
	fmt.Println("Doctor checking patient")
	p.DoctorCheckUpDone = true
	return
}

// Cashier 收费处处理器
type Cashier struct {
	Next
}
func (c *Cashier) Do(p *patient) (err error) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
		return
	}
	fmt.Println("Cashier getting money from patient patient")
	p.PaymentDone = true
	return
}

// Pharmacy 药房处理器
type Pharmacy struct {
	Next
}
func (m *Pharmacy) Do (p *patient) (err error) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
		return
	}
	fmt.Println("Pharmacy giving medicine to patient")
	p.MedicineDone = true
	return
}

//首次处理
type StartHandler struct {
	Next
}
func (h *StartHandler) Do (p *patient) (err error){
	// 空Handler 这里什么也不做 只是载体 do nothing...
	return
}

