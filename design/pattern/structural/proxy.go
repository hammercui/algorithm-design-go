package structural

import "fmt"

/**
结构型模式
代理模式
1 可用作cache layer
2 或者增强层，比如access控制
 */
type Vehicle interface {
	Drive()
}

type Car struct {
	
}
func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	driver *Driver
	vehicle Vehicle
}

func (c *CarProxy) Drive() {
	if c.driver.Age > 18{
		c.vehicle.Drive()
	}else {
		fmt.Println("Driver too young!")
	}
}

func NewCarProxy(driver *Driver) *CarProxy{
	return &CarProxy{driver: driver,vehicle: &Car{}}
}




