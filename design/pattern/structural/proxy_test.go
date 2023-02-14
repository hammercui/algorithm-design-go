package structural

import "testing"

func TestProxy(t *testing.T) {
	carProxy := NewCarProxy(&Driver{Age: 12})
	carProxy.Drive()


	carProxy2 := NewCarProxy(&Driver{Age: 19})
	carProxy2.Drive()
}
