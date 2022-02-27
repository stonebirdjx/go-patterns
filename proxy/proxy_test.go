// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: proxy_test.go
// @Date: 2022/2/27 9:41
// @Desc: 代理模式测试方法
package proxy

import "testing"

func TestProxy(t *testing.T) {
	var sub Subject
	sub = &Proxy{}

	res := sub.Do()

	if res != "pre:real:after" {
		t.Fail()
	}
}
