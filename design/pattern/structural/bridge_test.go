package main

import "testing"

func TestBridged(t *testing.T) {
	//抽象层
	abstraction := &AdvancedRemote{}

	//1 tv实例化
	tv := &TV{}
	abstraction.SetDevice(tv)
	abstraction.volumeUp()

	//2 radio实例化
	radio :=&Radio{}
	abstraction.SetDevice(radio)
	abstraction.volumeUp()
}
