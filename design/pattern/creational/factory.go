package creational

import "fmt"

/**
工厂模式
1. 简单工厂
2. 工厂方法
3. 抽象工厂
本文只讲下抽象工厂
 */

//电视这一物体的抽象
type ITelevision interface {
	Watch()
}
//温度计的抽象
type IAirConditioner interface {
	SetTemperature(int)
}
type AbstractFactory interface {
	CreateTelevision() ITelevision
	CreateAirConditioner() IAirConditioner
}

type HuaweiTV struct {
}
func (h *HuaweiTV) Watch() {
	fmt.Println("Huawei tv Watching")
}

type HuaweiAirConditioner struct {
	
}
func (h HuaweiAirConditioner) SetTemperature(i int) {
	fmt.Printf("Huawei set temperature: %d\n",i)
}

type HuaweiFactory struct {
}

func (p *HuaweiFactory) CreateTelevision() ITelevision {
	return &HuaweiTV{}
}

func (p *HuaweiFactory) CreateAirConditioner() IAirConditioner {
	return &HuaweiAirConditioner{}
}

type XiaomiTV struct {
}
func (h *XiaomiTV) Watch() {
	fmt.Println("Xiaomi tv Watching")
}

type XiaomiAirConditioner struct {

}
func (h XiaomiAirConditioner) SetTemperature(i int) {
	fmt.Printf("Xiaomi set temperature: %d\n",i)
}

type XiaomiFactory struct {
}

func (p *XiaomiFactory) CreateTelevision() ITelevision {
	return &XiaomiTV{}
}

func (p *XiaomiFactory) CreateAirConditioner() IAirConditioner {
	return &XiaomiAirConditioner{}
}