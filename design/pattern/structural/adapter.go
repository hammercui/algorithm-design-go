package main

import "fmt"

/**
结构型设计模式
适配器模式，适用于client与Service不兼容，实现的接口不同，需要一个适配器作为转换。
因此有以下四个角色
1 client interface: client与其交互，已存在的
2 client： 调用发起者
3 adapter: 实现client interface，包装server对象，可随时新增实现。
4 server: client需要的服务类 但与client不兼容
 */

//适配器接口
type ClientXMLAnalysis interface {
	analysis(xml string) string
}

//被2包装服务
type AnalysisServer struct {}
func (p *AnalysisServer) analysis(json string) string  {
	fmt.Println("AnalysisServer analysis json data")
	return "ok"
}

//适配器实现
type ClientXMLAnalysisAdapter struct {
	server *AnalysisServer
}
func NewClientXMLAnalysisAdapter() ClientXMLAnalysis {
	server := &AnalysisServer{}
	return &ClientXMLAnalysisAdapter{server}
}
func (p *ClientXMLAnalysisAdapter) analysis(xml string) string{
	fmt.Println("adapter convert xml    to json")
	json := xml
	return p.server.analysis(json)
}



