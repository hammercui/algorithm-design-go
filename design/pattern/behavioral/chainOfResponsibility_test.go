package behavioral

import (
	"fmt"
	"testing"
)

//测试第一种责任链模式
func TestChainOfResponsibilityV1(t *testing.T) {
	patient := &patient{
		Name:              "hammer",
		RegistrationDone:  false,
		DoctorCheckUpDone: false,
		MedicineDone:      false,
		PaymentDone:       false,
	}

	startHandler := StartHandler{}
	startHandler.SetNext(&Reception{}). //挂号
		SetNext(&Clinic{}). //诊室看病
		SetNext(&Cashier{}).//收费处交钱
		SetNext(&Pharmacy{}) //药房拿药
	if err := startHandler.Execute(patient);err != nil{
		fmt.Println("Fail | Error:" + err.Error())
	}else{
		fmt.Println("success")
	}
}
func TestChainOfResponsibilityV2(t *testing.T) {
	hospital := &HospitalEngine{
		handlerChain: []HandlerFunc{},
		index:        0,
	}
	//添加处理链路
	hospital.Use(ReceptionHandler, //挂号
		ClinicHandler, //诊室看病
		CashierHandler,//收费处交钱
		PharmacyHandler,//药房拿药
		)
	patient := &patientV2{
		Name:              "hammer",
		RegistrationDone:  false,
		DoctorCheckUpDone: false,
		MedicineDone:      false,
		PaymentDone:       false,
		hospital:hospital,
	}
	hospital.Start(patient)
}