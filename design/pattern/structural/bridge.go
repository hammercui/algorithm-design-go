package main
/**
结构型设计模式
桥接模式，是一种前置设计模式，用来拆分app的维度。其内部可使用其他设计模式
比如：
1不兼容的新功能使用adapter模式
2implementation实例化使用抽象工厂
3abstraction实例化使用Builder
下面用遥控器和电视的例子来解析bridge模式
1 Remote类：遥控器，属于abstraction层的父类
2 Device接口： 属于implementation层的接口
3 TV&Radio类： 属于implementation层的实现
4 AdvancedRemote类：属于abstraction的子类
 */
type Device interface {
	isEnable() bool
	getVolume() int
	setVolume(percent int)
	getChannel() int
	setChannel(channel int)
}
type BaseDevice struct {
	percent int
	channel int
	enable bool
}
type TV struct {
	BaseDevice
}

func (p *TV) isEnable() bool{
	return p.enable
}

func (p *TV) getVolume() int{
	return p.percent
}

func (p *TV) setVolume(percent int) {
	p.percent = percent
}

func (p *TV) getChannel() int {
	return p.channel
}

func (p *TV) setChannel(channel int) {
	p.channel = channel
}

type Radio struct{
	BaseDevice
}

func (p *Radio) isEnable() bool {
	return p.enable
}

func (p *Radio) getVolume() int {
	return p.percent
}

func (p *Radio) setVolume(percent int) {
	p.percent = percent
}

func (p *Radio) getChannel() int {
	return p.channel
}

func (p *Radio) setChannel(channel int) {
	p.channel = channel
}





