package main

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	//新建适配器
	adapter := NewClientXMLAnalysisAdapter()
	xml := "i am xml"

	//client调用适配器的功能
	result := adapter.analysis(xml)
	fmt.Println(result)
}
