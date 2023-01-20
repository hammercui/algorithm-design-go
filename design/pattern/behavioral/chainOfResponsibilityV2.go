package behavioral

import "fmt"

/**
基础函数式实现的责任链模式
1 定义一个医院引擎
2 每个病人持有一个医院引擎，每个人的看病流程不一样，因此engine可自定义
V2: 函数式的方式实现
 */
type HandlerFunc func(*patientV2)
type patientV2 struct {
	Name string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
    hospital *HospitalEngine
}
func (p *patientV2) Next()  {
	p.hospital.Next(p)
}

type HospitalEngine struct {
	handlerChain []HandlerFunc
	index int8
}
func (p *HospitalEngine) Start(c *patientV2)  {
	if len(p.handlerChain) > 0{
		p.handlerChain[0](c)
	}else{
		fmt.Println("handlerChain is empty")
	}
}
func (p *HospitalEngine) Next(c *patientV2)  {
	p.index++
	for p.index < int8(len(p.handlerChain)) {
		p.handlerChain[p.index](c)
	}
}
func (e *HospitalEngine) Use(handler ...HandlerFunc){
	e.handlerChain = append(e.handlerChain, handler...)
}

//挂号处理器
func ReceptionHandler(c *patientV2)  {
	if c.RegistrationDone {
		fmt.Println("Patient registration already done")
		c.Next()
		return
	}
	fmt.Println("Reception registering patient")
	c.RegistrationDone = true
	c.Next()
}

//诊室处理器
func ClinicHandler(c *patientV2)  {
	if c.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		c.Next()
		return
	}
	fmt.Println("Doctor checking patient")
	c.DoctorCheckUpDone = true
	c.Next()
}
//收费处理
func CashierHandler(c *patientV2) {
	if c.PaymentDone {
		fmt.Println("Payment Done")
		c.Next()
		return
	}
	fmt.Println("Cashier getting money from patient patient")
	c.PaymentDone = true
	c.Next()
}
//药房处理
func PharmacyHandler(c *patientV2)   {
	if c.MedicineDone {
		fmt.Println("Medicine already given to patient")
		c.Next()
		return
	}
	fmt.Println("Pharmacy giving medicine to patient")
	c.MedicineDone = true
	c.Next()
}




