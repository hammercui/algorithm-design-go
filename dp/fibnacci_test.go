package dp

import "testing"

func Test_fibnacci(t *testing.T)  {
	n := 10
	v1Result := getFibnacciV1(n)
	t.Log("v1",v1Result)
}
