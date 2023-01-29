package main

import (
	"fmt"
	"strings"
)

/**
组合模式 结构型模式
 */

//component接口
type Organization interface {
	display()
	duty()
}

//cpmposite点包含component列表，只关心接口实现的方法并调用
type CompositeOrganization struct {
	orgName string
	depth int
	list []Organization
}
func NewCompositeOrganization(orgName string, depth int) *CompositeOrganization {
	return &CompositeOrganization{orgName,depth,[]Organization{}}
}
func (o *CompositeOrganization) add(org Organization){
	o.list = append(o.list, org)
}
func (o *CompositeOrganization) remove(org Organization){
	for i,val := range o.list {
		if val == org{
			o.list = append(o.list[:i], o.list[i+1:]...)
			return
		}
	}
	return
}
func (o *CompositeOrganization) display() {
	fmt.Println(strings.Repeat("-", o.depth * 2), " ", o.orgName)
	for _, val := range o.list {
		val.display()
	}
}
func (o *CompositeOrganization) duty() {
	for _, val := range o.list {
		val.duty()
	}
}
//leaf对象-人力资源部门
type HRDOrg struct {
	orgName string
	depth int
}
func (o *HRDOrg) display() {
	fmt.Println(strings.Repeat("-", o.depth * 2), " ", o.orgName)
}
func (o *HRDOrg) duty() {
	fmt.Println(o.orgName, "员工招聘培训管理")
}

// Leaf对象--财务部门
type FinanceOrg struct {
	orgName string
	depth   int
}
func (f *FinanceOrg) display() {
	fmt.Println(strings.Repeat("-", f.depth * 2), " ", f.orgName)
}
func (f *FinanceOrg) duty() {
	fmt.Println(f.orgName, "员工招聘培训管理")
}


