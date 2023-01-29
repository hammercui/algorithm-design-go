package creational

import "sync"
/**
懒汉模式获得单例
 */
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
