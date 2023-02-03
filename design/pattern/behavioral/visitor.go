package behavioral

import "math"

/**
行为行模式
访问者模式，一种数据和操作分离的设计模式
1 element表示数据
2 visitor表示操作
3 client去创建这些对象
 */
type Visitor interface {
	visitSquare(square *Square) float64
	visitCircle(circle *Circle) float64
}

type Element interface {
	accept(Visitor) float64
}
type Square struct {
	side float64
}
func (p *Square) accept(visitor Visitor) float64{
	return visitor.visitSquare(p)
}

type Circle struct {
	radius float64
}
func (p *Circle) accept(visitor Visitor) float64{
	return visitor.visitCircle(p)
}
//面积visitor
type AreaVisitor struct {
	
}

func (a AreaVisitor) visitSquare(square *Square) float64{
	return square.side * square.side
}

func (a AreaVisitor) visitCircle(circle *Circle) float64{
	return circle.radius * circle.radius * math.Pi
}



