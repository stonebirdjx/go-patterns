// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: composite_test.go
// @Date: 2022/2/27 12:40
// @Desc: 组合模式的测试方法
package composite

import "testing"

func TestNewOrganization(t *testing.T) {
	got := NewOrganization().Count()
	if got != 20 {
		t.Fatal()
	}
}
