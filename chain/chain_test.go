// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: chain_test.go
// @Date: 2022/2/28 21:23
// @Desc: 责任链模式测试方法
package chain

import "testing"

func TestAdSensitiveWordFilter_Filter(t *testing.T) {
	chain := &SensitiveWordFilterChain{}
	chain.AddFilter(&AdSensitiveWordFilter{})
	if chain.Filter("test") {
		t.Fatal()
	}

	chain.AddFilter(&PoliticalWordFilter{})
	if !chain.Filter("test") {
		t.Fatal()
	}
}
