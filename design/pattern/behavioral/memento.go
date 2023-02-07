package behavioral
/**
行为型模式
镜像模式，用了生成对象状态的镜像，实现状态回滚。
分为：
1. originator 初始状态，并生成memento
2。 memento 状态的镜像，immutable
3 caretakers：存储memento的栈
 */
type Memento interface {
	getState() string
}
type StringMemento struct {
	state string
}

func (p *StringMemento) getState() string {
	return p.state
}
func NewStringMemento(state string) Memento {
	return &StringMemento{state: state}
}

type Originator struct {
	state string
}
func (p *Originator) createMemento() Memento {
	return NewStringMemento(p.state)
}
func (p *Originator) restoreMemento(m Memento)  {
	p.state = m.getState()
}
func (p *Originator) getState() string  {
	return p.state
}
func (p *Originator) setState(state string)  {
	p.state = state
}

type Caretaker struct {
	mementoArray []Memento
}

func (p *Caretaker) addMemento(m Memento)  {
	p.mementoArray = append(p.mementoArray,m)
}

func (p *Caretaker) getMemento() Memento  {
	size := len(p.mementoArray)
	if size > 0 {
		memento :=  p.mementoArray[size - 1]
		p.mementoArray = p.mementoArray[:size-1]
		return memento
	}
	return nil
}




