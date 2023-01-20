package creational

import "sync"

type LazyObject struct {
	Name string
}
var instance *LazyObject
var once sync.Once

func GetSingle() *LazyObject {
	once.Do(func() {
		instance = &LazyObject{Name:"hammer"}
	})
	return instance
}
