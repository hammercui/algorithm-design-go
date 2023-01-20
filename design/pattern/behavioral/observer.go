package behavioral

import (
	"fmt"
	"sync"
)

type Action func()
type Function func(args ...interface{})

//生命消息总线
type EventBus interface {
	Subscribe(topic string, handler Action) error
	SubscribeFunc(topic string, handler Function) error
	Publish(topic string,args ...interface{})
}


type AsyncEventBus struct {
	actions map[string][]Action
	functions map[string][]Function
	lock sync.Mutex
}
// NewAsyncEventBus new
func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		actions: map[string][]Action{},
		functions: map[string][]Function{},
		lock:     sync.Mutex{},
	}
}

func (p *AsyncEventBus) SubscribeAction(topic string, handler Action) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	if s,ok := p.actions[topic];ok {
		s = append(s, handler)
		p.actions[topic] = s
	}else{
		p.actions[topic] = []Action{handler}
	}
	return nil
}

func (p *AsyncEventBus) SubscribeFunc(topic string, handler Function) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	if s,ok := p.functions[topic];ok {
		s = append(s, handler)
		p.functions[topic] = s
	}else{
		p.functions[topic] = []Function{handler}
	}
	return nil
}

func (p *AsyncEventBus) Publish(topic string, args ...interface{}) {
	if len(args) == 0 {
		if s,ok := p.actions[topic];ok {
			for i := 0; i < len(s); i++ {
				go s[i]()
			}
		}else {
			fmt.Printf("topic:%s action is empty \n",topic)
		}
	}else{
		if s,ok := p.functions[topic];ok {
			for i := 0; i < len(s); i++ {
				go s[i](args...)
			}
		}else {
			fmt.Printf("topic:%s function is empty \n",topic)
		}
	}
}
