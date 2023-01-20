package behavioral

import (
	"fmt"
	"testing"
	"time"
)

func subscribe1(){
	fmt.Println("subscribe1 receive ")
}
func subscribe2(args ...interface{}){
	fmt.Printf("subscribe2 receive a:%d,b:%d\n",args[0],args[1])
}

func TestObserver(t *testing.T) {
	eventBus := NewAsyncEventBus()
	eventBus.Subscribe("action",subscribe1)
	eventBus.SubscribeFunc("function",subscribe2)

	eventBus.Publish("action")
	eventBus.Publish("function",1,2)
	time.Sleep(1 * time.Second)

	eventBus.Remove("action",subscribe1)
	eventBus.RemoveFunc("function",subscribe2)

	time.Sleep(2 * time.Second)
}