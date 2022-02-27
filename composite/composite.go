// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: composite.go
// @Date: 2022/2/27 12:40
// @Desc: 组合模式：它在我们树型结构的问题中，模糊了简单元素和复杂元素的概念，
//		  客户程序可以像处理简单元素一样来处理复杂元素，从而使得客户程序与复杂元素的内部结构解耦。
package composite

// IOrganization 组织接口，都实现统计人数的功能
type IOrganization interface {
	Count() int
}

// Employee 员工
type Employee struct {
	Name string
}

// Count 人数统计
func (Employee) Count() int {
	return 1
}

// Department 部门
type Department struct {
	Name string

	SubOrganizations []IOrganization
}

// Count 人数统计
func (d Department) Count() int {
	c := 0
	for _, org := range d.SubOrganizations {
		c += org.Count()
	}
	return c
}

// AddSub 添加子节点
func (d *Department) AddSub(org IOrganization) {
	d.SubOrganizations = append(d.SubOrganizations, org)
}

// NewOrganization 构建组织架构 demo
func NewOrganization() IOrganization {
	root := &Department{Name: "root"}
	for i := 0; i < 10; i++ {
		root.AddSub(&Employee{})
		root.AddSub(&Department{Name: "sub", SubOrganizations: []IOrganization{&Employee{}}})
	}
	return root
}
