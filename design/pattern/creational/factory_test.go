package creational

import "testing"

func TestAbstractFactory(t *testing.T) {
	var factory AbstractFactory
	var tv ITelevision
	var air IAirConditioner

	factory = &HuaweiFactory{}
	tv = factory.CreateTelevision()
	tv.Watch()
	air = factory.CreateAirConditioner()
	air.SetTemperature(1)

	factory = &XiaomiFactory{}
	tv = factory.CreateTelevision()
	tv.Watch()
	air = factory.CreateAirConditioner()
	air.SetTemperature(5)
}
