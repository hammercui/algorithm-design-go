package convert

import "testing"

func Test_str2byte(t *testing.T)  {
	str := "hello world"
	t.Log("from",str)
	bytes := StringToBytes(str)
	outStr := BytesToString(bytes)
	t.Log("bytes",bytes)
	t.Log("outStr",outStr)
	//var a = [...]int{1,2,3}
}
