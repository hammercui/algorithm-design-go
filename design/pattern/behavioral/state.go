package behavioral

import "fmt"

/**
行为型模式
状态机模式，核心是把一个对象的大的行为逻辑，拆解为一个个小的类
主要包含:
1. Context 原始对象
2. State 公共接口
3. ConcreteState
 */

type State interface{
	clickLock()
	clickPlay()
	clickNext()
	clickPrevious()
}
type BaseState struct {
	player *PlayerContext
}
type ReadyState struct {
	BaseState
}

func (p *ReadyState) clickLock() {
	p.player.SetState(&LockedState{BaseState{player: p.player}})
}

func (p *ReadyState) clickPlay() {
	p.player.SetState(&PlayingState{BaseState{player: p.player}})
}

func (p *ReadyState) clickNext() {
	fmt.Println("ready state ,can't play next")
}

func (p *ReadyState) clickPrevious() {
	fmt.Println("ready state ,can't play previous")
}

type LockedState struct {
	BaseState
}

func (p *LockedState) clickLock() {
	fmt.Println("locked state ,do not again")
}

func (p *LockedState) clickPlay() {
	p.player.SetState(&ReadyState{BaseState{player: p.player}})
}

func (p *LockedState) clickNext() {
	fmt.Println("locked state ,can't play next")
}

func (p *LockedState) clickPrevious() {
	fmt.Println("locked state ,can't play next")
}

type PlayingState struct {
	BaseState
}

func (p *PlayingState) clickLock() {
	p.player.SetState(&LockedState{BaseState{player: p.player}})
}

func (p *PlayingState) clickPlay() {
	fmt.Println("playing state")
}

func (p *PlayingState) clickNext() {
	fmt.Println("playing next")
}

func (p *PlayingState) clickPrevious() {
	fmt.Println("playing previous")
}

type PlayerContext struct {
	state State
	playlist []string
	currentSong string
}
func (p *PlayerContext) clickLock() {
	p.state.clickLock()
}
func (p *PlayerContext) clickPlay() {
	p.state.clickPlay()
}
func (p *PlayerContext) clickNext() {
	p.state.clickNext()
}
func (p *PlayerContext) clickPrevious() {
	p.state.clickPrevious()
}
func (p *PlayerContext) SetState(state State){
	p.state = state
}
func NewPlayer() *PlayerContext  {
	initState := &ReadyState{}
	player :=  &PlayerContext{state: initState}
	initState.player = player
	return player
}

