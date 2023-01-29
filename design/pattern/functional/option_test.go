package functional

import "testing"

func TestOption(t *testing.T) {

	err := httpRequest("GET","https://google.com",
		WithTimeout(5),WithData("test"))
	if err != nil{
		t.Fatal(err)
	}
}
