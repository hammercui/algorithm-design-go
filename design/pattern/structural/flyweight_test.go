package structural

import "testing"

func TestFlyweight(t *testing.T) {
	//设置5种树木
	pear := NewTreeContext(1,"梨树","白色","梨树纹理",1,1)

	peach := NewTreeContext(2,"桃树","蓝色","桃树纹理",1,1)

	elm := NewTreeContext(3,"榆树","绿色","榆树纹理",1,1)

	locust := NewTreeContext(4,"槐树","橙色","槐树纹理",1,1)

	pine := NewTreeContext(5,"松树","黑色","松树纹理",1,1)

	var list []*TreeContext = []*TreeContext{}
	list = append(list, pear)
	list = append(list,peach)
	list = append(list,elm)
	list = append(list,locust)
	list = append(list,pine)
	for i := 0; i < 1000; i++ {
		treeType := i % 5 + 1
		tree := NewTreeContext(treeType,"","","",i,i)
		tree.draw()
	}
}
