// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: factorysimple_test.go
// @Date: 2022/2/26 10:41
// @Desc: 简单工厂模式测试方法
package factory

import "testing"

// TestNewApiHi 是 hiApi的测试方法
func TestNewApiHi(t *testing.T) {
	resApi := NewApi(hi)
	if resApi == nil {
		t.Error("hi Api is nil")
		return
	}
	if resApi.say("hjx") != "hi,hjx" {
		t.Error("hi Api return err")
	}
}

// TestNewApiHollo 是 helloApi的测试方法
func TestNewApiHollo(t *testing.T) {
	resApi := NewApi(hello)
	if resApi == nil {
		t.Error("hello Api is nil")
		return
	}
	if resApi.say("hjx") != "hello,hjx" {
		t.Error("hello Api return err")
	}
}

// TestNewApiNil 测试nil 值
func TestNewApiNil(t *testing.T) {
	resApi := NewApi(selector(3))
	if resApi != nil {
		t.Error("resApi is not nil")
		return
	}
}
