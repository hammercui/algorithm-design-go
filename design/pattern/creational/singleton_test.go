package creational

import (
	"fmt"
	"testing"
	"time"
)

func TestLazyObject(t *testing.T) {
	for i := 0; i <5;i++{
		go func() {
			instance := GetSingle()
			if instance.Name != "hammer" {
				t.Errorf("failed!")
			}else{
				fmt.Println("instance ok ")
			}

		}()
	}
	time.Sleep(1* time.Second)
}
