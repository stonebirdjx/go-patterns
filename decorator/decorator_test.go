// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: decorator_test.go
// @Date: 2022/2/27 10:15
// @Desc: 装饰器模式测试方法
package decorator

import "testing"

func TestColorSquare_Draw(t *testing.T) {
	sq := Square{}
	csq := NewColorSquare(sq, "red")
	got := csq.Draw()
	if got != "this is a square, color is red" {
		t.Fatal()
	}
}
