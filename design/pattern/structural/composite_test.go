package main

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	root := NewCompositeOrganization("北京总公司", 1)
	root.add(&HRDOrg{orgName: "总公司人力资源部", depth: 2})
	root.add(&FinanceOrg{orgName: "总公司财务部", depth: 2})

	compSh := NewCompositeOrganization("上海分公司", 2)
	compSh.add(&HRDOrg{orgName: "上海分公司人力资源部", depth: 3})
	compSh.add(&FinanceOrg{orgName: "上海分公司财务部", depth: 3})
	root.add(compSh)

	compGd := NewCompositeOrganization("广东分公司", 2)
	compGd.add(&HRDOrg{orgName: "广东分公司人力资源部", depth: 3})
	compGd.add(&FinanceOrg{orgName: "南京办事处财务部", depth: 3})
	root.add(compGd)

	fmt.Println("公司组织架构：")
	root.display()

	fmt.Println("各组织的职责：")
	root.duty()
}
