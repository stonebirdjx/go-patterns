// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: template_test.go
// @Date: 2022/2/28 20:49
// @Desc: 模板模式测试方法
package template

import "testing"

func TestSend(t *testing.T) {
	tel := NewTelecomSms()
	err := tel.Send("test", 1239999)
	if err != nil {
		t.Fatal()
	}
}
