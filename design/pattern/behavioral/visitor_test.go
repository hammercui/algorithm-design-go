package behavioral

import (
	"fmt"
	"testing"
)

func TestVisitor(t *testing.T) {
	square := &Square{side:2}
	circle := &Circle{radius:2}

	//计算面积的visitor
	visitor := &AreaVisitor{}
	fmt.Printf("square面积: %.2f\n", visitor.visitSquare(square))
	fmt.Printf("circle面积: %.2f\n", visitor.visitCircle(circle))
}
