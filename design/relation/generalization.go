package relation
/**
泛化 包含接口与实现
接口：虚线空心三角
实现：实线空心三角
 */

type HumanLike interface{
	Run()
	Walk()
}

type Human struct {
}

type StudentG struct {
	Human
}